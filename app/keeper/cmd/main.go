//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"log"

	"github.com/zerotohero-dev/spike/internal/spiffe"
	"github.com/zerotohero-dev/spike/internal/system"
)

const appName = "keeper"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("SVID: %s", spiffe.AppSpiffeId(ctx))
	log.Println(appName, "is running... Press Ctrl+C to exit")

	// Start the server

	system.KeepAlive()
}
