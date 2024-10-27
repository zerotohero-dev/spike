//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"log"

	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
)

func UpdateCache() error {
	log.Println("hello " + state.RootKey())

	// TODO: implement me.

	return nil
}
