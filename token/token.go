package token

import "strings"

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

func IsBuiltinSymbol(strLiteral string) bool {
	trimmedStr := strings.TrimSpace(strLiteral)

	if trimmedStr == "+" ||
		trimmedStr == "-" ||
		trimmedStr == "*" ||
		trimmedStr == "/" ||
		trimmedStr == "==" ||
		trimmedStr == "!=" ||
		trimmedStr == ">" ||
		trimmedStr == "<" ||
		trimmedStr == ">=" ||
		trimmedStr == "<=" ||
		trimmedStr == "!" ||
		trimmedStr == "at" ||
		trimmedStr == "print" ||
		trimmedStr == "len" {
		return true
	}

	return false
}

func IsSymbol(strLiteral string) bool {
	trimmedStr := strings.TrimSpace(strLiteral)

	// not contain any whitespace
	if strings.Contains(trimmedStr, " ") {
		return false
	}

	// starts with a dollar sign
	return strings.HasPrefix(trimmedStr, "$")
}
