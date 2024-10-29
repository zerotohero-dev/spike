//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/pilot/internal/entity/reqres"
	"github.com/zerotohero-dev/spike/internal/net"
	"net/http"
	"net/url"

	"github.com/zerotohero-dev/spike/app/pilot/internal/entity/data"
)

func PutSecret() error {
	return nil
}

func GetSecret(path string, version int) (*data.Secret, error) {
	// TODO: create a server at Nexus 8553 to listen and respond a dummy response.
	secretUrl, err := url.JoinPath("https://localhost:8553/v1/", path,
		fmt.Sprintf("?version=%d", version),
	)
	if err != nil {
		return nil, errors.Join(errors.New("GetSecret: failed to join secret url"), err)
	}

	fmt.Println("fetch:", secretUrl)

	return nil, nil
}

func DeleteSecret() error {
	return nil
}

func DestroySecret() error {
	return nil
}

func UndeleteSecret() error {
	return nil
}

func ListSecretKeys() error {
	return nil
}

func newAuthorizer() *tlsconfig.Authorizer {
	auth := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		// TODO: implement me.
		return nil
	})

	return &auth
}

//func doPost(client *http.Client, path string, mr []byte) error {
//	r, err := client.Post(path, "application/json", bytes.NewBuffer(mr))
//
//	if err != nil {
//		return errors.Join(
//			err,
//			errors.New("post: Problem connecting to SPIKE Nexus API endpoint URL"),
//		)
//	}
//
//	if r.StatusCode != http.StatusOK {
//		return errors.New("post: Problem connecting to SPIKE Nexus API endpoint URL")
//	}
//
//	respond(r)
//	return nil
//}

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

	path := "https://localhost:8553/v1/init"

	return net.Post(client, path, mr, func(*http.Response) {})
}
