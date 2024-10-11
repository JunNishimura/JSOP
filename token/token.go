package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT    = "INT"
	STRING = "STRING"

	MINUS = "-"

	TRUE  = "TRUE"
	FALSE = "false"

	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	DOUBLE_QUOTE = "\""
	COLON        = ":"
	COMMA        = ","
)
