//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package poll

import (
	"context"
	"log"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/nexus/internal/net"
)

// Tick continuously updates SPIKE Keeper, sending the root key to be backed up
// in memory.
//
// It runs until the provided context is cancelled.
//
// The function uses a select statement to either:
// 1. Update the cache when the ticker signals, or
// 2. Exit when the context is done
//
// Parameters:
//   - ctx: A context.Context for cancellation control
//   - source: A pointer to workloadapi.X509Source that provides the source data
//   - ticker: A time.Ticker that determines the update interval
//
// The function will log any errors that occur during cache updates but
// continue running.
//
// To stop the function, cancel the provided context.
//
// Example usage:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	source := &workloadapi.X509Source{...}
//	ticker := time.NewTicker(5 * time.Minute)
//	defer ticker.Stop()
//
//	go Tick(ctx, source, ticker)
func Tick(ctx context.Context,
	source *workloadapi.X509Source, ticker *time.Ticker) {
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
}
