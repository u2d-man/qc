package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type SQL struct {
	Fields    []string
	TableName string
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*SQL, error) {
	stmt := &SQL{}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	for {
		tok, lit := p.scanIgnoreWhiteSpace()
		if tok != IDENT && tok != ASTERISK {
			return nil, fmt.Errorf("found %q, expected fields", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)

		if tok, _ := p.scanIgnoreWhiteSpace(); tok != COMMA {
			p.s.unread()
			break
		}
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	tok, lit := p.scanIgnoreWhiteSpace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	return stmt, nil
}

func readFile(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("File open error: %v", err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("File read error: %v", err)
	}

	return string(b), nil
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreWhiteSpace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}

	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}
