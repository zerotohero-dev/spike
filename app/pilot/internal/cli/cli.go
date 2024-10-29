//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package cli

import "github.com/zerotohero-dev/spike/app/pilot/internal/cli/handle"

// TODO: document all public methods.

func Parse(args []string) {
	if len(args) < 2 {
		handle.Usage(args)
		return
	}

	command := args[1]
	switch command {
	case "init":
		handle.Init(args)
	case "put":
		handle.Put(args)
	case "get":
		handle.Get(args)
	case "delete":
		handle.Delete(args)
	case "undelete":
		handle.Undelete(args)
	case "list":
		handle.List(args)
	default:
		handle.Default(args)
	}
}
