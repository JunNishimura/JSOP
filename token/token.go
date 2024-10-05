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

var reservedWords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupStringTokenType(word string) TokenType {
	if tok, ok := reservedWords[word]; ok {
		return tok
	}
	return STRING
}

func IsBuiltinSymbol(strLiteral string) bool {
	if strLiteral == "+" ||
		strLiteral == "-" ||
		strLiteral == "*" ||
		strLiteral == "/" ||
		strLiteral == "==" ||
		strLiteral == "!=" ||
		strLiteral == ">" ||
		strLiteral == "<" ||
		strLiteral == ">=" ||
		strLiteral == "<=" ||
		strLiteral == "!" ||
		strLiteral == "at" {
		return true
	}

	return false
}
