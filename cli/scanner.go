package cli

import (
	"bufio"
	"io"
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() {
	r := s.read()

	if isWhiteSpace(r) {

		return
	}
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return r
}

func isWhiteSpace(r rune) bool {
	return r == ' '
}
