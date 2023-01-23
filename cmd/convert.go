package cmd

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

	err := flags.Parse(args[1:])
	if err != nil {
		return ExitCodeParseFlagError
	}

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
	r, err := ReadFile(target)
	if err != nil {
		return ExitCodeFail
	}

	fmt.Println(string(r))

	return ExitCodeOK
}

func ReadFile(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("File open error %v", err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("File read error %v", err)
	}

	return string(b), nil
}
