//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
	"github.com/zerotohero-dev/spike/internal/entity/v1/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
)

func routeGetSecret(r *http.Request, w http.ResponseWriter) {
	fmt.Println("routeGetSecret:", r.Method, r.URL.Path, r.URL.RawQuery)

	w.WriteHeader(http.StatusOK)

	body := net.ReadRequestBody(r, w)
	if body == nil {
		return
	}

	var req reqres.SecretReadRequest
	if err := net.HandleRequestError(w, json.Unmarshal(body, &req)); err != nil {
		log.Println("routeInit: Problem handling request:", err.Error())
		return
	}

	version := req.Version
	path := req.Path

	secret, exists := state.GetSecret(path, version)
	if !exists {
		log.Println("routeGetSecret: Secret not found")
		w.WriteHeader(http.StatusNotFound)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routeGetSecret: Problem writing response:", err.Error())
		}
		return
	}

	res := reqres.SecretReadResponse{Data: secret}
	md, err := json.Marshal(res)
	if err != nil {
		log.Println("routeGetSecret: Problem generating response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("routeGetSecret: got secret")

	_, err = io.WriteString(w, string(md))
	if err != nil {
		log.Println("routeGetSecret: Problem writing response:", err.Error())
	}
}
