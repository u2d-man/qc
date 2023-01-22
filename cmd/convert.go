package cmd

import (
	"flag"
	"io"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
	ExitCodeFail           = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func NewCli(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Execute(args []string) int {
	var filename string

	flags := flag.NewFlagSet("qc", flag.ExitOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&filename, "filename", "", "Allowed extensions: .sql")

	flag.Parse()

	argv := flags.Args()
	target := ""
	if len(argv) == 1 {
		target = argv[0]
	} else {
		return ExitCodeParseFlagError
	}

	return c.run(target)
}

func (c *CLI) run(target string) int {
	return ExitCodeOK
}
