//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/zerotohero-dev/spike/app/nexus/internal/route"
)

func InitializeRoutes(source *workloadapi.X509Source) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement me!

		//cid := crypto.Id()
		//
		//validation.EnsureSafe(source)
		//
		//id, err := s.IdFromRequest(r)
		//
		//if err != nil {
		//	log.WarnLn(&cid, "Handler: blocking insecure svid", id, err)
		//
		//	routeFallback.Fallback(cid, r, w)
		//
		//	return
		//}
		//
		//sid := s.IdAsString(r)
		//

		route.Route(r, w)
	})
}
