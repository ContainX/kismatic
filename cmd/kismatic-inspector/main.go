package main

import (
	"os"

	"github.com/apprenda/kismatic/pkg/inspector/cmd"
)

func main() {
	cmd := cmd.NewCmdKismaticInspector(os.Stdout)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
