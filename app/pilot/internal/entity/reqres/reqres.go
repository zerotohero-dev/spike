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
}

// SecretWriteRequest for creating/updating secrets
type SecretWriteRequest struct {
	Options *SecretWriteOptions `json:"options,omitempty"`
	Data    map[string]string   `json:"data"`
}

type SecretWriteOptions struct {
	Cas int `json:"cas,omitempty"` // Compare-and-swap version number
}

// SecretWriteResponse after successful write
type SecretWriteResponse struct {
	SecretResponseMetadata
}

// ReadRequest for getting secrets (query params in URL)
type ReadRequest struct {
	Version int `json:"version,omitempty"` // Optional specific version
}

// SecretReadResponse for getting secrets
type SecretReadResponse struct {
	Data     map[string]string      `json:"data"`
	Metadata SecretResponseMetadata `json:"metadata"`
}

// SecretDeleteRequest for soft-deleting secret versions
type SecretDeleteRequest struct {
	Versions []int `json:"versions"` // Empty means latest version
}

// SecretDeleteResponse after soft-delete
type SecretDeleteResponse struct {
	Metadata SecretResponseMetadata `json:"metadata"`
}

// SecretDestroyRequest for permanently removing secret versions
type SecretDestroyRequest struct {
	Versions []int `json:"versions"` // Empty means latest version
}

// SecretDestroyResponse after permanent deletion
type SecretDestroyResponse struct {
	Metadata SecretResponseMetadata `json:"metadata"`
}

// SecretUndeleteRequest for recovering soft-deleted versions
type SecretUndeleteRequest struct {
	Versions []int `json:"versions"`
}

// SecretUndeleteResponse after recovery
type SecretUndeleteResponse struct {
	Metadata SecretResponseMetadata `json:"metadata"`
}

// SecretListResponse for listing secrets
type SecretListResponse struct {
	Keys []string `json:"keys"`
}

// SecretMetadataConfig represents the configuration for a path
type SecretMetadataConfig struct {
	MaxVersions        int               `json:"max_versions,omitempty"`
	CasRequired        bool              `json:"cas_required,omitempty"`
	DeleteVersionAfter time.Duration     `json:"delete_version_after,omitempty"`
	CustomMetadata     map[string]string `json:"custom_metadata,omitempty"`
}

// MetadataResponse when getting path metadata
type MetadataResponse struct {
	CreatedTime        time.Time         `json:"created_time"`
	CurrentVersion     int               `json:"current_version"`
	OldestVersion      int               `json:"oldest_version"`
	UpdatedTime        time.Time         `json:"updated_time"`
	MaxVersions        int               `json:"max_versions"`
	CasRequired        bool              `json:"cas_required"`
	DeleteVersionAfter time.Duration     `json:"delete_version_after"`
	CustomMetadata     map[string]string `json:"custom_metadata,omitempty"`
}

/*
API Endpoints and their request/response mappings:

1. Write Secret
POST /v2/data/{path}
Request:  SecretWriteRequest
Response: SecretWriteResponse

Example:
POST /v2/data/secret/foo
Request:
{
    "options": {
        "cas": 0
    },
    "data": {
        "username": "admin",
        "password": "secret123"
    }
}
Response:
{
    "created_time": "2024-01-01T12:00:00Z",
    "version": 1,
    "destroyed": false
}

2. Read Secret
GET /v2/data/{path}
Query Params: ?version=N
Response: SecretReadResponse

Example:
GET /v2/data/secret/foo?version=1
Response:
{
    "data": {
        "username": "admin",
        "password": "secret123"
    },
    "metadata": {
        "created_time": "2024-01-01T12:00:00Z",
        "version": 1,
        "destroyed": false
    }
}

3. Soft Delete
DELETE /v2/data/{path}
Request: SecretDeleteRequest
Response: SecretDeleteResponse

Example:
DELETE /v2/data/secret/foo
Request:
{
    "versions": [1, 2]
}
Response:
{
    "metadata": {
        "deleted_time": "2024-01-01T12:30:00Z",
        "destroyed": false,
        "version": 2
    }
}

4. Destroy
DELETE /v2/destroy/{path}
Request: SecretDestroyRequest
Response: SecretDestroyResponse

Example:
DELETE /v2/destroy/secret/foo
Request:
{
    "versions": [1]
}
Response:
{
    "metadata": {
        "destroyed": true,
        "version": 1
    }
}

5. Undelete
POST /v2/undelete/{path}
Request: SecretUndeleteRequest
Response: SecretUndeleteResponse

Example:
POST /v2/undelete/secret/foo
Request:
{
    "versions": [1, 2]
}
Response:
{
    "metadata": {
        "deleted_time": null,
        "destroyed": false,
        "version": 2
    }
}

6. List Secrets
LIST /v2/metadata/{path}
Response: SecretListResponse

Example:
LIST /v2/metadata/secret/
Response:
{
    "keys": ["foo", "bar"]
}

7. Read Metadata
GET /v2/metadata/{path}
Response: MetadataResponse

Example:
GET /v2/metadata/secret/foo
Response:
{
    "created_time": "2024-01-01T12:00:00Z",
    "current_version": 3,
    "oldest_version": 1,
    "updated_time": "2024-01-01T13:00:00Z",
    "max_versions": 10,
    "cas_required": false,
    "delete_version_after": "3600s",
    "custom_metadata": {
        "owner": "team-a"
    }
}

8. Update Metadata
POST /v2/metadata/{path}
Request: SecretMetadataConfig
Response: Empty (204 No Content)

Example:
POST /v2/metadata/secret/foo
Request:
{
    "max_versions": 5,
    "cas_required": true,
    "delete_version_after": "3600s",
    "custom_metadata": {
        "owner": "team-a"
    }
}
*/
