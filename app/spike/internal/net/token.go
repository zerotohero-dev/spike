//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/spike/internal/entity/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
)

func SaveAdminToken(source *workloadapi.X509Source, token string) error {
	// TODO: if SPIKE Nexus has an existing admin token,
	// it should reject creating a new admin token.
	// the admin token change shall be done
	// either by updating the SPIKE Nexus db, and
	// it would not be a standard operation.

	r := reqres.AdminTokenWriteRequest{
		Data: token,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("token: I am having problem generating the payload"),
			err,
		)
	}

	authorizer := newAuthorizer()
	tlsConfig := tlsconfig.MTLSClientConfig(source, source, *authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// TODO: magic string.
	path := "https://localhost:8553/v1/init"

	return net.Post(client, path, mr, func(*http.Response) {})
}

// TODO: better to use a full-blown CLI parser.
// cobra is typically the de-facto choice.
/*
package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "github.com/spf13/cobra"
)

type KeyValue struct {
    Key   string
    Value string
}

// parseKVPairs parses key=value arguments into structured data
func parseKVPairs(args []string) ([]KeyValue, error) {
    var pairs []KeyValue
    for _, arg := range args {
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid key-value pair: %s (expected format: key=value)", arg)
        }
        pairs = append(pairs, KeyValue{
            Key:   parts[0],
            Value: parts[1],
        })
    }
    return pairs, nil
}

// parseVersionList parses comma-separated version numbers
func parseVersionList(versions string) ([]int, error) {
    if versions == "" {
        return nil, nil
    }

    var result []int
    for _, v := range strings.Split(versions, ",") {
        num, err := strconv.Atoi(strings.TrimSpace(v))
        if err != nil {
            return nil, fmt.Errorf("invalid version number: %s", v)
        }
        result = append(result, num)
    }
    return result, nil
}

var rootCmd = &cobra.Command{
    Use:   "spike",
    Short: "Spike is a key-value store CLI",
    Long: `Spike is a command-line interface for managing key-value pairs
with versioning support. It allows you to store, retrieve, and manage
versioned key-value pairs at specific paths.`,
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize a new spike store",
    Args:  cobra.NoArgs,
    RunE: func(cmd *cobra.Command, args []string) error {
        // TODO: Implement initialization logic
        fmt.Println("Initializing spike store...")
        return nil
    },
}

var putCmd = &cobra.Command{
    Use:   "put <path> <key=value>...",
    Short: "Store key-value pairs at the specified path",
    Example: `  spike put /acme/app/db username=admin password=secret
  spike put /acme/app/settings port=8080 debug=true`,
    Args: cobra.MinimumNArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
        path := args[0]

        kvPairs, err := parseKVPairs(args[1:])
        if err != nil {
            return err
        }

        // TODO: Implement storage logic
        fmt.Printf("Storing at path: %s\n", path)
        for _, kv := range kvPairs {
            fmt.Printf("  %s: %s\n", kv.Key, kv.Value)
        }
        return nil
    },
}

var getCmd = &cobra.Command{
    Use:   "get <path>",
    Short: "Retrieve values at the specified path",
    Example: `  spike get /acme/app/db
  spike get /acme/app/settings -version=2`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        path := args[0]
        version, _ := cmd.Flags().GetInt("version")

        // TODO: Implement retrieval logic
        fmt.Printf("Retrieving from path: %s (version: %d)\n", path, version)
        return nil
    },
}

var deleteCmd = &cobra.Command{
    Use:   "delete <path>",
    Short: "Delete values at the specified path",
    Example: `  spike delete /acme/app/db
  spike delete /acme/app/settings -versions=1,2,3`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        path := args[0]
        versionsStr, _ := cmd.Flags().GetString("versions")

        versions, err := parseVersionList(versionsStr)
        if err != nil {
            return err
        }

        // TODO: Implement deletion logic
        fmt.Printf("Deleting from path: %s\n", path)
        if len(versions) > 0 {
            fmt.Printf("Versions to delete: %v\n", versions)
        }
        return nil
    },
}

var undeleteCmd = &cobra.Command{
    Use:   "undelete <path>",
    Short: "Restore previously deleted values at the specified path",
    Example: `  spike undelete /acme/app/db
  spike undelete /acme/app/settings -versions=1,2,3`,
    Args: cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        path := args[0]
        versionsStr, _ := cmd.Flags().GetString("versions")

        versions, err := parseVersionList(versionsStr)
        if err != nil {
            return err
        }

        // TODO: Implement undelete logic
        fmt.Printf("Undeleting at path: %s\n", path)
        if len(versions) > 0 {
            fmt.Printf("Versions to restore: %v\n", versions)
        }
        return nil
    },
}

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all stored paths",
    Args:  cobra.NoArgs,
    RunE: func(cmd *cobra.Command, args []string) error {
        // TODO: Implement listing logic
        fmt.Println("Listing all paths...")
        return nil
    },
}

func init() {
    // Add flags
    getCmd.Flags().Int("version", 0, "Version number to retrieve")
    deleteCmd.Flags().String("versions", "", "Comma-separated list of versions to delete")
    undeleteCmd.Flags().String("versions", "", "Comma-separated list of versions to restore")

    // Add commands to root
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(putCmd)
    rootCmd.AddCommand(getCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(undeleteCmd)
    rootCmd.AddCommand(listCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

*/
