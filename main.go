package main

import (
	"github.com/strattonw/aug/cmd"
	"os"
)

func main() {
	if err := cmd.AugCommand.Execute(); err != nil {
		os.Exit(0)
	}
}
