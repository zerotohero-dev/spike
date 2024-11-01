//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"log"

	"github.com/zerotohero-dev/spike/app/keeper/internal/server"
	"github.com/zerotohero-dev/spike/internal/config"
	"github.com/zerotohero-dev/spike/internal/spiffe"
)

const appName = "SPIKE Keeper"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, spiffeid, err := spiffe.AppSpiffeSource(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer spiffe.CloseSource(source)

	if !config.IsKeeper(spiffeid) {
		log.Fatalf("SPIFFE ID %s is not valid.\n", spiffeid)
	}

	log.Printf("Started service: %s v%s\n", appName, config.KeeperVersion)
	if err := server.Serve(source); err != nil {
		log.Fatalf("%s: Failed to serve: %s\n", appName, err.Error())
	}
}
