//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package state

import (
	"sync"

	"github.com/zerotohero-dev/spike/app/nexus/internal/crypto"
)

var (
	rootKey   string
	rootKeyMu sync.RWMutex
)

func Initialize() error {
	r, err := crypto.Aes256Seed()
	if err != nil {
		return err
	}

	// TODO: save initialization status to Postgres.

	// TODO: ADR: Use postgres as a backing store.

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
