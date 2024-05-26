//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/redjolr/goherent/cmd"
)

func TestWithLocalCmd() {
	fmt.Println(os.Args)
	extraCmdArgs := os.Args[2:]
	os.Exit(cmd.Main(extraCmdArgs))
}
