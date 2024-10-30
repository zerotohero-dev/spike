//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
)

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
