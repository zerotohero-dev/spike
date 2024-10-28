//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/zerotohero-dev/spike/app/pilot/internal/net"
	"github.com/zerotohero-dev/spike/internal/crypto"
)

func printUsage() {
	fmt.Println("Usage: pilot <command> [args...]")
	fmt.Println("Commands:")
	fmt.Println("  init <username>")
	fmt.Println("  put <path> <key=value>... [-cas=<version>]")
	fmt.Println("  get <path> [-version=<n>]")
	fmt.Println("  delete <path> [-versions=<n1,n2,...>]")
	fmt.Println("  destroy <path> [-versions=<n1,n2,...>]")
	fmt.Println("  undelete <path> [-versions=<n1,n2,...>]")
	fmt.Println("  metadata get <path>")
	fmt.Println("  metadata delete <path>")
	fmt.Println("  list")
}

var pilotToken = ""
var tokenMutex sync.RWMutex

func getPilotToken() string {
	tokenMutex.RLock()
	// First check if token exists in memory
	if pilotToken != "" {
		defer tokenMutex.RUnlock()
		return pilotToken
	}
	tokenMutex.RUnlock()

	// If not in memory, try to read from file
	tokenBytes, err := os.ReadFile(".pilot-token")
	if err != nil {
		log.Printf("Failed to read token from file: %v", err)
		return ""
	}

	// Update memory with token from file
	token := string(tokenBytes)
	setPilotToken(token)
	return token
}

func setPilotToken(token string) {
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

func hasPilotToken() bool {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	if pilotToken != "" {
		return true
	}
	pt := getPilotToken()
	if pt != "" {
		pilotToken = pt
		return true
	}
	return false
}

func handleInit(args []string) {
	if hasPilotToken() {
		fmt.Println("SPIKE Pilot is already initialized.")
		fmt.Println("Nothing to do.")
		return
	}

	// Generate and set the token
	token := crypto.Token()
	setPilotToken(token)

	// Save the token to SPIKE Nexus
	// This token will be used for Nexus to generated
	// short-lived session tokens for the admin user.
	err := net.SaveAdminToken(token)
	if err != nil {
		log.Println("Failed to save admin token to SPIKE Nexus: " + err.Error())
		return
	}

	fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
	fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
	fmt.Println(` \\\\\\\ SPDX-License-Identifier: Apache-2.0`)
	fmt.Println("")
	fmt.Println("SPIKE system initialization completed.")
	fmt.Println("")
	fmt.Println("Admin Token:", token)
	fmt.Println("")
	fmt.Println("Please save this token securely. It will not be shown again.")
	fmt.Println("Token is saved to ./.pilot-token for future use.")
	fmt.Println("")
}

func handleGet(args []string) {
	log.Printf("Command: %s", args[1])
}

func handlePut(args []string) {
	log.Printf("Command: %s", args[1])
}

func handleDelete(args []string) {
	panic("handleDelete not implemented")
}

func handleDestroy(args []string) {
	panic("handleDestroy not implemented")
}

func handleUndelete(args []string) {
	panic("handleUndelete not implemented")
}

func handleList(args []string) {
	panic("handleList not implemented")
}

func handleDefault(args []string) {
	printUsage()
}

func HandleCommand(args []string) {
	if len(args) < 2 {
		printUsage()
	}

	command := args[1]
	switch command {
	case "init":
		handleInit(args)
	case "put":
		handlePut(args)
	case "get":
		handleGet(args)
	case "delete":
		handleDelete(args)
	case "destroy":
		handleDestroy(args)
	case "undelete":
		handleUndelete(args)
	case "list":
		handleList(args)
	default:
		handleDefault(args)
	}
}
