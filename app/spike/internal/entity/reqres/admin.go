//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package reqres

type AdminTokenWriteRequest struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}

type AdminTokenWriteResponse struct {
	Err string `json:"err,omitempty"`
}