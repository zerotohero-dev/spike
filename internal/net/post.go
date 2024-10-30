//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func body(r *http.Response) (bod []byte, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err = errors.Join(err, b.Close())
	}(r.Body)

	return body, err
}

func Post(client *http.Client, path string, mr []byte,
	respond func(*http.Response)) ([]byte, error) {
	r, err := client.Post(path, "application/json", bytes.NewBuffer(mr))

	if err != nil {
		return []byte{}, errors.Join(
			errors.New("post: Problem connecting to peer"),
			err,
		)
	}

	if r.StatusCode != http.StatusOK {
		return []byte{}, errors.New("post: Problem connecting to peer")
	}

	b, err := body(r)
	if err != nil {
		return []byte{}, errors.Join(
			errors.New("post: Problem reading response body"),
			err,
		)
	}

	respond(r)

	return b, nil
}
