//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

var reader = rand.Read

func Aes256Seed() (string, error) {
	// Generate a 256-bit key
	key := make([]byte, 32)

	_, err := reader(key)
	if err != nil {
		return "", errors.Join(
			err,
			errors.New("Aes256Seed: failed to generate random key"),
		)
	}

	return hex.EncodeToString(key), nil
}
