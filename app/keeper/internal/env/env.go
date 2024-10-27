//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package env

func TlsPort() string {
	return ":8443"

	// TODO: make dynamic:
	//p := env.Value(env.VSecMSafeTlsPort)
	//if p == "" {
	//	p = string(env.VSecMSafeTlsPortDefault)
	//}
	//return p
}
