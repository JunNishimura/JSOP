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
	SYMBOL = "SYMBOL"

	PLUS  = "+"
	MINUS = "-"

	LBRACE       = "{"
	RBRACE       = "}"
	DOUBLE_QUOTE = "\""
	COLON        = ":"

	// reserved tokens
	ATOM = "ATOM"
)

var reservedWords = map[string]TokenType{
	"atom": ATOM,
}

func LookupStringTokenType(word string) TokenType {
	if tok, ok := reservedWords[word]; ok {
		return tok
	}
	return SYMBOL
}
