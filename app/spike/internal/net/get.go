//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/spike/internal/entity/data"
	"github.com/zerotohero-dev/spike/internal/entity/v1/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
)

// TODO: verify and cleanup.

func GetSecret(source *workloadapi.X509Source, path string, version int) (*data.Secret, error) {
	secretUrl, err := url.JoinPath("https://localhost:8553/v1/secrets?action=get")
	if err != nil {
		return nil,
			errors.Join(errors.New("GetSecret: failed to join secret url"), err)
	}

	r := reqres.SecretReadRequest{
		Path:    path,
		Version: version,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Join(
			errors.New("getSecret: I am having problem generating the payload"),
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

	endpoint := secretUrl

	body, err := net.Post(client, endpoint, mr, func(resp *http.Response) {})

	var res reqres.SecretReadResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Join(
			errors.New("getSecret: Problem parsing response body"),
			err,
		)
	}

	return &data.Secret{
		Data: res.Data,
	}, nil
}
