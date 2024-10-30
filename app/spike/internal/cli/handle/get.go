//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/spike/app/spike/internal/net"
	"github.com/zerotohero-dev/spike/app/spike/internal/state"
	"strings"
)

func Get(source *workloadapi.X509Source, args []string) {

	if len(args) < 3 {
		fmt.Println("Usage: spike get <path> [-version=<n>]")
		return
	}

	// TODO: implement me.

	adminToken, err := state.AdminToken()
	if err != nil {
		fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
		fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
		fmt.Println(` \\\\\\\ web: spike.ist source: github.com/zerotohero-dev/spike`)
		fmt.Println("")
		fmt.Println("SPIKE is not initialized.")
		fmt.Println("Please run `spike init` to initialize SPIKE.")
		return
	}

	// TODO: for now we verify the admin token; later down the line, we will
	// exchange the admin token with a short-lived token with `spike login`.
	if adminToken == "" {
		fmt.Println(`    \\ SPIKE: Keep your secrets secret with SPIFFE.`)
		fmt.Println(`  \\\\\ Copyright 2024-present SPIKE contributors.`)
		fmt.Println(` \\\\\\\ web: spike.ist source: github.com/zerotohero-dev/spike`)
		fmt.Println("")
		fmt.Println("SPIKE is not initialized.")
		fmt.Println("Please run `spike init` to initialize SPIKE.")
		return
	}

	// TODO: we need a more capable arg parser.
	version := 0
	path := args[2]
	if len(args) > 3 && strings.HasPrefix(args[3], "-version=") {
		fmt.Sscanf(args[3], "-version=%d", &version)
	}

	secret, err := net.GetSecret(source, path, version)
	if err != nil {
		fmt.Println("Error reading secret:", err.Error())
		return
	}

	data := secret.Data
	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}

/*
func parseVersionFlag(arg string) (int, error) {
    var version int
    if strings.HasPrefix(arg, "-version=") {
        _, err := fmt.Sscanf(arg, "-version=%d", &version)
        return version, err
    }
    return 0, fmt.Errorf("invalid version flag: %s", arg)
}

func parseVersionsFlag(arg string) ([]int, error) {
    if !strings.HasPrefix(arg, "-versions=") {
        return nil, fmt.Errorf("invalid versions flag: %s", arg)
    }
    versions := strings.Split(strings.TrimPrefix(arg, "-versions="), ",")
    result := make([]int, 0, len(versions))
    for _, v := range versions {
        n, err := strconv.Atoi(v)
        if err != nil {
            return nil, fmt.Errorf("invalid version number: %s", v)
        }
        result = append(result, n)
    }
    return result, nil
}

*/
