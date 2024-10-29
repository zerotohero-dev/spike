//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package store

import "time"

func (kv *KV) Delete(path string, versions []int) {
	secret, exists := kv.data[path]
	if !exists {
		return
	}

	now := time.Now()

	// If no versions specified, mark the latest version as deleted
	if len(versions) == 0 {
		if v, exists := secret.Versions[secret.Metadata.CurrentVersion]; exists {
			v.DeletedTime = &now
			secret.Versions[secret.Metadata.CurrentVersion] = v
		}
		return
	}

	// Delete specific versions
	for _, version := range versions {
		if v, exists := secret.Versions[version]; exists {
			v.DeletedTime = &now
			secret.Versions[version] = v
		}
	}
}
