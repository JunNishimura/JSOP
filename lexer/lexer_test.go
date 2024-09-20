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
