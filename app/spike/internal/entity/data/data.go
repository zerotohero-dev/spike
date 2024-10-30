//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package data

import "time"

type SecretMetaData struct {
	CreatedTime time.Time  `json:"created_time"`
	Version     int        `json:"version"`
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
}

type Secret struct {
	Data     map[string]string `json:"data"`
	Metadata SecretMetaData    `json:"metadata"`
}
