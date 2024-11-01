//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

import "os"

// TlsPort returns the TLS port for the Spike Keeper service.
// It first checks for a port specified in the SPIKE_KEEPER_TLS_PORT
// environment variable.
// If no environment variable is set, it defaults to ":8443".
//
// The returned string is in the format ":port" suitable for use with
// net/http Listen functions.
func TlsPort() string {
	p := os.Getenv("SPIKE_KEEPER_TLS_PORT")

	if p != "" {
		return p
	}

	return ":8443"
}
