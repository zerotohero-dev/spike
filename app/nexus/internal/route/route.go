//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"fmt"
	"net/http"
)

func factory(p, a, m string) handler {
	fmt.Println("factory:", "p", p, "a", a, "m", m)

	switch {
	// Route to fetch the Keystone status.
	// The status can be "pending" or "ready".
	case m == http.MethodPost && a == "" && p == urlInit:
		return routeInit
	case m == http.MethodPost && a == "" && p == urlSecrets:
		return routePostSecret
	case m == http.MethodPost && a == "get" && p == urlSecrets:
		return routeGetSecret
	// Fallback route.
	default:
		return routeFallback
	}
}

func Route(w http.ResponseWriter, r *http.Request) {
	factory(r.URL.Path, r.URL.Query().Get("action"), r.Method)(r, w)
}
