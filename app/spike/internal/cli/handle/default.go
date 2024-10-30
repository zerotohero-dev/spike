//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func Default(source *workloadapi.X509Source, args []string) {
	printUsage()
}

func Usage(args []string) {
	if len(args) <= 1 {
		printUsage()
		return
	}

	fmt.Println("Unknown command:" + args[1])
	fmt.Println("")
	printUsage()
}

/*
	======


		# First time system setup
		$ pilot init
		System initialized


		# Login with admin token
		$ pilot login --token abc123def456...
		Login successful. Session token saved.

		# Use the system with session token
		$ pilot get secret/foo

		# Admin can create more user tokens
		$ pilot token create --role operator
		Created token: xyz789...

*/
