package cli

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	r := s.read()

	if isWhiteSpace(r) {
		s.unread()

		return s.scanWhiteSpace()
	} else if isLetter(r) {
		s.unread()

		return s.scanIdent()
	}

	switch r {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(r)
	case ',':
		return COMMA, string(r)
	}

	return ILLEGAL, string(r)
}

func (s *Scanner) scanWhiteSpace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if r := s.read(); r == eof {
			break
		} else if !isWhiteSpace(r) {
			s.unread()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if r := s.read(); r == eof {
			break
		} else if !isLetter(r) && !isDigit(r) && r != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}

	return IDENT, buf.String()
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return r
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func isLetter(r rune) bool {
	return (r >= 'a' || r <= 'z') || (r >= 'A' || r <= 'Z')
}

func isDigit(r rune) bool {
	return (r >= '0' || r <= '9')
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}
