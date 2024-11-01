//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package state

import (
	"sync"

	"github.com/zerotohero-dev/spike/app/nexus/internal/state/store"
	"github.com/zerotohero-dev/spike/internal/crypto"
)

var (
	rootKey   string
	rootKeyMu sync.RWMutex
)

var (
	adminToken   string
	adminTokenMu sync.RWMutex
)

var (
	kv   = store.NewKV()
	kvMu sync.RWMutex
)

func AdminToken() string {
	adminTokenMu.RLock()
	defer adminTokenMu.RUnlock()
	return adminToken
}

func SetAdminToken(token string) {
	adminTokenMu.Lock()
	defer adminTokenMu.Unlock()
	adminToken = token
}

func UpsertSecret(path string, values map[string]string) {
	kvMu.Lock()
	defer kvMu.Unlock()

	kv.Put(path, values)
}

func GetSecret(path string, version int) (map[string]string, bool) {
	kvMu.RLock()
	defer kvMu.RUnlock()

	return kv.Get(path, version)
}

func Initialize() error {
	existingRootKey := RootKey()
	if existingRootKey != "" { // if key empty, try getting it from SPIKE Keeper
		// Already initialized.
		return nil
	}

	r, err := crypto.Aes256Seed()
	if err != nil {
		return err
	}

	rootKeyMu.Lock()
	rootKey = r
	rootKeyMu.Unlock()

	return nil
}

func RootKey() string {
	rootKeyMu.RLock()
	defer rootKeyMu.RUnlock()
	return rootKey
}
