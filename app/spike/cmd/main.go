//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/spike/internal/cli"
	"github.com/zerotohero-dev/spike/internal/spiffe"
	"log"
	"os"
)

const appName = "pilot"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// TODO: validate self spiffe id.

	source, _ := spiffe.AppSpiffeSource(ctx)
	defer func(source *workloadapi.X509Source) {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Printf("Unable to close X509Source: %v", err)
		}
	}(source)

	cli.Parse(source, os.Args)
}
