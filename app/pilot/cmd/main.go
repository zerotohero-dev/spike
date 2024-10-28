//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
	"os"
	"time"

	"github.com/zerotohero-dev/spike/internal/spiffe"
)

type Version struct {
	Data        map[string]string
	CreatedTime time.Time
	Version     int
	DeletedTime *time.Time // nil if not deleted
}

type Metadata struct {
	CurrentVersion int
	OldestVersion  int
	CreatedTime    time.Time
	UpdatedTime    time.Time
	MaxVersions    int
}

type Secret struct {
	Versions map[int]Version
	Metadata Metadata
}

// KVStore represents an in-memory key-value store with versioning
type KVStore struct {
	data map[string]*Secret
}

// NewKVStore creates a new KVStore instance
func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]*Secret),
	}
}

func (kv *KVStore) put(path string, values map[string]string) {
	secret, exists := kv.data[path]
	if !exists {
		secret = &Secret{
			Versions: make(map[int]Version),
			Metadata: Metadata{
				CreatedTime:    time.Now(),
				UpdatedTime:    time.Now(),
				MaxVersions:    3,
				CurrentVersion: 0,
				OldestVersion:  0,
			},
		}
		kv.data[path] = secret
	}

	// Increment version
	newVersion := secret.Metadata.CurrentVersion + 1

	// Add new version
	secret.Versions[newVersion] = Version{
		Data:        values,
		CreatedTime: time.Now(),
		Version:     newVersion,
	}

	// Update metadata
	secret.Metadata.CurrentVersion = newVersion
	secret.Metadata.UpdatedTime = time.Now()
	if secret.Metadata.OldestVersion == 0 {
		secret.Metadata.OldestVersion = 1
	}

	// Cleanup old versions if exceeding MaxVersions
	for version := range secret.Versions {
		if version <= secret.Metadata.CurrentVersion-secret.Metadata.MaxVersions {
			delete(secret.Versions, version)
			if version == secret.Metadata.OldestVersion {
				secret.Metadata.OldestVersion = version + 1
			}
		}
	}
}

func (kv *KVStore) get(path string, version int) (map[string]string, bool) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, false
	}

	// If version not specified, use current version
	if version == 0 {
		version = secret.Metadata.CurrentVersion
	}

	v, exists := secret.Versions[version]
	if !exists || v.DeletedTime != nil {
		return nil, false
	}

	return v.Data, true
}

func (kv *KVStore) delete(path string, versions []int) {
	secret, exists := kv.data[path]
	if !exists {
		return
	}

	now := time.Now()

	// If no versions specified, mark the latest version as deleted
	if len(versions) == 0 {
		if v, exists := secret.Versions[secret.Metadata.CurrentVersion]; exists {
			v.DeletedTime = &now
			secret.Versions[secret.Metadata.CurrentVersion] = v
		}
		return
	}

	// Delete specific versions
	for _, version := range versions {
		if v, exists := secret.Versions[version]; exists {
			v.DeletedTime = &now
			secret.Versions[version] = v
		}
	}
}

func (kv *KVStore) getMetadata(path string) (*Metadata, bool) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, false
	}
	return &secret.Metadata, true
}

func (kv *KVStore) deleteMetadata(path string) {
	delete(kv.data, path)
}

func (kv *KVStore) list() []string {
	keys := make([]string, 0, len(kv.data))
	for k := range kv.data {
		keys = append(keys, k)
	}
	return keys
}

