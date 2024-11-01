//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"errors"
	"fmt"
	"github.com/zerotohero-dev/spike/internal/config"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

// CreateMtlsServer creates an HTTP server configured for mutual TLS (mTLS)
// authentication using SPIFFE X.509 certificates. It sets up the server with a
// custom authorizer that validates client SPIFFE IDs against a provided
// predicate function.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. It must be initialized and valid.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//   - predicate: A function that takes a SPIFFE ID string and returns true if
//     the client should be allowed access, false otherwise.
//
// Returns:
//   - *http.Server: A configured HTTP server ready to be started with TLS
//     enabled.
//   - error: An error if the server configuration fails.
//
// The server uses the provided X509Source for both its own identity and for
// validating client certificates. Client connections are only accepted if their
// SPIFFE ID passes the provided predicate function.
func CreateMtlsServer(source *workloadapi.X509Source,
	tlsPort string,
	predicate func(string) bool) (*http.Server, error) {
	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if predicate(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: I don't know you, and it's crazy '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      tlsPort,
		TLSConfig: tlsConfig,
	}
	return server, nil
}

// Serve initializes and starts an HTTPS server using mTLS authentication with
// SPIFFE X.509 certificates. It sets up the server routes using the provided
// initialization function and listens for incoming connections on the specified
// port.
//
// Parameters:
//   - source: An X509Source that provides the server's identity credentials and
//     validates client certificates. Must not be nil.
//   - initializeRoutes: A function that sets up the HTTP route handlers for the
//     server. This function is called before the server starts.
//   - tlsPort: The network address and port for the server to listen on
//     (e.g., ":8443").
//
// Returns:
//   - error: Returns nil if the server starts successfully, otherwise returns
//     an error explaining the failure. Specific error cases include:
//   - If source is nil
//   - If server creation fails
//   - If the server fails to start or encounters an error while running
//
// The function uses empty strings for the certificate and key file parameters
// in ListenAndServeTLS as the certificates are provided by the X509Source. The
// server's mTLS configuration is determined by the CreateMtlsServer function.
func Serve(source *workloadapi.X509Source,
	initializeRoutes func(), tlsPort string) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	initializeRoutes()

	server, err := CreateMtlsServer(source, tlsPort, config.IsNexus)
	if err != nil {
		return err
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Join(
			err,
			errors.New("serve: failed to listen and serve"),
		)
	}

	return nil
}
