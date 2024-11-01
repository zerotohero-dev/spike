//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"io"
	"log"
	"net/http"
)

func routeKeep(_ *http.Request, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)

	log.Println("routeKeep::implement Me")

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.Println("routeKeep: Problem writing response:", err.Error())
	}
}
