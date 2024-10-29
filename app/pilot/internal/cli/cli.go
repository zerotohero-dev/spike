//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/pilot/internal/cli/handle"
)

// TODO: document all public methods.

func Parse(source *workloadapi.X509Source, args []string) {
	if len(args) < 2 {
		handle.Usage(args)
		return
	}

	command := args[1]
	switch command {
	case "init":
		handle.Init(source, args)
	case "put":
		handle.Put(source, args)
	case "get":
		handle.Get(source, args)
	case "delete":
		handle.Delete(source, args)
	case "undelete":
		handle.Undelete(source, args)
	case "list":
		handle.List(source, args)
	default:
		handle.Default(source, args)
	}
}
