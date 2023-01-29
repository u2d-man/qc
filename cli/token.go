package cli

type Token int

const (
	ILLEGAL Token = iota

	// Whtie Space: ' '
	WS

	// *
	ASTERISK
	// ,
	COMMA

	SELECT
	FROM
	WHERE
)
