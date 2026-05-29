// SPDX-License-Identifier: MIT

package main

import "github.com/jojin1709/403-bypass/cmd"

// Version and BuildDate are set via ldflags at build time.
var (
	Version   = "dev"
	BuildDate = "unknown"
)

func main() {
	cmd.SetVersionInfo(Version, BuildDate)
	cmd.Execute()
}
