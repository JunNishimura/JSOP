package parser

import (
	"testing"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/token"
)

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

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			intAtom, ok := program.(*ast.IntegerLiteral)
			if !ok {
				t.Fatalf("exp not *ast.Integer. got=%T", program)
			}
			if intAtom.Value != tt.expected {
				t.Fatalf("intAtom.Value not %d. got=%d", tt.expected, intAtom.Value)
			}
		})
	}
}

func TestStringAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "string with double quotes",
			input:    `"hello"`,
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			strAtom, ok := program.(*ast.StringLiteral)
			if !ok {
				t.Fatalf("exp not *ast.String. got=%T", program)
			}
			if strAtom.Value != tt.expected {
				t.Fatalf("strAtom.Value not %q. got=%q", tt.expected, strAtom.Value)
			}
		})
	}
}

func TestBooleanAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.Boolean
	}{
		{
			name:  "true",
			input: "true",
			expected: &ast.Boolean{
				Token: token.Token{Type: token.TRUE, Literal: "true"},
				Value: true,
			},
		},
		{
			name:  "false",
			input: "false",
			expected: &ast.Boolean{
				Token: token.Token{Type: token.FALSE, Literal: "false"},
				Value: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			boolean, ok := program.(*ast.Boolean)
			if !ok {
				t.Fatalf("exp not *ast.Boolean. got=%T", program)
			}
			if boolean.String() != tt.expected.String() {
				t.Fatalf("boolean.String() not %q. got=%q", tt.expected.String(), boolean.String())
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

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			prefixAtom, ok := program.(*ast.PrefixAtom)
			if !ok {
				t.Fatalf("exp not *ast.PrefixAtom. got=%T", program)
			}
			if prefixAtom.String() != tt.expected.String() {
				t.Fatalf("prefixAtom.String() not %q. got=%q", tt.expected.String(), prefixAtom.String())
			}
		})
	}
}

func TestArrayAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.Array
	}{
		{
			name:  "empty array",
			input: "[]",
			expected: &ast.Array{
				Token:    token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{},
			},
		},
		{
			name:  "array with 1 element",
			input: "[1]",
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
				},
			},
		},
		{
			name:  "array with multiple elements",
			input: "[1, 2, 3]",
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
					&ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "2"},
						Value: 2,
					},
					&ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "3"},
						Value: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			array, ok := program.(*ast.Array)
			if !ok {
				t.Fatalf("exp not *ast.Array. got=%T", program)
			}
			if array.String() != tt.expected.String() {
				t.Fatalf("array.String() not %q. got=%q", tt.expected.String(), array.String())
			}
		})
	}
}

func TestCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.CommandObject
	}{
		{
			name: "addition command",
			input: `
				{
					"command": {
						"symbol": "+",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.PLUS, Literal: "+"},
					Value: "+",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "subtraction command",
			input: `
				{
					"command": {
						"symbol": "-",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.MINUS, Literal: "-"},
					Value: "-",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "multiplication command",
			input: `
				{
					"command": {
						"symbol": "*",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.ASTERISK, Literal: "*"},
					Value: "*",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "division command",
			input: `
				{
					"command": {
						"symbol": "/",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.SLASH, Literal: "/"},
					Value: "/",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "equation command",
			input: `
				{
					"command": {
						"symbol": "==",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.EQ, Literal: "=="},
					Value: "==",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "negation of equation command",
			input: `
				{
					"command": {
						"symbol": "!=",
						"args": [1, 2]
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.NOT_EQ, Literal: "!="},
					Value: "!=",
				},
				Args: &ast.Array{
					Token: token.Token{Type: token.LBRACKET, Literal: "["},
					Elements: []ast.Expression{
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						&ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "2"},
							Value: 2,
						},
					},
				},
			},
		},
		{
			name: "negation of true",
			input: `
				{
					"command": {
						"symbol": "!",
						"args": true
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.EXCLAM, Literal: "!"},
					Value: "!",
				},
				Args: &ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
		},
		{
			name: "no arguments command",
			input: `
				{
					"command": {
						"symbol": "$hoge"
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.SYMBOL, Literal: "$hoge"},
					Value: "$hoge",
				},
				Args: nil,
			},
		},
		{
			name: "only one argument command",
			input: `
				{
					"command": {
						"symbol": "$hoge",
						"args": 1
					}
				}`,
			expected: &ast.CommandObject{
				Token: token.Token{Type: token.COMMAND, Literal: "command"},
				Symbol: &ast.Symbol{
					Token: token.Token{Type: token.SYMBOL, Literal: "$hoge"},
					Value: "$hoge",
				},
				Args: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			command, ok := program.(*ast.CommandObject)
			if !ok {
				t.Fatalf("exp not *ast.Command. got=%T", program)
			}
			if command.String() != tt.expected.String() {
				t.Fatalf("command.String() not %q. got=%q", tt.expected.String(), command.String())
			}
		})
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.IfExpression
	}{
		{
			name: "if expression",
			input: `
				{
					"if": {
						"cond": true,
						"conseq": 1
					}
				}`,
			expected: &ast.IfExpression{
				Token: token.Token{Type: token.IF, Literal: "if"},
				Condition: &ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
				Consequence: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
				Alternative: nil,
			},
		},
		{
			name: "if expression with alternative",
			input: `
				{
					"if": {
						"cond": true,
						"conseq": 1,
						"alt": 2
					}
				}`,
			expected: &ast.IfExpression{
				Token: token.Token{Type: token.IF, Literal: "if"},
				Condition: &ast.Boolean{
					Token: token.Token{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
				Consequence: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
				Alternative: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "2"},
					Value: 2,
				},
			},
		},
		{
			name: "if expression: condition, consequence, and alternative are objects",
			input: `
				{
					"if": {
						"cond": {
							"command": {
								"symbol": "==",
								"args": [1, 2]
							}
						},
						"conseq": {
							"command": {
								"symbol": "+",
								"args": [3, 4]
							}
						},
						"alt": {
							"command": {
								"symbol": "*",
								"args": [5, 6]
							}
						}
					}
				}`,
			expected: &ast.IfExpression{
				Token: token.Token{Type: token.IF, Literal: "if"},
				Condition: &ast.CommandObject{
					Token: token.Token{Type: token.COMMAND, Literal: "command"},
					Symbol: &ast.Symbol{
						Token: token.Token{Type: token.EQ, Literal: "=="},
						Value: "==",
					},
					Args: &ast.Array{
						Token: token.Token{Type: token.LBRACKET, Literal: "["},
						Elements: []ast.Expression{
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "1"},
								Value: 1,
							},
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "2"},
								Value: 2,
							},
						},
					},
				},
				Consequence: &ast.CommandObject{
					Token: token.Token{Type: token.COMMAND, Literal: "command"},
					Symbol: &ast.Symbol{
						Token: token.Token{Type: token.PLUS, Literal: "+"},
						Value: "+",
					},
					Args: &ast.Array{
						Token: token.Token{Type: token.LBRACKET, Literal: "["},
						Elements: []ast.Expression{
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "3"},
								Value: 3,
							},
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "4"},
								Value: 4,
							},
						},
					},
				},
				Alternative: &ast.CommandObject{
					Token: token.Token{Type: token.COMMAND, Literal: "command"},
					Symbol: &ast.Symbol{
						Token: token.Token{Type: token.ASTERISK, Literal: "*"},
						Value: "*",
					},
					Args: &ast.Array{
						Token: token.Token{Type: token.LBRACKET, Literal: "["},
						Elements: []ast.Expression{
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "5"},
								Value: 5,
							},
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "6"},
								Value: 6,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			ifExp, ok := program.(*ast.IfExpression)
			if !ok {
				t.Fatalf("exp not *ast.IfExpression. got=%T", program)
			}
			if ifExp.String() != tt.expected.String() {
				t.Fatalf("ifExp.String() not %q. got=%q", tt.expected.String(), ifExp.String())
			}
		})
	}
}

func TestLoopExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.LoopExpression
	}{
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
								"symbol": "+",
								"args": ["$i", 1]
							}
						}
					}
				}`,
			expected: &ast.LoopExpression{
				Token: token.Token{Type: token.LOOP, Literal: "loop"},
				Index: &ast.Symbol{
					Token: token.Token{Type: token.SYMBOL, Literal: "$i"},
					Value: "$i",
				},
				From: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "0"},
					Value: 0,
				},
				To: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "10"},
					Value: 10,
				},
				Body: &ast.CommandObject{
					Token: token.Token{Type: token.COMMAND, Literal: "command"},
					Symbol: &ast.Symbol{
						Token: token.Token{Type: token.PLUS, Literal: "+"},
						Value: "+",
					},
					Args: &ast.Array{
						Token: token.Token{Type: token.LBRACKET, Literal: "["},
						Elements: []ast.Expression{
							&ast.Symbol{
								Token: token.Token{Type: token.SYMBOL, Literal: "$i"},
								Value: "$i",
							},
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "1"},
								Value: 1,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			loopExp, ok := program.(*ast.LoopExpression)
			if !ok {
				t.Fatalf("exp not *ast.LoopExpression. got=%T", program)
			}
			if loopExp.String() != tt.expected.String() {
				t.Fatalf("loopExp.String() not %q. got=%q", tt.expected.String(), loopExp.String())
			}
		})
	}
}

func TestSetExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ast.SetExpression
	}{
		{
			name: "set variable",
			input: `
				{
					"set": {
						"var": "$x",
						"val": 1
					}
				}`,
			expected: &ast.SetExpression{
				Token: token.Token{Type: token.SET, Literal: "set"},
				Name: &ast.Symbol{
					Token: token.Token{Type: token.SYMBOL, Literal: "$x"},
					Value: "$x",
				},
				Value: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "1"},
					Value: 1,
				},
			},
		},
		{
			name: "set variable with command",
			input: `
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
				}`,
			expected: &ast.SetExpression{
				Token: token.Token{Type: token.SET, Literal: "set"},
				Name: &ast.Symbol{
					Token: token.Token{Type: token.SYMBOL, Literal: "$x"},
					Value: "$x",
				},
				Value: &ast.CommandObject{
					Token: token.Token{Type: token.COMMAND, Literal: "command"},
					Symbol: &ast.Symbol{
						Token: token.Token{Type: token.SYMBOL, Literal: "+"},
						Value: "+",
					},
					Args: &ast.Array{
						Token: token.Token{Type: token.LBRACKET, Literal: "["},
						Elements: []ast.Expression{
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "1"},
								Value: 1,
							},
							&ast.IntegerLiteral{
								Token: token.Token{Type: token.INT, Literal: "2"},
								Value: 2,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error:	%v", err)
			}

			setExp, ok := program.(*ast.SetExpression)
			if !ok {
				t.Fatalf("exp not *ast.SetExpression. got=%T", program)
			}
			if setExp.String() != tt.expected.String() {
				t.Fatalf("setExp.String() not %q. got=%q", tt.expected.String(), setExp.String())
			}
		})
	}
}

func TestPrograms(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Expression
	}{
		{
			name: "multiple atoms",
			input: `
				[
					1,
					true
				]
			`,
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "1"},
						Value: 1,
					},
					&ast.Boolean{
						Token: token.Token{Type: token.TRUE, Literal: "true"},
						Value: true,
					},
				},
			},
		},
		{
			name: "multiple objects",
			input: `
				[
					{
						"command": {
							"symbol": "+",
							"args": [1, 2]
						}
					},
					{
						"if": {
							"cond": true,
							"conseq": 1
						}
					}
				]
			`,
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.CommandObject{
						Token: token.Token{Type: token.COMMAND, Literal: "command"},
						Symbol: &ast.Symbol{
							Token: token.Token{Type: token.PLUS, Literal: "+"},
							Value: "+",
						},
						Args: &ast.Array{
							Token: token.Token{Type: token.LBRACKET, Literal: "["},
							Elements: []ast.Expression{
								&ast.IntegerLiteral{
									Token: token.Token{Type: token.INT, Literal: "1"},
									Value: 1,
								},
								&ast.IntegerLiteral{
									Token: token.Token{Type: token.INT, Literal: "2"},
									Value: 2,
								},
							},
						},
					},
					&ast.IfExpression{
						Token: token.Token{Type: token.IF, Literal: "if"},
						Condition: &ast.Boolean{
							Token: token.Token{Type: token.TRUE, Literal: "true"},
							Value: true,
						},
						Consequence: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
						Alternative: nil,
					},
				},
			},
		},
		{
			name: "multiple objects with set expression",
			input: `
				[
					{
						"set": {
							"var": "$x",
							"val": 1
						}
					},
					"$x"
				]
			`,
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.SetExpression{
						Token: token.Token{Type: token.SET, Literal: "set"},
						Name: &ast.Symbol{
							Token: token.Token{Type: token.SYMBOL, Literal: "$x"},
							Value: "$x",
						},
						Value: &ast.IntegerLiteral{
							Token: token.Token{Type: token.INT, Literal: "1"},
							Value: 1,
						},
					},
					&ast.Symbol{
						Token: token.Token{Type: token.SYMBOL, Literal: "$x"},
						Value: "$x",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)

			program, err := p.ParseProgram()
			if err != nil {
				t.Fatalf("ParseProgram() error: %v", err)
			}

			array, ok := program.(*ast.Array)
			if !ok {
				t.Fatalf("exp not *ast.Array. got=%T", program)
			}

			if array.String() != tt.expected.String() {
				t.Fatalf("array.String() not %q. got=%q", tt.expected.String(), array.String())
			}
		})
	}
}
