//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cli

import (
	"fmt"
	"log"
)

func printUsage() {
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
}

func handleInit(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleGet(args []string) {
	log.Printf("Command: %s", args[1])
}
func handlePut(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleDelete(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleDestroy(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleUndelete(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleList(args []string) {
	log.Printf("Command: %s", args[1])
}
func handleDefault(args []string) {
	log.Printf("Command: %s", args[1])
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
