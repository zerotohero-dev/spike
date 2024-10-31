//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"log"
	"os"

	"github.com/zerotohero-dev/spike/app/spike/internal/cli"
	"github.com/zerotohero-dev/spike/internal/config"
	"github.com/zerotohero-dev/spike/internal/spiffe"
)

const appName = "SPIKE"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, spiffeid := spiffe.AppSpiffeSource(ctx)

	if !config.IsPilot(spiffeid) {
		log.Fatalf("SPIFFE ID %s is not valid.\n", spiffeid)
	}

	defer spiffe.CloseSource(source)

	log.Printf("%s v%s\n", appName, config.PilotVersion)

	cli.Parse(source, os.Args)
}
