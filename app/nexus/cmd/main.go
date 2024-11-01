//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"log"
	"time"

	"github.com/zerotohero-dev/spike/app/nexus/internal/poll"
	"github.com/zerotohero-dev/spike/app/nexus/internal/server"
	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
	"github.com/zerotohero-dev/spike/internal/config"
	"github.com/zerotohero-dev/spike/internal/spiffe"
)

const appName = "SPIKE Nexus"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, spiffeid, err := spiffe.AppSpiffeSource(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer spiffe.CloseSource(source)

	if !config.IsNexus(spiffeid) {
		log.Fatalf("SPIFFE ID %s is not valid.\n", spiffeid)
	}

	err := state.Initialize()
	if err != nil {
		log.Fatalf("Unable to initialize state: " + err.Error())
	}

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	go poll.Tick(ctx, source, ticker)

	log.Printf("Started service: %s v%s\n", appName, config.NexusVersion)
	if err := server.Serve(source); err != nil {
		log.Fatalf("%s: Failed to serve: %s\n", appName, err.Error())
	}
}
