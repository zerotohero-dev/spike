//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package validation

import "github.com/zerotohero-dev/spike/app/keeper/internal/env"

// IsNexus checks if the provided SPIFFE ID matches the SPIKE Nexus SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured Spike Nexus
// SPIFFE ID from the environment. This is typically used for validating whether
// a given identity represents the Nexus service.
//
// Parameters:
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches the Nexus SPIFFE ID, false otherwise
func IsNexus(spiffeid string) bool {
	return spiffeid == env.SpikeNexusSpiffeId()
}

// IsKeeper checks if the provided SPIFFE ID matches the Spike Keeper SPIFFE ID.
//
// The function compares the input SPIFFE ID against the configured Spike Keeper
// SPIFFE ID from the environment. This is typically used for validating whether
// a given identity represents the Keeper service.
//
// Parameters:
//   - spiffeid: The SPIFFE ID string to check
//
// Returns:
//   - bool: true if the SPIFFE ID matches the Keeper SPIFFE ID, false otherwise
func IsKeeper(spiffeid string) bool {
	return spiffeid == env.SpikeKeeperSpiffeId()
}
