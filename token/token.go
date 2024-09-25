package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT       = "INT"
	SYMBOLKEY = "SYMBOL_KEY"
	COMMAND   = "COMMAND"
	ARGS      = "ARGS"
	IF        = "IF"
	COND      = "COND"
	CONSEQ    = "CONSEQ"
	ALT       = "ALT"

	SYMBOL = "SYMBOL"

	MINUS  = "-"
	EXCLAM = "!"

	TRUE  = "true"
	FALSE = "false"

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
	"symbol":  SYMBOLKEY,
	"args":    ARGS,
	"if":      IF,
	"cond":    COND,
	"conseq":  CONSEQ,
	"alt":     ALT,
	"true":    TRUE,
	"false":   FALSE,
}

func LookupStringTokenType(word string) TokenType {
	if tok, ok := reservedWords[word]; ok {
		return tok
	}
	return SYMBOL
}
