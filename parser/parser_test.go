package parser

import (
	"testing"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/token"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIntegerAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "1 digit integer",
			input:    "1",
			expected: 1,
		},
		{
			name:     "integer more than 1 digit",
			input:    "123",
			expected: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			checkParserErrors(t, p)

			program := p.ParseProgram()
			if len(program.Expressions) != 1 {
				t.Fatalf("program.Expressions does not contain 1 expression. got=%d", len(program.Expressions))
			}

			intAtom, ok := program.Expressions[0].(*ast.IntegerLiteral)
			if !ok {
				t.Fatalf("exp not *ast.Integer. got=%T", program.Expressions[0])
			}
			if intAtom.Value != tt.expected {
				t.Fatalf("intAtom.Value not %d. got=%d", tt.expected, intAtom.Value)
			}
		})
	}
}

func TestPrefixAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.PrefixAtom
	}{
		{
			name:  "negative integer",
			input: "-123",
			expected: &ast.PrefixAtom{
				Token:    token.Token{Type: token.MINUS, Literal: "-"},
				Operator: "-",
				Right: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "123"},
					Value: 123,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			checkParserErrors(t, p)

			program := p.ParseProgram()
			if len(program.Expressions) != 1 {
				t.Fatalf("program.Expressions does not contain 1 expression. got=%d", len(program.Expressions))
			}

			prefixAtom, ok := program.Expressions[0].(*ast.PrefixAtom)
			if !ok {
				t.Fatalf("exp not *ast.PrefixAtom. got=%T", program.Expressions[0])
			}
			if prefixAtom.String() != tt.expected.String() {
				t.Fatalf("prefixAtom.String() not %q. got=%q", tt.expected.String(), prefixAtom.String())
			}
		})
	}
}
