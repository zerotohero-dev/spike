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

const urlInit = "/v1/init"

func routeFallback(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.Println("routeFallback: Problem writing response:", err.Error())
	}
}

func routeInit(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)

	// TODO: implement me.

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.Println("routeInit: Problem writing response:", err.Error())
	}
}

func factory(p, m string) handler {
	log.Println("Factory:", p, urlInit)

	switch {
	// Route to fetch the Keystone status.
	// The status can be "pending" or "ready".
	case m == http.MethodPost && p == urlInit:
		return routeInit
	// Fallback route.
	default:
		return routeFallback
	}
}

func Route(r *http.Request, w http.ResponseWriter) {
	factory(r.URL.Path, r.Method)(r, w)
}
