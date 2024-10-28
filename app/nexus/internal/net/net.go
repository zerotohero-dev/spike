//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-jose/go-jose/v4/json"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/nexus/internal/entity/v1/reqres"
	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
	"github.com/zerotohero-dev/spike/app/nexus/internal/validation"
)

func newRootKeyCacheRequest(rootKey string) reqres.RootKeyCacheRequest {
	return reqres.RootKeyCacheRequest{
		RootKey: rootKey,
	}
}

func createAuthorizer() tlsconfig.Authorizer {
	return tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsKeeper(id.String()) {
			return nil
		}

		return errors.New("Post: I don't know you, and it's crazy: '" +
			id.String() + "'")
	})
}

func respond(r *http.Response) {
	if r == nil {
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Post: Problem closing request body.", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(
			"Post: Unable to read the response body from VSecM Safe.",
			err.Error())
		return
	}

	log.Println("response:", string(body))
}

func doPost(client *http.Client, p string, md []byte) error {
	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))

	if err != nil {
		return errors.Join(
			err,
			errors.New("post: Problem connecting to SPIKE Keep:"+err.Error()),
		)
	}

	if r.StatusCode != http.StatusOK {
		return errors.New("post: Problem connecting SPIKE Keep: status:" + r.Status)
	}

	respond(r)

	return nil
}

func UpdateCache(source *workloadapi.X509Source) error {
	if source == nil {
		return errors.New("UpdateCache: got nil source")
	}

	authorizer := createAuthorizer()

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	rr := newRootKeyCacheRequest(state.RootKey())
	md, err := json.Marshal(rr)
	if err != nil {
		return errors.New("UpdateCache: failed to marshal request: " + err.Error())
	}

	path := "https://localhost:8443/v1/keep"

	return doPost(client, path, md)
}
