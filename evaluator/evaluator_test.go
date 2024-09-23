package evaluator

import (
	"testing"

	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/object"
	"github.com/JunNishimura/jsop/parser"
)

func testEval(t *testing.T, input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program, err := p.ParseProgram()
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("object is not Integer. got=%T", obj)
	}

	if result.Value != expected {
		t.Fatalf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "1 digit integer",
			input:    "5",
			expected: 5,
		},
		{
			name:     "more than 1 digit integer",
			input:    "10",
			expected: 10,
		},
		{
			name:     "negative integer",
			input:    "-5",
			expected: -5,
		},
		{
			name: "addition",
			input: `
				{
					"command": {
						"symbol": "+",
						"args": [1, 2]
					}
				}
			`,
			expected: 3,
		},
		{
			name: "subtraction",
			input: `
				{
					"command": {
						"symbol": "-",
						"args": [1, 2]
					}
				}
			`,
			expected: -1,
		},
		{
			name: "multiplication",
			input: `
				{
					"command": {
						"symbol": "*",
						"args": [1, 2]
					}
				}
			`,
			expected: 2,
		},
		{
			name: "division",
			input: `
				{
					"command": {
						"symbol": "/",
						"args": [4, 2]
					}
				}
			`,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(t, tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}
