//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package state

import (
	"log"
	"os"
	"sync"
)

var pilotToken = ""
var tokenMutex sync.RWMutex

func PilotToken() string {
	tokenMutex.RLock()
	// First check if token exists in memory
	if pilotToken != "" {
		defer tokenMutex.RUnlock()
		return pilotToken
	}
	tokenMutex.RUnlock()

	// If not in memory, try to read from file:
	tokenBytes, err := os.ReadFile(".pilot-token")
	if err != nil {
		log.Printf("Failed to read token from file: %v", err)
		return ""
	}

	// Update memory with token from file:
	token := string(tokenBytes)
	SavePilotToken(token)
	return token
}

func SavePilotToken(token string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	// Save token to memory:
	pilotToken = token
	// Save token to file:
	err := os.WriteFile(".pilot-token", []byte(token), 0600)
	if err != nil {
		log.Printf("Failed to save token to file: %v", err)
	}
}

func PilotTokenExists() bool {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	if pilotToken != "" {
		return true
	}
	pt := PilotToken()
	if pt != "" {
		pilotToken = pt
		return true
	}
	return false
}
