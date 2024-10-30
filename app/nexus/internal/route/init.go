//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zerotohero-dev/spike/app/nexus/internal/entity/v1/reqres"
	"github.com/zerotohero-dev/spike/app/nexus/internal/state"
)

func routeInit(r *http.Request, w http.ResponseWriter) {
	fmt.Println("routeInit:", r.Method, r.URL.Path, r.URL.RawQuery)

	w.WriteHeader(http.StatusOK)

	body, err := body(r)
	if err != nil {
		log.Println("routeInit: Problem reading request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routeInit: Problem writing response:", err.Error())
		}
		return
	}
	if body == nil {
		log.Println("routeInit: No request body.")
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routeInit: Problem writing response:", err.Error())
		}
		return
	}

	var req reqres.AdminTokenWriteRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("routeInit: Problem parsing request body:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("routeInit: Problem writing response:", err.Error())
		}
		return
	}

	adminToken := req.Data

	// TODO: do not log this!
	// TODO: save this in memory
	// TODO: sync this with keep
	// TODO: if there is an admin token in memory; reject updating it.

	// TODO: anyone has access to the machine has access to admin token (because it's saved on the file system)
	// We can keep it in a folder and configure file/folder access permissions to reduce the impact,
	// but still, if the machine is compromised, the token is compromised too.
	// as an added security measure, the user can opt-in to encrypt the config.

	// TODO: provide ability for the admin to rotate the root key

	// log.Println("Received admin token:", adminToken)
	// TODO: sanitize admin token.
	state.SetAdminToken(adminToken)
	log.Println("routeInit: Admin token saved")

	// TODO: sync admin token with keep.

	_, err = io.WriteString(w, "")
	if err != nil {
		log.Println("routeInit: Problem writing response:", err.Error())
	}
}
