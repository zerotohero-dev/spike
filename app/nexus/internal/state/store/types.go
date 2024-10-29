//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package store

import "time"

type Version struct {
	Data        map[string]string
	CreatedTime time.Time
	Version     int
	DeletedTime *time.Time // nil if not deleted
}

type Metadata struct {
	CurrentVersion int
	OldestVersion  int
	CreatedTime    time.Time
	UpdatedTime    time.Time
	MaxVersions    int
}

type Secret struct {
	Versions map[int]Version
	Metadata Metadata
}
