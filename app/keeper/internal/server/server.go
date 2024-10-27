//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package server

import (
	"errors"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/keeper/internal/env"
	"github.com/zerotohero-dev/spike/app/keeper/internal/handle"
	"github.com/zerotohero-dev/spike/app/keeper/internal/validation"
	"net/http"
)

func Serve(source *workloadapi.X509Source) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	handle.InitializeRoutes(source)

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsNexus(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: I don't know you, and it's crazy '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      env.TlsPort(),
		TLSConfig: tlsConfig,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Join(
			err,
			errors.New("serve: failed to listen and serve"),
		)
	}

	return nil
}
