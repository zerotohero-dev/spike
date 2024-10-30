//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "time"

type RootKeyCacheRequest struct {
	RootKey string `json:"rootKey"`
	Err     string `json:"err,omitempty"`
}

type RootKeyCacheResponse struct {
	Err string `json:"err,omitempty"`
}

type AdminTokenWriteRequest struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}

type AdminTokenWriteResponse struct {
	Err string `json:"err,omitempty"`
}

type SecretResponseMetadata struct {
	CreatedTime time.Time  `json:"created_time"`
	Version     int        `json:"version"`
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	Err         string     `json:"err,omitempty"`
}

// SecretPutRequest for creating/updating secrets
type SecretPutRequest struct {
	Path   string            `json:"path"`
	Values map[string]string `json:"values"`
	Err    string            `json:"err,omitempty"`
}

// SecretPutResponse after successful write
type SecretPutResponse struct {
	SecretResponseMetadata
	Err string `json:"err,omitempty"`
}
