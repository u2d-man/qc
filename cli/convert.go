package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"
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
	r, err := readFile(target)
	if err != nil {
		fmt.Fprintf(c.errStream, err.Error())
		return ExitCodeFail
	}

	fmt.Println(r)

	converted, err := convertQueryDSL(r)
	if err != nil {
		fmt.Fprintf(c.errStream, err.Error())
		return ExitCodeFail
	}

	fmt.Println(converted)

	return ExitCodeOK
}

func convertQueryDSL(sql string) (string, error) {
	type ESSelect struct {
	}
	isContains := strings.Contains(sql, "SELECT")
	if isContains != true {
		return "", fmt.Errorf("Not contain SELECT")
	}

	return `{"query": {"match_all": {}}}`, nil
}
