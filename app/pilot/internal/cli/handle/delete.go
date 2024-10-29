//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package handle

//// parseVersions helper function to parse version numbers from command args
//func parseVersions(args []string) []int {
//	versions := []int{}
//	for _, arg := range args {
//		if strings.HasPrefix(arg, "-versions=") {
//			versionStr := strings.TrimPrefix(arg, "-versions=")
//			for _, v := range strings.Split(versionStr, ",") {
//				var version int
//				fmt.Sscanf(v, "%d", &version)
//				versions = append(versions, version)
//			}
//			break
//		}
//	}
//	return versions
//}

func Delete(args []string) {
	panic("handleDelete not implemented")

	//		if len(args) < 3 {
	//			fmt.Println("Usage: pilot delete <path> [-versions=<n1,n2,...>]")
	//			return
	//		}
	//		versions := parseVersions(args)
	//		if err := store.softDelete(args[2], versions); err != nil {
	//			fmt.Printf("Error: %v\n", err)
	//			return
	//		}
	//		fmt.Printf("Success! Versions marked as deleted at: %s\n", args[2])
}
