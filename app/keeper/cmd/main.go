//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/keeper/internal/server"
	"log"

	"github.com/zerotohero-dev/spike/internal/spiffe"
)

const appName = "keeper"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, spiffeid := spiffe.AppSpiffeSource(ctx)
	defer func(source *workloadapi.X509Source) {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Printf("Unable to close X509Source: %v", err)
		}
	}(source)

	// TODO: validate self spiffeid.

	log.Printf("SPIFFEID: %s", spiffeid)

	// Start the server
	log.Println("Started server")
	err := server.Serve(source)
	if err != nil {
		log.Fatalln("failed to serve:", err.Error())
		return
	}
}
