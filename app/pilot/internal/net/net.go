//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

func PutSecret() error {
	return nil
}

func GetSecret() error {
	return nil
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