func handleCommand(store *KVStore, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: pilot <command> [args...]")
		fmt.Println("Commands:")
		fmt.Println("  put <path> <key=value>... [-cas=<version>]")
		fmt.Println("  get <path> [-version=<n>]")
		fmt.Println("  delete <path> [-versions=<n1,n2,...>]")
		fmt.Println("  destroy <path> [-versions=<n1,n2,...>]")
		fmt.Println("  undelete <path> [-versions=<n1,n2,...>]")
		fmt.Println("  metadata get <path>")
		fmt.Println("  metadata delete <path>")
		fmt.Println("  list")
		return
	}

	/*
		most secret management systems (including Vault) return the secrets in plaintext because:

		They rely on TLS for transport security
		They enforce short-lived sessions/tokens
		They implement audit logging to track access
		They assume if you can authenticate and are authorized, you should get the actual secret value
		 that your policy authorizes you.


		======


			# First time system setup
			$ pilot init
			System initialized
			Admin Token: abc123def456...
			Please save this token securely. It will not be shown again.

			# Login with admin token
			$ pilot login --token abc123def456...
			Login successful. Session token saved.

			# Use the system with session token
			$ pilot get secret/foo

			# Admin can create more user tokens
			$ pilot token create --role operator
			Created token: xyz789...

	*/

	// TODO: these will be REST mTLS requests to Nexus
	// TODO: anything that "reads" a secret will read it encrypted
	// (so the operator will need to provide a key to decrypt it)
	// (but is that necessary? pilot is a superadmin tool anyway;
	// superadmin can provide a AES key but they will use the same
	// key to decrypt the secret, and when they do, the secret will
	// be visible on the terminal in plain text anyway.
	// unless we block reading secrets from SPIKE Pilot, it does not
	// make sense to return the secret encrypted.
	// if `get`
	command := args[1]
	switch command {
	case "put":
		log.Printf("Command: %s", command)
	case "get":
		log.Printf("Command: %s", command)
	case "delete":
		log.Printf("Command: %s", command)
	case "destroy":
		log.Printf("Command: %s", command)
	case "undelete":
		log.Printf("Command: %s", command)
	case "list":
		log.Printf("Command: %s", command)
	default:
		log.Printf("Unknown command: %s\n", command)
	}
}

const appName = "pilot"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// TODO: validate self spiffe id.

	source, spiffeid := spiffe.AppSpiffeSource(ctx)
	defer func(source *workloadapi.X509Source) {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Printf("Unable to close X509Source: %v", err)
		}
	}(source)

	store := NewKVStore()

	handleCommand(store, os.Args)

	log.Printf("SVID: %s", spiffeid)
	//log.Println(appName, "is running... Press Ctrl+C to exit")
	//system.KeepAlive()
}

