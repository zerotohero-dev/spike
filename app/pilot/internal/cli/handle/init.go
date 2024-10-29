//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/pilot/internal/state"
	"github.com/zerotohero-dev/spike/internal/crypto"
)

// TODO: create sequence diagrams to reason about what's happening and see
// if there are any gaps in the implementation.

// TODO: initialize a website; does not have to be fancy.

func Init(source *workloadapi.X509Source, args []string) {
	if state.AdminTokenExists() {
		fmt.Println("SPIKE Pilot is already initialized.")
		fmt.Println("Nothing to do.")
		return
	}

	// Generate and set the token
	token := crypto.Token()
	err := state.SaveAdminToken(source, token)
	if err != nil {
		fmt.Println("Failed to save admin token:")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
	fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
	fmt.Println(` \\\\\\\ web: spike.ist source: github.com/zerotohero-dev/spike`)
	fmt.Println("")
	fmt.Println("SPIKE system initialization completed.")
	fmt.Println("")
	fmt.Println("Admin Token:", token)
	fmt.Println("")
	fmt.Println("Please save this token securely. It will not be shown again.")
	fmt.Println("Token is saved to ./.pilot-token for future use.")
	fmt.Println("")
}
