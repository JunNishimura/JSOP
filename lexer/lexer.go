package lexer

import "github.com/JunNishimura/jsop/token"

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
		tok = newToken(token.COMMA, l.curChar)
	case '+':
		tok = newToken(token.PLUS, l.curChar)
	case '-':
		tok = newToken(token.MINUS, l.curChar)
	case '*':
		tok = newToken(token.ASTERISK, l.curChar)
	case '/':
		tok = newToken(token.SLASH, l.curChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.curChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.curChar) {
			tok.Literal = l.readString()
			tok.Type = token.LookupStringTokenType(tok.Literal)
			return tok
		}
		tok = newToken(token.ILLEGAL, l.curChar)
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.curChar == ' ' || l.curChar == '\t' || l.curChar == '\n' || l.curChar == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Lexer) readNumber() string {
	startPos := l.curPos
	for isDigit(l.curChar) {
		l.readChar()
	}
	return l.input[startPos:l.curPos]
}

func (l *Lexer) readString() string {
	startPos := l.curPos
	for isLetter(l.curChar) {
		l.readChar()
	}
	return l.input[startPos:l.curPos]
}
