//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/spike/internal/entity/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
)

func SaveAdminToken(source *workloadapi.X509Source, token string) error {
	// TODO: if SPIKE Nexus has an existing admin token,
	// it should reject creating a new admin token.
	// the admin token change shall be done
	// either by updating the SPIKE Nexus db, and
	// it would not be a standard operation.

	r := reqres.AdminTokenWriteRequest{
		Data: token,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("token: I am having problem generating the payload"),
			err,
		)
	}

	authorizer := newAuthorizer()
	tlsConfig := tlsconfig.MTLSClientConfig(source, source, *authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// TODO: magic string.
	path := "https://localhost:8553/v1/init"

	return net.Post(client, path, mr, func(*http.Response) {})
}
