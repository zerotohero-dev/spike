//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package system

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func KeepAlive() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	log.Printf("\nReceived %v signal, shutting down gracefully...\n", sig)
}
