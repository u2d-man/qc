package main

import (
	"os"

	cmd "github.com/UserKazun/qc/cli"
)

func main() {
	cmd := cmd.NewCli(os.Stdout, os.Stderr)
	os.Exit(cmd.Execute(os.Args))
}
