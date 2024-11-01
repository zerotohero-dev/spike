//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/keeper/internal/route"
	"net/http"
)

func InitializeRoutes(source *workloadapi.X509Source) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement me!

		route.Route(r, w)
	})
}
