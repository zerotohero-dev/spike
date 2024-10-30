//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/spike/internal/net"
	"github.com/zerotohero-dev/spike/app/spike/internal/state"
	"strings"
)

func Put(source *workloadapi.X509Source, args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: pilot put <path> <key=value>...")
		return
	}

	adminToken, err := state.AdminToken()
	if err != nil {
		fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
		fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
		fmt.Println(` \\\\\\\ web: spike.ist source: github.com/zerotohero-dev/spike`)
		fmt.Println("")
		fmt.Println("SPIKE is not initialized.")
		fmt.Println("Please run `spike init` to initialize SPIKE.")
		return
	}

	// TODO: for now we verify the admin token; later down the line, we will
	// exchange the admin token with a short-lived token with `spike login`.
	if adminToken == "" {
		fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
		fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
		fmt.Println(` \\\\\\\ web: spike.ist source: github.com/zerotohero-dev/spike`)
		fmt.Println("")
		fmt.Println("SPIKE is not initialized.")
		fmt.Println("Please run `spike init` to initialize SPIKE.")
		return
	}

	path := args[2]
	values := make(map[string]string)
	for _, arg := range args[3:] {
		kv := strings.Split(arg, "=")
		if len(kv) == 2 {
			values[kv[0]] = kv[1]
		}
	}

	err = net.PutSecret(source, path, values)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("OK")
}
