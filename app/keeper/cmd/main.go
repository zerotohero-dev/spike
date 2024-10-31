//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"github.com/zerotohero-dev/spike/app/keeper/internal/server"
	"github.com/zerotohero-dev/spike/app/keeper/internal/validation"
	"github.com/zerotohero-dev/spike/internal/config"
	"log"

	"github.com/zerotohero-dev/spike/internal/spiffe"
)

const appName = "SPIKE Keeper"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, spiffeid := spiffe.AppSpiffeSource(ctx)
	defer spiffe.CloseSource(source)

	if !validation.IsKeeper(spiffeid) {
		log.Fatalf("SPIFFE ID %s is not valid.\n", spiffeid)
	}

	// Start the server
	log.Printf("Started service: %s v%s\n", appName, config.KeeperVersion)
	if err := server.Serve(source); err != nil {
		log.Fatalf("%s: Failed to serve: %s\n", appName, err.Error())
	}
}
