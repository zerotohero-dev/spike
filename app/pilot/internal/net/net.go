//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"fmt"
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

func SaveAdminToken(token string) error {
	// TODO: if SPIKE Nexus has an existing admin token,
	// it should reject creating a new admin token.
	// the admin token change shall be done
	// either by updating the SPIKE Nexus db, and
	// it would not be a standard operation.

	return nil
}
