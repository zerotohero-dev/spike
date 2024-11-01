//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"net/http"
)

func factory(p, m string) handler {
	switch {
	// Route to fetch the Keystone status.
	// The status can be "pending" or "ready".
	case m == http.MethodPost && p == urlKeep:
		return routeKeep
	// Fallback route.
	default:
		return routeFallback
	}
}
