//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "time"

// TODO: these entities are duplicated. move this to a higher level.

type RootKeyCacheRequest struct {
	RootKey string `json:"rootKey"`
	// TODO: we don't use Err fields at all: Maybe remove them?
	Err string `json:"err,omitempty"`
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

// SecretReadRequest for getting secrets (query params in URL)
type SecretReadRequest struct {
	Path    string `json:"path"`
	Version int    `json:"version,omitempty"` // Optional specific version
	Err     string `json:"err,omitempty"`
}

// SecretReadResponsej for getting secrets
type SecretReadResponse struct {
	Data map[string]string `json:"data"`
	Err  string            `json:"err,omitempty"`
}