//
//
//
//var (
//	ErrVersionNotFound = errors.New("version not found")
//	ErrSecretNotFound = errors.New("secret not found")
//	ErrVersionDestroyed = errors.New("version has been destroyed")
//	ErrInvalidVersion = errors.New("invalid version")
//)
//
//type Version struct {
//	Data        map[string]string
//	CreatedTime time.Time
//	Version     int
//	DeletedTime *time.Time // nil if not deleted (soft delete)
//	Destroyed   bool       // true if permanently destroyed
//}
//
//type Metadata struct {
//	CurrentVersion     int
//	OldestVersion     int
//	CreatedTime       time.Time
//	UpdatedTime       time.Time
//	MaxVersions       int
//	CasRequired       bool
//	DeleteVersionAfter time.Duration
//	CustomMetadata    map[string]string
//}
//
//type Secret struct {
//	Versions map[int]Version
//	Metadata Metadata
//}
//
//type KVStore struct {
//	data map[string]*Secret
//}
//
//func NewKVStore() *KVStore {
//	return &KVStore{
//		data: make(map[string]*Secret),
//	}
//}
//
//// put creates a new version of a secret
//func (kv *KVStore) put(path string, values map[string]string, cas int) error {
//	secret, exists := kv.data[path]
//	if !exists {
//		secret = &Secret{
//			Versions: make(map[int]Version),
//			Metadata: Metadata{
//				CreatedTime:       time.Now(),
//				UpdatedTime:       time.Now(),
//				MaxVersions:       10,
//				CurrentVersion:    0,
//				OldestVersion:     0,
//				CasRequired:       false,
//				DeleteVersionAfter: 0,
//				CustomMetadata:    make(map[string]string),
//			},
//		}
//		kv.data[path] = secret
//	}
//
//	// Check CAS if required
//	if secret.Metadata.CasRequired {
//		if cas != secret.Metadata.CurrentVersion {
//			return fmt.Errorf("cas mismatch: expected %d, got %d",
//				secret.Metadata.CurrentVersion, cas)
//		}
//	}
//
//	// Increment version
//	newVersion := secret.Metadata.CurrentVersion + 1
//
//	// Add new version
//	secret.Versions[newVersion] = Version{
//		Data:        values,
//		CreatedTime: time.Now(),
//		Version:     newVersion,
//		DeletedTime: nil,
//		Destroyed:   false,
//	}
//
//	// Update metadata
//	secret.Metadata.CurrentVersion = newVersion
//	secret.Metadata.UpdatedTime = time.Now()
//	if secret.Metadata.OldestVersion == 0 {
//		secret.Metadata.OldestVersion = 1
//	}
//
//	// Cleanup old versions if exceeding MaxVersions
//	kv.cleanupOldVersions(path)
//
//	return nil
//}
//
//// get retrieves a specific version of a secret
//func (kv *KVStore) get(path string, version int) (map[string]string, *Version, error) {
//	secret, exists := kv.data[path]
//	if !exists {
//		return nil, nil, ErrSecretNotFound
//	}
//
//	// If version not specified, use current version
//	if version == 0 {
//		version = secret.Metadata.CurrentVersion
//	}
//
//	v, exists := secret.Versions[version]
//	if !exists {
//		return nil, nil, ErrVersionNotFound
//	}
//
//	if v.Destroyed {
//		return nil, nil, ErrVersionDestroyed
//	}
//
//	return v.Data, &v, nil
//}
//
//// softDelete marks specified versions as deleted
//func (kv *KVStore) softDelete(path string, versions []int) error {
//	secret, exists := kv.data[path]
//	if !exists {
//		return ErrSecretNotFound
//	}
//
//	now := time.Now()
//
//	// If no versions specified, delete the latest version
//	if len(versions) == 0 {
//		versions = []int{secret.Metadata.CurrentVersion}
//	}
//
//	for _, version := range versions {
//		if v, exists := secret.Versions[version]; exists {
//			if v.Destroyed {
//				continue // Skip already destroyed versions
//			}
//			v.DeletedTime = &now
//			secret.Versions[version] = v
//		}
//	}
//
//	return nil
//}
//
//// destroy permanently removes the data for specified versions
//func (kv *KVStore) destroy(path string, versions []int) error {
//	secret, exists := kv.data[path]
//	if !exists {
//		return ErrSecretNotFound
//	}
//
//	// If no versions specified, destroy the latest version
//	if len(versions) == 0 {
//		versions = []int{secret.Metadata.CurrentVersion}
//	}
//
//	for _, version := range versions {
//		if v, exists := secret.Versions[version]; exists {
//			v.Destroyed = true
//			v.Data = nil // Remove sensitive data from memory
//			secret.Versions[version] = v
//		}
//	}
//
//	return nil
//}
//
//// undelete recovers soft-deleted versions
//func (kv *KVStore) undelete(path string, versions []int) error {
//	secret, exists := kv.data[path]
//	if !exists {
//		return ErrSecretNotFound
//	}
//
//	for _, version := range versions {
//		if v, exists := secret.Versions[version]; exists {
//			if v.Destroyed {
//				return ErrVersionDestroyed
//			}
//			v.DeletedTime = nil
//			secret.Versions[version] = v
//		}
//	}
//
//	return nil
//}
//
//// cleanupOldVersions removes versions exceeding MaxVersions
//func (kv *KVStore) cleanupOldVersions(path string) {
//	secret := kv.data[path]
//	if secret == nil {
//		return
//	}
//
//	// Sort versions by number
//	versions := make([]int, 0, len(secret.Versions))
//	for v := range secret.Versions {
//		versions = append(versions, v)
//	}
//
//	// Remove oldest versions if exceeding MaxVersions
//	if len(versions) > secret.Metadata.MaxVersions {
//		for version := range secret.Versions {
//			if version <= secret.Metadata.CurrentVersion-secret.Metadata.MaxVersions {
//				delete(secret.Versions, version)
//				if version == secret.Metadata.OldestVersion {
//					secret.Metadata.OldestVersion = version + 1
//				}
//			}
//		}
//	}
//}
//
//func handleCommand(store *KVStore, args []string) {
//	if len(args) < 2 {
//		fmt.Println("Usage: pilot <command> [args...]")
//		fmt.Println("Commands:")
//		fmt.Println("  put <path> <key=value>... [-cas=<version>]")
//		fmt.Println("  get <path> [-version=<n>]")
//		fmt.Println("  delete <path> [-versions=<n1,n2,...>]")
//		fmt.Println("  destroy <path> [-versions=<n1,n2,...>]")
//		fmt.Println("  undelete <path> [-versions=<n1,n2,...>]")
//		fmt.Println("  metadata get <path>")
//		fmt.Println("  metadata delete <path>")
//		fmt.Println("  list")
//		return
//	}
//
//	command := args[1]
//	switch command {
//	case "put":
//		if len(args) < 4 {
//			fmt.Println("Usage: pilot put <path> <key=value>... [-cas=<version>]")
//			return
//		}
//		values := make(map[string]string)
//		cas := 0
//		for _, arg := range args[3:] {
//			if strings.HasPrefix(arg, "-cas=") {
//				fmt.Sscanf(arg, "-cas=%d", &cas)
//				continue
//			}
//			kv := strings.Split(arg, "=")
//			if len(kv) == 2 {
//				values[kv[0]] = kv[1]
//			}
//		}
//		if err := store.put(args[2], values, cas); err != nil {
//			fmt.Printf("Error: %v\n", err)
//			return
//		}
//		fmt.Printf("Success! Data written to: %s\n", args[2])
//
//	case "get":
//		if len(args) < 3 {
//			fmt.Println("Usage: pilot get <path> [-version=<n>]")
//			return
//		}
//		version := 0
//		if len(args) > 3 && strings.HasPrefix(args[3], "-version=") {
//			fmt.Sscanf(args[3], "-version=%d", &version)
//		}
//		data, v, err := store.get(args[2], version)
//		if err != nil {
//			fmt.Printf("Error: %v\n", err)
//			return
//		}
//		fmt.Printf("=== Version %d ===\n", v.Version)
//		if v.DeletedTime != nil {
//			fmt.Printf("(Deleted at: %v)\n", *v.DeletedTime)
//		}
//		for k, v := range data {
//			fmt.Printf("%s: %s\n", k, v)
//		}
//
//	case "delete":
//		if len(args) < 3 {
//			fmt.Println("Usage: pilot delete <path> [-versions=<n1,n2,...>]")
//			return
//		}
//		versions := parseVersions(args)
//		if err := store.softDelete(args[2], versions); err != nil {
//			fmt.Printf("Error: %v\n", err)
//			return
//		}
//		fmt.Printf("Success! Versions marked as deleted at: %s\n", args[2])
//
//	case "destroy":
//		if len(args) < 3 {
//			fmt.Println("Usage: pilot destroy <path> [-versions=<n1,n2,...>]")
//			return
//		}
//		versions := parseVersions(args)
//		if err := store.destroy(args[2], versions); err != nil {
//			fmt.Printf("Error: %v\n", err)
//			return
//		}
//		fmt.Printf("Success! Versions destroyed at: %s\n", args[2])
//
//	case "undelete":
//		if len(args) < 3 {
//			fmt.Println("Usage: pilot undelete <path> [-versions=<n1,n2,...>]")
//			return
//		}
//		versions := parseVersions(args)
//		if err := store.undelete(args[2], versions); err != nil {
//			fmt.Printf("Error: %v\n", err)
//			return
//		}
//		fmt.Printf("Success! Versions recovered at: %s\n", args[2])
//
//	case "list":
//		keys := make([]string, 0, len(store.data))
//		for k := range store.data {
//			keys = append(keys, k)
//		}
//		if len(keys) == 0 {
//			fmt.Println("No secrets found")
//			return
//		}
//		fmt.Println("Secrets:")
//		for _, key := range keys {
//			fmt.Printf("- %s\n", key)
//		}
//
//	default:
//		fmt.Printf("Unknown command: %s\n", command)
//	}
//}
//
//// parseVersions helper function to parse version numbers from command args
//func parseVersions(args []string) []int {
//	versions := []int{}
//	for _, arg := range args {
//		if strings.HasPrefix(arg, "-versions=") {
//			versionStr := strings.TrimPrefix(arg, "-versions=")
//			for _, v := range strings.Split(versionStr, ",") {
//				var version int
//				fmt.Sscanf(v, "%d", &version)
//				versions = append(versions, version)
//			}
//			break
//		}
//	}
//	return versions
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	source, spiffeid := spiffe.AppSpiffeSource(ctx)
//	defer func(source *workloadapi.X509Source) {
//		if source == nil {
//			return
//		}
//		err := source.Close()
//		if err != nil {
//			log.Printf("Unable to close X509Source: %v", err)
//		}
//	}(source)
//
//	log.Printf("SVID: %s", spiffeid)
//
//	store := NewKVStore()
//
//	if len(os.Args) == 1 {
//		log.Println(appName, "is running in server mode... Press Ctrl+C to exit")
//		system.KeepAlive()
//		return
//	}
//
//	handleCommand(store, os.Args)
//}
