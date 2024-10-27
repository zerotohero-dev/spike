//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
	"time"

	"github.com/zerotohero-dev/spike/internal/spiffe"
	"github.com/zerotohero-dev/spike/internal/system"
)

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

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("hello")
			case <-ctx.Done():
				return
			}
		}
	}()

	system.KeepAlive()
}
