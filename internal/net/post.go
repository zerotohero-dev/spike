//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"errors"
	"net/http"
)

func Post(client *http.Client, path string, mr []byte,
	respond func(*http.Response)) error {
	r, err := client.Post(path, "application/json", bytes.NewBuffer(mr))

	if err != nil {
		return errors.Join(
			err,
			errors.New("post: Problem connecting to peer"),
		)
	}

	if r.StatusCode != http.StatusOK {
		return errors.New("post: Problem connecting to peer")
	}

	respond(r)
	return nil
}
