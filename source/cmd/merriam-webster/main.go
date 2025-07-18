package main

import (
	"github.com/example/merriam-webster/source/cmd/merriam-webster/cmd"
)

var version = "unknown"

func main() {
	cmd.Version = version
	cmd.Execute()
}
