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
			name:  "string literal",
			input: `"hello"`,
			expected: []token.Token{
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "hello"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "array",
			input: "[1, 2]",
			expected: []token.Token{
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "+"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "-"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "*"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "/"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "=="},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "!="},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "<"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: ">"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "<="},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: ">="},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
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
			name: "negation expression",
			input: `
				{
					"command": {
						"symbol": "!",
						"args": true
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "!"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.TRUE, Literal: "true"},
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
				{Type: token.STRING, Literal: "if"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "cond"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "conseq"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "alt"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "set expression",
			input: `
				{
					"set": {
						"var": "$x",
						"val": 1
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "set"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "var"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$x"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "val"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "1"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "lambda expression",
			input: `
				{
					"lambda": {
						"params": ["$x", "$y"],
						"body": {
							"command": {
								"symbol": "+",
								"args": ["$x", "$y"]
							}
						}
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "lambda"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "params"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$x"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$y"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "body"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "+"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$x"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$y"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "loop expression",
			input: `
				{
					"loop": {
						"for": "$i",
						"from": 0,
						"to": 10,
						"do": {
							"command": {
								"symbol": "==",
								"args": ["$i", 5]
							}
						}
					}
				}`,
			expected: []token.Token{
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "loop"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "for"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$i"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "from"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "0"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "to"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.INT, Literal: "10"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "do"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "=="},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$i"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "5"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
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

func TestMultiplePrograms(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name: "multiple atoms",
			input: `
				[
					1,
					2
				]
			`,
			expected: []token.Token{
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "multiple commands",
			input: `
				[
					{
						"command": {
							"symbol": "+",
							"args": [1, 2]
						}
					},
					{
						"command": {
							"symbol": "-",
							"args": [3, 4]
						}
					}
				]
			`,
			expected: []token.Token{
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "+"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "-"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "3"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "4"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name: "read symbol from environment",
			input: `
				[
					{
						"set": {
							"var": "$x",
							"val": {
								"command": {
									"symbol": "+",
									"args": [1, 2]
								}
							}
						}
					},
					"$x"
				]
			`,
			expected: []token.Token{
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "set"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "var"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$x"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "val"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "command"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "symbol"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "+"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "args"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.COLON, Literal: ":"},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "1"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "2"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.STRING, Literal: "$x"},
				{Type: token.DOUBLE_QUOTE, Literal: "\""},
				{Type: token.RBRACKET, Literal: "]"},
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
