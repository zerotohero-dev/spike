//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

import "time"

type SecretResponseMetadata struct {
	CreatedTime time.Time  `json:"created_time"`
	Destroyed   bool       `json:"destroyed"`
	Version     int        `json:"version"`
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	Err         string     `json:"err,omitempty"`
}

// SecretWriteRequest for creating/updating secrets
type SecretWriteRequest struct {
	Data map[string]string `json:"data"`
	Err  string            `json:"err,omitempty"`
}

// SecretWriteResponse after successful write
type SecretWriteResponse struct {
	SecretResponseMetadata
	Err string `json:"err,omitempty"`
}

// SecretReadRequest for getting secrets (query params in URL)
type SecretReadRequest struct {
	Version int    `json:"version,omitempty"` // Optional specific version
	Err     string `json:"err,omitempty"`
}

// SecretReadResponse for getting secrets
type SecretReadResponse struct {
	Data     map[string]string      `json:"data"`
	Metadata SecretResponseMetadata `json:"metadata"`
	Err      string                 `json:"err,omitempty"`
}

// SecretDeleteRequest for soft-deleting secret versions
type SecretDeleteRequest struct {
	Versions []int  `json:"versions"` // Empty means latest version
	Err      string `json:"err,omitempty"`
}

// SecretDeleteResponse after soft-delete
type SecretDeleteResponse struct {
	Metadata SecretResponseMetadata `json:"metadata"`
	Err      string                 `json:"err,omitempty"`
}

// SecretUndeleteRequest for recovering soft-deleted versions
type SecretUndeleteRequest struct {
	Versions []int  `json:"versions"`
	Err      string `json:"err,omitempty"`
}

// SecretUndeleteResponse after recovery
type SecretUndeleteResponse struct {
	Metadata SecretResponseMetadata `json:"metadata"`
	Err      string                 `json:"err,omitempty"`
}

// SecretListResponse for listing secrets
type SecretListResponse struct {
	Keys []string `json:"keys"`
	Err  string   `json:"err,omitempty"`
}
