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
	STRING    = "STRING"
	SYMBOLKEY = "SYMBOL_KEY"
	COMMAND   = "COMMAND"
	ARGS      = "ARGS"
	IF        = "IF"
	COND      = "COND"
	CONSEQ    = "CONSEQ"
	ALT       = "ALT"
	SET       = "SET"
	VAR       = "VAR"
	VAL       = "VAL"
	LOOP      = "LOOP"
	FOR       = "FOR"
	FROM      = "FROM"
	TO        = "TO"
	DO        = "DO"

	SYMBOL = "SYMBOL"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	LTE      = "<="
	GT       = ">"
	GTE      = ">="
	EQ       = "=="
	NOT_EQ   = "!="
	EXCLAM   = "!"

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
	"set":     SET,
	"var":     VAR,
	"val":     VAL,
	"loop":    LOOP,
	"for":     FOR,
	"from":    FROM,
	"to":      TO,
	"do":      DO,
}

func LookupStringTokenType(word string) TokenType {
	if tok, ok := reservedWords[word]; ok {
		return tok
	}
	return STRING
}
