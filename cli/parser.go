package cli

import (
	"fmt"
	"io"
)

type SelectStatement struct {
	Fields          []string
	TableName       string
	SearchCondition []string
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

func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	if tok, lit := p.scanIgnoreWhitespace(); tok != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	for {
		tok, lit := p.scanIgnoreWhitespace()
		if tok != IDENT && tok != ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}
		stmt.Fields = append(stmt.Fields, lit)

		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	if tok, lit := p.scanIgnoreWhitespace(); tok == WHERE {
		for {
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT {
				return nil, fmt.Errorf("found %q, expected field", lit)
			}
			stmt.SearchCondition = append(stmt.SearchCondition, lit)

			if tok, _ := p.scanIgnoreWhitespace(); tok == SEMICOLON {
				break
			}
		}
	}

	return stmt, nil
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

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }
