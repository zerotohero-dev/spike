//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"fmt"
	"github.com/zerotohero-dev/spike/app/spike/internal/entity/data"
	"net/url"
)

func GetSecret(path string, version int) (*data.Secret, error) {
	// TODO: create a server at Nexus 8553 to listen and respond a dummy response.
	secretUrl, err := url.JoinPath("https://localhost:8553/v1/", path,
		fmt.Sprintf("?version=%d", version),
	)
	if err != nil {
		return nil,
			errors.Join(errors.New("GetSecret: failed to join secret url"), err)
	}

	fmt.Println("fetch:", secretUrl)

	return nil, nil
}
