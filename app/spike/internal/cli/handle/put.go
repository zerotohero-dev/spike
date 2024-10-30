//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
)

func Put(source *workloadapi.X509Source, args []string) {
	log.Printf("Command: %s", args[1])

	//		if len(args) < 4 {
	//			fmt.Println("Usage: pilot put <path> <key=value>... [-cas=<version>]")
	//			return
	//		}
	//		values := make(map[string]string)
	//		cas := 0
	//		for _, arg := range args[3:] {
	//			kv := strings.Split(arg, "=")
	//			if len(kv) == 2 {
	//				values[kv[0]] = kv[1]
	//			}
	//		}
	//		if err := store.put(args[2], values); err != nil {
	//			fmt.Printf("Error: %v\n", err)
	//			return
	//		}
	//		fmt.Printf("Success! Data written to: %s\n", args[2])
}
