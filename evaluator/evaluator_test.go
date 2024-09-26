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

func testBooleanObject(t *testing.T, obj object.Object, expected bool) {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Fatalf("object is not Boolean. got=%T", obj)
	}

	if result.Value != expected {
		t.Fatalf("object has wrong value. got=%t, want=%t", result.Value, expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "true",
			input:    "true",
			expected: true,
		},
		{
			name:     "false",
			input:    "false",
			expected: false,
		},
		{
			name:     "negation of true",
			input:    "!true",
			expected: false,
		},
		{
			name:     "negation of false",
			input:    "!false",
			expected: true,
		},
		{
			name: "equation symbol: return true",
			input: `
				{
					"command": {
						"symbol": "==",
						"args": [true, true]
					}
				}`,
			expected: true,
		},
		{
			name: "equation symbol: return false",
			input: `
				{
					"command": {
						"symbol": "==",
						"args": [true, false]
					}
				}`,
			expected: false,
		},
		{
			name: "inequation symbol: return true",
			input: `
				{
					"command": {
						"symbol": "!=",
						"args": [true, false]
					}
				}`,
			expected: true,
		},
		{
			name: "inequation symbol: return false",
			input: `
				{
					"command": {
						"symbol": "!=",
						"args": [true, true]
					}
				}`,
			expected: false,
		},
		{
			name: "comparison symbol(greater than): return true",
			input: `
				{
					"command": {
						"symbol": ">",
						"args": [2, 1]
					}
				}`,
			expected: true,
		},
		{
			name: "comparison symbol(greater than): return false",
			input: `
				{
					"command": {
						"symbol": ">",
						"args": [1, 2]
					}
				}`,
			expected: false,
		},
		{
			name: "comparison symbol(less than): return true",
			input: `
				{
					"command": {
						"symbol": "<",
						"args": [1, 2]
					}
				}`,
			expected: true,
		},
		{
			name: "comparison symbol(less than): return false",
			input: `
				{
					"command": {
						"symbol": "<",
						"args": [2, 1]
					}
				}`,
			expected: false,
		},
		{
			name: "comparison symbol(greater than or equal to): return true",
			input: `
				{
					"command": {
						"symbol": ">=",
						"args": [2, 1]
					}
				}`,
			expected: true,
		},
		{
			name: "comparison symbol(greater than or equal to): return false",
			input: `
				{
					"command": {
						"symbol": ">=",
						"args": [1, 2]
					}
				}`,
			expected: false,
		},
		{
			name: "comparison symbol(less than or equal to): return true",
			input: `
				{
					"command": {
						"symbol": "<=",
						"args": [1, 2]
					}
				}`,
			expected: true,
		},
		{
			name: "comparison symbol(less than or equal to): return false",
			input: `
				{
					"command": {
						"symbol": "<=",
						"args": [2, 1]
					}
				}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(t, tt.input)

			testBooleanObject(t, evaluated, tt.expected)
		})
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name: "if the condition is true with literal boolean",
			input: `
				{
					"if": {
						"cond": true,
						"conseq": 10
					}
				}`,
			expected: 10,
		},
		{
			name: "if the condition is false with literal boolean",
			input: `
				{
					"if": {
						"cond": false,
						"conseq": 10,
						"alt": 20
					}
				}`,
			expected: 20,
		},
		{
			name: "if expression with object",
			input: `
				{
					"if": {
						"cond": {
							"command": {
								"symbol": ">",
								"args": [2, 1]
							}
						},
						"conseq": {
							"command": {
								"symbol": "+",
								"args": [1, 2]
							}
						},
						"alt": {
							"command": {
								"symbol": "-",
								"args": [1, 2]
							}
						}
					}
				}`,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(t, tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}
