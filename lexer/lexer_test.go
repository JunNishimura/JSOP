package lexer

import (
	"testing"

	"github.com/JunNishimura/jsop/token"
)

func TestSingleProgram(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name: "integer atom",
			input: `
				{
					"atom": 1
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ATOM, Literal: "atom"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "1"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "integer more than 1 digit",
			input: `
				{
					"atom": 123
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ATOM, Literal: "atom"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "123"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "integer with plus sign",
			input: `
				{
					"atom": +123
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ATOM, Literal: "atom"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "123"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "negative integer",
			input: `
				{
					"atom": -123
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ATOM, Literal: "atom"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "-123"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "mathematical operation: addition",
			input: `
				{
					"command": {
						"symbol": "+",
						"args": [1, 2]
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMAND, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ARGS, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "mathematical operation: subtraction",
			input: `
				{
					"command": {
						"symbol": "-",
						"args": [1, 2]
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMAND, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ARGS, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "mathematical operation: multiplication",
			input: `
				{
					"command": {
						"symbol": "*",
						"args": [1, 2]
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMAND, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ARGS, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "mathematical operation: division",
			input: `
				{
					"command": {
						"symbol": "/",
						"args": [1, 2]
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMAND, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ARGS, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			for i, expected := range tt.expected {
				tok := l.NextToken()
				if tok.Type != expected.Type {
					t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expected.Type, tok.Type)
				}
				if tok.Literal != expected.Literal {
					t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expected.Literal, tok.Literal)
				}
			}
		})
	}
}
