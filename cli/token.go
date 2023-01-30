package cli

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	IDENT

	ASTERISK // *
	COMMA    // ,

	// Keywords
	SELECT
	FROM
	WHERE
)
