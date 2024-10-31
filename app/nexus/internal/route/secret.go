//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"encoding/json"
	"fmt"
	"github.com/zerotohero-dev/spike/internal/entity/v1/reqres"
	"io"
	"log"
	"net/http"

	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
)

func routePostSecret(r *http.Request, w http.ResponseWriter) {
	fmt.Println("routePostSecret:", r.Method, r.URL.Path, r.URL.RawQuery)

	w.WriteHeader(http.StatusOK)

	body, err := body(r)
	if err != nil {
		log.Println("routePostSecret: Problem reading request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}
	if body == nil {
		log.Println("routePostSecret: No request body.")
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}

	var req reqres.SecretPutRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("routePostSecret: Problem parsing request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}

	values := req.Values
	path := req.Path

	state.UpsertSecret(path, values)
	log.Println("routePostSecret: Secret upserted")

	_, err = io.WriteString(w, "")
	if err != nil {
		log.Println("routePostSecret: Problem writing response:", err.Error())
	}
}

func routeGetSecret(r *http.Request, w http.ResponseWriter) {
	fmt.Println("routeGetSecret:", r.Method, r.URL.Path, r.URL.RawQuery)

	w.WriteHeader(http.StatusOK)

	body, err := body(r)
	if err != nil {
		log.Println("routePostSecret: Problem reading request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}
	if body == nil {
		log.Println("routePostSecret: No request body.")
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}

	var req reqres.SecretReadRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("routePostSecret: Problem parsing request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}

	version := req.Version
	path := req.Path

	secret, exists := state.GetSecret(path, version)
	if !exists {
		log.Println("routePostSecret: Secret not found")
		w.WriteHeader(http.StatusNotFound)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routePostSecret: Problem writing response:", err.Error())
		}
		return
	}

	var res reqres.SecretReadResponse
	res.Data = secret

	md, err := json.Marshal(res)
	if err != nil {
		log.Println("routePostSecret: Problem generating response:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("routePostSecret: got secret")

	_, err = io.WriteString(w, string(md))
	if err != nil {
		log.Println("routePostSecret: Problem writing response:", err.Error())
	}
}
