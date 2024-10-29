//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import "github.com/spiffe/go-spiffe/v2/workloadapi"

func Undelete(source *workloadapi.X509Source, args []string) {
	panic("handleUndelete not implemented")

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
}
