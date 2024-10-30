//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/spike/internal/entity/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
	"net/http"
)

func PutSecret(source *workloadapi.X509Source,
	path string, values map[string]string) error {

	r := reqres.SecretPutRequest{
		Path:   path,
		Values: values,
	}
	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("putSecret: I am having problem generating the payload"),
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
	endpoint := "https://localhost:8553/v1/secrets"

	_, err = net.Post(client, endpoint, mr, func(*http.Response) {})

	return err
}
