//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package state

import (
	"errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/spike/internal/net"
	"os"
	"sync"
)

var tokenMutex sync.RWMutex

func AdminToken() (string, error) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()

	// Try to read from file:
	tokenBytes, err := os.ReadFile(".spike-token")
	if err != nil {
		return "", errors.Join(
			errors.New("failed to read token from file"),
			err,
		)
	}

	return string(tokenBytes), nil
}

func SaveAdminToken(source *workloadapi.X509Source, token string) error {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	// TODO: pilot login => exchanges admin token with a short-lived token.
	// At initial phase there is only one admin user.
	// Later, the admin user can be able to create username/token associations.

	// Save token to file:
	err := os.WriteFile(".pilot-token", []byte(token), 0600)
	if err != nil {
		return errors.Join(errors.New("failed to save token to file"), err)
	}

	// Save the token to SPIKE Nexus
	// This token will be used for Nexus to generated
	// short-lived session tokens for the admin user.
	err = net.SaveAdminToken(source, token)
	if err != nil {
		return errors.Join(errors.New("failed to save token to nexus"), err)
	}

	return nil
}

func AdminTokenExists() bool {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	token, _ := AdminToken()
	return token != ""
}
