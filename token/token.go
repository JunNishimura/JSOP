package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT     = "INT"
	SYMBOL  = "SYMBOL"
	COMMAND = "COMMAND"
	ARGS    = "ARGS"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	DOUBLE_QUOTE = "\""
	COLON        = ":"
	COMMA        = ","
)

var reservedWords = map[string]TokenType{
	"command": COMMAND,
	"args":    ARGS,
}

func LookupStringTokenType(word string) TokenType {
	if tok, ok := reservedWords[word]; ok {
		return tok
	}
	return SYMBOL
}
