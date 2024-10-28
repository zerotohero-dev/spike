//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"io"
	"log"
	"net/http"
)

type handler func(*http.Request, http.ResponseWriter)

const urlKeep = "/v1/keep"

func routeFallback(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.Println("routeFallback: Problem writing response:", err.Error())
	}
}

func routeKeep(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.Println("routeKeep: Problem writing response:", err.Error())
	}
}

func factory(p, m string) handler {
	log.Println("Factory:", p, urlKeep)

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

func Route(r *http.Request, w http.ResponseWriter) {
	factory(r.URL.Path, r.Method)(r, w)
}
