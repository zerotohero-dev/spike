//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import "fmt"

func printUsage() {
	fmt.Println("Usage: spike <command> [args...]")
	fmt.Println("Commands:")
	fmt.Println("  init")
	fmt.Println("  put <path> <key=value>...")
	fmt.Println("  get <path> [-version=<n>]")
	fmt.Println("  delete <path> [-versions=<n1,n2,...>]")
	fmt.Println("  undelete <path> [-versions=<n1,n2,...>]")
	fmt.Println("  list")
}
