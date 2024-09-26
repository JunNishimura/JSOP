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
			name:  "1 digit integer",
			input: "1",
			expected: []token.Token{
				{Type: token.INT, Literal: "1"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "integer more than 1 digit",
			input: "123",
			expected: []token.Token{
				{Type: token.INT, Literal: "123"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "negative integer",
			input: "-123",
			expected: []token.Token{
				{Type: token.MINUS, Literal: "-"},
				{Type: token.INT, Literal: "123"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "true",
			input: "true",
			expected: []token.Token{
				{Type: token.TRUE, Literal: "true"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "false",
			input: "false",
			expected: []token.Token{
				{Type: token.FALSE, Literal: "false"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "exclamation",
			input: "!true",
			expected: []token.Token{
				{Type: token.EXCLAM, Literal: "!"},
				{Type: token.TRUE, Literal: "true"},
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "+"},
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "-"},
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "*"},
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "/"},
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
			name: "equation",
			input: `
				{
					"command": {
						"symbol": "==",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "=="},
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
			name: "nonequation",
			input: `
				{
					"command": {
						"symbol": "!=",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "!="},
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
			name: "less than expression",
			input: `
				{
					"command": {
						"symbol": "<",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "<"},
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
			name: "greater than expression",
			input: `
				{
					"command": {
						"symbol": ">",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: ">"},
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
			name: "less than or equal to expression",
			input: `
				{
					"command": {
						"symbol": "<=",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: "<="},
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
			name: "greater than or equal to expression",
			input: `
				{
					"command": {
						"symbol": ">=",
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
				{Type: token.SYMBOLKEY, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.SYMBOL, Literal: ">="},
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
			name: "if expression",
			input: `
				{
					"if": {
						"cond": true,
						"conseq": 1,
						"alt": 2
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.IF, Literal: "if"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COND, Literal: "cond"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.CONSEQ, Literal: "conseq"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.ALT, Literal: "alt"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "2"},
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
