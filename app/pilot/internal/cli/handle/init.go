//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/zerotohero-dev/spike/app/pilot/internal/net"
	"github.com/zerotohero-dev/spike/internal/crypto"
	"log"
)

func Init(args []string) {
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
