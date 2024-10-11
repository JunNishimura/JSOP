package lexer

import (
	"strings"

	"github.com/JunNishimura/jsop/token"
)

type Lexer struct {
	input   string
	curPos  int
	nextPos int
	curChar byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.curChar = 0
	} else {
		l.curChar = l.input[l.nextPos]
	}
	l.curPos = l.nextPos
	l.nextPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.curChar {
	case '{':
		tok = newToken(token.LBRACE, l.curChar)
	case '}':
		tok = newToken(token.RBRACE, l.curChar)
	case '[':
		tok = newToken(token.LBRACKET, l.curChar)
	case ']':
		tok = newToken(token.RBRACKET, l.curChar)
	case '"':
		tok = newToken(token.DOUBLE_QUOTE, l.curChar)
	case ':':
		tok = newToken(token.COLON, l.curChar)
	case ',':
		if isLetter(l.peekChar()) {
			l.readChar()
			literal := "," + l.readString(isLetter, isSpecialChar)
			return token.Token{
				Type:    token.STRING,
				Literal: literal,
			}
		}
		tok = newToken(token.COMMA, l.curChar)
	case '-':
		if l.peekChar() == '"' {
			tok = newToken(token.STRING, l.curChar)
		} else {
			tok = newToken(token.MINUS, l.curChar)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.curChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.curChar) || isSpecialChar(l.curChar) {
			readStr := l.readString(isLetter, isSpecialChar, isWhitespace)

			trimmedStr := strings.TrimSpace(readStr)
			if trimmedStr == "true" {
				return token.Token{Type: token.TRUE, Literal: trimmedStr}
			} else if trimmedStr == "false" {
				return token.Token{Type: token.FALSE, Literal: trimmedStr}
			}

			return token.Token{Type: token.STRING, Literal: readStr}
		}
		tok = newToken(token.ILLEGAL, l.curChar)
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.curChar) {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isSpecialChar(ch byte) bool {
	return ch == '=' ||
		ch == '+' ||
		ch == '-' ||
		ch == '*' ||
		ch == '/' ||
		ch == '<' ||
		ch == '>' ||
		ch == '!' ||
		ch == '$'
}

func (l *Lexer) readNumber() string {
	startPos := l.curPos
	for isDigit(l.curChar) {
		l.readChar()
	}
	return l.input[startPos:l.curPos]
}

func (l *Lexer) readString(filters ...func(byte) bool) string {
	startPos := l.curPos
	for {
	LOOP:
		for _, filter := range filters {
			if filter(l.curChar) {
				l.readChar()
				goto LOOP
			}
		}
		break
	}
	return l.input[startPos:l.curPos]
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}
