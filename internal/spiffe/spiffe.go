//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package spiffe

import (
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
)

func AppSpiffeSource(ctx context.Context) (*workloadapi.X509Source, string) {
	// TODO: get this from env.
	socketPath := "unix:///tmp/spire-agent/public/api.sock"

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr(socketPath),
		),
	)
	if err != nil {
		log.Fatalf("Unable to create X509Source: %v", err)
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.Fatalf("Unable to get X509SVID: %v", err)
	}

	return source, svid.ID.String()
}
