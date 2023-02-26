package cli

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	IDENT // fields, table name

	ASTERISK // *
	COMMA    // ,
	EQUAL    // =

	// Keywords
	SELECT
	FROM
	WHERE
)
