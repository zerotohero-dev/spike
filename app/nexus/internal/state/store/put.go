//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package store

import "time"

// TODO: make store access thread-safe.

func (kv *KV) Put(path string, values map[string]string) {
	rightNow := time.Now()

	secret, exists := kv.data[path]
	if !exists {
		secret = &Secret{
			Versions: make(map[int]Version),
			Metadata: Metadata{
				CreatedTime:    rightNow,
				UpdatedTime:    rightNow,
				MaxVersions:    3,
				CurrentVersion: 0,
				OldestVersion:  0,
			},
		}
		kv.data[path] = secret
	}

	// Increment version
	newVersion := secret.Metadata.CurrentVersion + 1

	// Add new version
	secret.Versions[newVersion] = Version{
		Data:        values,
		CreatedTime: rightNow,
		Version:     newVersion,
	}

	// Update metadata
	secret.Metadata.CurrentVersion = newVersion
	secret.Metadata.UpdatedTime = rightNow
	if secret.Metadata.OldestVersion == 0 {
		secret.Metadata.OldestVersion = 1
	}

	// Cleanup old versions if exceeding MaxVersions
	for version := range secret.Versions {
		if secret.Metadata.CurrentVersion-version >= secret.Metadata.MaxVersions {
			delete(secret.Versions, version)
			if version == secret.Metadata.OldestVersion {
				secret.Metadata.OldestVersion = version + 1
			}
		}
	}
}
