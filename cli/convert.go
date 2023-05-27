package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
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

	flags.StringVar(&filename, "f", "", "Allowed extensions: .sql")

	err := flags.Parse(args[1:])
	if err != nil {
		return ExitCodeParseFlagError
	}

	return c.run(filename)
}

func (c *CLI) run(filename string) int {
	r, err := readFile(filename)
	if err != nil {
		_, _ = fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeFail
	}

	parser := NewParser(strings.NewReader(r))
	stmt, err := parser.Parse()
	if err != nil {
		_, _ = fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeFail
	}

	converted, err := c.convertToQueryDSL(stmt)
	if err != nil {
		_, _ = fmt.Fprintln(c.errStream, err.Error())
		return ExitCodeFail
	}

	fmt.Println(converted)

	return ExitCodeOK
}

func readFile(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("file open error: %v", err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	b, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("file read error: %v", err)
	}

	return string(b), nil
}

func (c *CLI) convertToQueryDSL(stmt *SelectStatement) (string, error) {
	// queryDSL := QueryDSL{}
	queryMap := c.generateQueryMap(stmt)
	if stmt != nil {
		marshaled, err := json.Marshal(queryMap)
		if err != nil {
			return "", fmt.Errorf("cannot marshal: %w", err)
		}

		return string(marshaled), nil
	}

	return "", nil
}

func (c *CLI) generateQueryMap(stmt *SelectStatement) map[string]interface{} {
	var queryDSL map[string]interface{}
	if stmt != nil && len(stmt.Fields) != 0 {
		queryDSL = map[string]interface{}{
			"_source": stmt.Fields,
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}
	}

	return queryDSL
}
