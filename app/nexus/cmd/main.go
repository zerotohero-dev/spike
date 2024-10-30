//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"github.com/zerotohero-dev/spike/app/nexus/internal/server"
	"log"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/nexus/internal/net"
	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
	"github.com/zerotohero-dev/spike/internal/spiffe"
)

// TODO: update the webiste once there is something demoable.

const appName = "nexus"

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

	log.Printf("SPIFFEID: %s\n", spiffeid)
	log.Println(appName, "is running... Press Ctrl+C to exit")

	// TODO: if initialized already, do not re-init.

	err := state.Initialize()
	if err != nil {
		panic("Unable to initialize state: " + err.Error())
	}

	// TODO: validate self spiffe id.

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				err := net.UpdateCache(source)
				if err != nil {
					log.Printf("Unable to update cache: %v\n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	log.Println("Started server")
	err = server.Serve(source)
	if err != nil {
		log.Fatalln("failed to serve:", err.Error())
		return
	}
}
