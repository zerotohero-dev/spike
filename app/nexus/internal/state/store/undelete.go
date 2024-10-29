//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package store

// Undelete recovers soft-deleted versions
func (kv *KV) Undelete(path string, versions []int) error {
	secret, exists := kv.data[path]
	if !exists {
		return ErrSecretNotFound
	}

	for _, version := range versions {
		if v, exists := secret.Versions[version]; exists {
			v.DeletedTime = nil
			secret.Versions[version] = v
		}
	}

	return nil
}
