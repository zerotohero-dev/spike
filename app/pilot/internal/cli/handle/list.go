//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import "github.com/spiffe/go-spiffe/v2/workloadapi"

func List(source *workloadapi.X509Source, args []string) {
	panic("handleList not implemented")

	//		keys := make([]string, 0, len(store.data))
	//		for k := range store.data {
	//			keys = append(keys, k)
	//		}
	//		if len(keys) == 0 {
	//			fmt.Println("No secrets found")
	//			return
	//		}
	//		fmt.Println("Secrets:")
	//		for _, key := range keys {
	//			fmt.Printf("- %s\n", key)
	//		}
}
