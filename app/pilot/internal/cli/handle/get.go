//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func Get(source *workloadapi.X509Source, args []string) {

	if len(args) < 3 {
		fmt.Println("Usage: pilot get <path> [-version=<n>]")
		return
	}

	// TODO: implement me.

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
}
