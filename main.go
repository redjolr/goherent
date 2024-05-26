package main

import (
	"os"

	"github.com/redjolr/goherent/cmd"
)

func main() {
	extraCmdArgs := os.Args[1:]

	os.Exit(cmd.Main(extraCmdArgs))
}
