//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package store

func (kv *KV) Get(path string, version int) (map[string]string, bool) {
	secret, exists := kv.data[path]
	if !exists {
		return nil, false
	}

	// If version not specified, use current version
	if version == 0 {
		version = secret.Metadata.CurrentVersion
	}

	v, exists := secret.Versions[version]
	if !exists || v.DeletedTime != nil {
		return nil, false
	}

	return v.Data, true
}
