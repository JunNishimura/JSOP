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
		{
			name:     "string with space",
			input:    `"hello world"`,
			expected: "hello world",
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
		expected *ast.KeyValueObject
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "+"},
										Value: "+",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "-"},
										Value: "-",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "*"},
										Value: "*",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "/"},
										Value: "/",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "=="},
										Value: "==",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "!="},
										Value: "!=",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "!"},
										Value: "!",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Boolean{
										Token: token.Token{Type: token.TRUE, Literal: "true"},
										Value: true,
									},
								},
							},
						},
					},
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$hoge"},
										Value: "$hoge",
									},
								},
							},
						},
					},
				},
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$hoge"},
										Value: "$hoge",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "access array element",
			input: `
				{
					"command": {
						"symbol": "at",
						"args": ["$hoge", 1]
					}
				}`,
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "at"},
										Value: "at",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
										Token: token.Token{Type: token.LBRACKET, Literal: "["},
										Elements: []ast.Expression{
											&ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$hoge"},
												Value: "$hoge",
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
					},
				},
			},
		},
		{
			name: "single line comment",
			input: `
				{
					"command": {
						"//": "this is a comment",
						"symbol": "+",
						"args": [1, 2]
					}
				}`,
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "command"},
							Value: "command",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "//"},
										Value: "//",
									},
									Value: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "this is a comment"},
										Value: "this is a comment",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "symbol"},
										Value: "symbol",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "+"},
										Value: "+",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "args"},
										Value: "args",
									},
									Value: &ast.Array{
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

			kvObject, ok := program.(*ast.KeyValueObject)
			if !ok {
				t.Fatalf("exp not *ast.KeyValueObject. got=%T", program)
			}
			if kvObject.String() != tt.expected.String() {
				t.Fatalf("kvObject.String() not %q. got=%q", tt.expected.String(), kvObject.String())
			}
		})
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Object
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "if"},
							Value: "if",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "cond"},
										Value: "cond",
									},
									Value: &ast.Boolean{
										Token: token.Token{Type: token.TRUE, Literal: "true"},
										Value: true,
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "conseq"},
										Value: "conseq",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
							},
						},
					},
				},
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "if"},
							Value: "if",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "cond"},
										Value: "cond",
									},
									Value: &ast.Boolean{
										Token: token.Token{Type: token.TRUE, Literal: "true"},
										Value: true,
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "conseq"},
										Value: "conseq",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "alt"},
										Value: "alt",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "2"},
										Value: 2,
									},
								},
							},
						},
					},
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "if"},
							Value: "if",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "cond"},
										Value: "cond",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "=="},
																Value: "==",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
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
											},
										},
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "conseq"},
										Value: "conseq",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "+"},
																Value: "+",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
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
													},
												},
											},
										},
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "alt"},
										Value: "alt",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "*"},
																Value: "*",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
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
											},
										},
									},
								},
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

			kvObject, ok := program.(*ast.KeyValueObject)
			if !ok {
				t.Fatalf("exp not *ast.KeyValueObject. got=%T", program)
			}
			if kvObject.String() != tt.expected.String() {
				t.Fatalf("kvObject.String() not %q. got=%q", tt.expected.String(), kvObject.String())
			}
		})
	}
}

func TestLoopExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Expression
	}{
		{
			name: "loop expression",
			input: `
				{
					"loop": {
						"for": "$i",
						"from": 0,
						"until": 10,
						"do": {
							"command": {
								"symbol": "+",
								"args": ["$i", 1]
							}
						}
					}
				}`,
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "loop"},
							Value: "loop",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "for"},
										Value: "for",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$i"},
										Value: "$i",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "from"},
										Value: "from",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "0"},
										Value: 0,
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "until"},
										Value: "until",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "10"},
										Value: 10,
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "do"},
										Value: "do",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "+"},
																Value: "+",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
																Token: token.Token{Type: token.LBRACKET, Literal: "["},
																Elements: []ast.Expression{
																	&ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "$i"},
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
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "loop expression with in keyword",
			input: `
				[
					{
						"set": {
							"var": "$arr",
							"val": [10, 20, 30]	
						}
					},
					{
						"loop": {
							"for": "$element",
							"in": "$arr",
							"do": {
								"command": {
									"symbol": "print",
									"args": "$element"
								}
							}
						}
					}
				]`,
			expected: &ast.Array{
				Token: token.Token{Type: token.LBRACKET, Literal: "["},
				Elements: []ast.Expression{
					&ast.KeyValueObject{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						KV: []*ast.KeyValuePair{
							{
								Key: &ast.StringLiteral{
									Token: token.Token{Type: token.STRING, Literal: "set"},
									Value: "set",
								},
								Value: &ast.KeyValueObject{
									Token: token.Token{Type: token.LBRACE, Literal: "{"},
									KV: []*ast.KeyValuePair{
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "var"},
												Value: "var",
											},
											Value: &ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$arr"},
												Value: "$arr",
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "val"},
												Value: "val",
											},
											Value: &ast.Array{
												Token: token.Token{Type: token.LBRACKET, Literal: "["},
												Elements: []ast.Expression{
													&ast.IntegerLiteral{
														Token: token.Token{Type: token.INT, Literal: "10"},
														Value: 10,
													},
													&ast.IntegerLiteral{
														Token: token.Token{Type: token.INT, Literal: "20"},
														Value: 20,
													},
													&ast.IntegerLiteral{
														Token: token.Token{Type: token.INT, Literal: "30"},
														Value: 30,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					&ast.KeyValueObject{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						KV: []*ast.KeyValuePair{
							{
								Key: &ast.StringLiteral{
									Token: token.Token{Type: token.STRING, Literal: "loop"},
									Value: "loop",
								},
								Value: &ast.KeyValueObject{
									Token: token.Token{Type: token.LBRACE, Literal: "{"},
									KV: []*ast.KeyValuePair{
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "for"},
												Value: "for",
											},
											Value: &ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$element"},
												Value: "$element",
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "in"},
												Value: "in",
											},
											Value: &ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$arr"},
												Value: "$arr",
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "do"},
												Value: "do",
											},
											Value: &ast.KeyValueObject{
												Token: token.Token{Type: token.LBRACE, Literal: "{"},
												KV: []*ast.KeyValuePair{
													{
														Key: &ast.StringLiteral{
															Token: token.Token{Type: token.STRING, Literal: "command"},
															Value: "command",
														},
														Value: &ast.KeyValueObject{
															Token: token.Token{Type: token.LBRACE, Literal: "{"},
															KV: []*ast.KeyValuePair{
																{
																	Key: &ast.StringLiteral{
																		Token: token.Token{Type: token.STRING, Literal: "symbol"},
																		Value: "symbol",
																	},
																	Value: &ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "print"},
																		Value: "print",
																	},
																},
																{
																	Key: &ast.StringLiteral{
																		Token: token.Token{Type: token.STRING, Literal: "args"},
																		Value: "args",
																	},
																	Value: &ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "$element"},
																		Value: "$element",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
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

			if program.String() != tt.expected.String() {
				t.Fatalf("program.String() not %q. got=%q", tt.expected.String(), program.String())
			}
		})
	}
}

func TestSetExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Object
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "set"},
							Value: "set",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "var"},
										Value: "var",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$x"},
										Value: "$x",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "val"},
										Value: "val",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
							},
						},
					},
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "set"},
							Value: "set",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "var"},
										Value: "var",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$x"},
										Value: "$x",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "val"},
										Value: "val",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "+"},
																Value: "+",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
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
											},
										},
									},
								},
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

			kvObject, ok := program.(*ast.KeyValueObject)
			if !ok {
				t.Fatalf("exp not *ast.KeyValueObject. got=%T", program)
			}
			if kvObject.String() != tt.expected.String() {
				t.Fatalf("kvObject.String() not %q. got=%q", tt.expected.String(), kvObject.String())
			}
		})
	}
}

func TestLambdaExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Object
	}{
		{
			name: "lambda expression without parameters",
			input: `
				{
					"lambda": {
						"body": 1
					}
				}`,
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "lambda"},
							Value: "lambda",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "body"},
										Value: "body",
									},
									Value: &ast.IntegerLiteral{
										Token: token.Token{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "lambda expression with 1 argument",
			input: `
				{
					"lambda": {
						"params": "$x",
						"body": {
							"command": {
								"symbol": "+",
								"args": ["$x", 1]
							}
						}
					}
				}`,
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "lambda"},
							Value: "lambda",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "params"},
										Value: "params",
									},
									Value: &ast.Symbol{
										Token: token.Token{Type: token.STRING, Literal: "$x"},
										Value: "$x",
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "body"},
										Value: "body",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "+"},
																Value: "+",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
																Token: token.Token{Type: token.LBRACKET, Literal: "["},
																Elements: []ast.Expression{
																	&ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "$x"},
																		Value: "$x",
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
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "lambda expression with multiple argument",
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
			expected: &ast.KeyValueObject{
				Token: token.Token{Type: token.LBRACE, Literal: "{"},
				KV: []*ast.KeyValuePair{
					{
						Key: &ast.StringLiteral{
							Token: token.Token{Type: token.STRING, Literal: "lambda"},
							Value: "lambda",
						},
						Value: &ast.KeyValueObject{
							Token: token.Token{Type: token.LBRACE, Literal: "{"},
							KV: []*ast.KeyValuePair{
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "params"},
										Value: "params",
									},
									Value: &ast.Array{
										Token: token.Token{Type: token.LBRACKET, Literal: "["},
										Elements: []ast.Expression{
											&ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$x"},
												Value: "$x",
											},
											&ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$y"},
												Value: "$y",
											},
										},
									},
								},
								{
									Key: &ast.StringLiteral{
										Token: token.Token{Type: token.STRING, Literal: "body"},
										Value: "body",
									},
									Value: &ast.KeyValueObject{
										Token: token.Token{Type: token.LBRACE, Literal: "{"},
										KV: []*ast.KeyValuePair{
											{
												Key: &ast.StringLiteral{
													Token: token.Token{Type: token.STRING, Literal: "command"},
													Value: "command",
												},
												Value: &ast.KeyValueObject{
													Token: token.Token{Type: token.LBRACE, Literal: "{"},
													KV: []*ast.KeyValuePair{
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "symbol"},
																Value: "symbol",
															},
															Value: &ast.Symbol{
																Token: token.Token{Type: token.STRING, Literal: "+"},
																Value: "+",
															},
														},
														{
															Key: &ast.StringLiteral{
																Token: token.Token{Type: token.STRING, Literal: "args"},
																Value: "args",
															},
															Value: &ast.Array{
																Token: token.Token{Type: token.LBRACKET, Literal: "["},
																Elements: []ast.Expression{
																	&ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "$x"},
																		Value: "$x",
																	},
																	&ast.Symbol{
																		Token: token.Token{Type: token.STRING, Literal: "$y"},
																		Value: "$y",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
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

			kvObject, ok := program.(*ast.KeyValueObject)
			if !ok {
				t.Fatalf("exp not *ast.KeyValueObject. got=%T", program)
			}
			if kvObject.String() != tt.expected.String() {
				t.Fatalf("kvObject.String() not %q. got=%q", tt.expected.String(), kvObject.String())
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
					&ast.KeyValueObject{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						KV: []*ast.KeyValuePair{
							{
								Key: &ast.StringLiteral{
									Token: token.Token{Type: token.STRING, Literal: "command"},
									Value: "command",
								},
								Value: &ast.KeyValueObject{
									Token: token.Token{Type: token.LBRACE, Literal: "{"},
									KV: []*ast.KeyValuePair{
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "symbol"},
												Value: "symbol",
											},
											Value: &ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "+"},
												Value: "+",
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "args"},
												Value: "args",
											},
											Value: &ast.Array{
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
							},
						},
					},
					&ast.KeyValueObject{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						KV: []*ast.KeyValuePair{
							{
								Key: &ast.StringLiteral{
									Token: token.Token{Type: token.STRING, Literal: "if"},
									Value: "if",
								},
								Value: &ast.KeyValueObject{
									Token: token.Token{Type: token.LBRACE, Literal: "{"},
									KV: []*ast.KeyValuePair{
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "cond"},
												Value: "cond",
											},
											Value: &ast.Boolean{
												Token: token.Token{Type: token.TRUE, Literal: "true"},
												Value: true,
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "conseq"},
												Value: "conseq",
											},
											Value: &ast.IntegerLiteral{
												Token: token.Token{Type: token.INT, Literal: "1"},
												Value: 1,
											},
										},
									},
								},
							},
						},
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
					&ast.KeyValueObject{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						KV: []*ast.KeyValuePair{
							{
								Key: &ast.StringLiteral{
									Token: token.Token{Type: token.STRING, Literal: "set"},
									Value: "set",
								},
								Value: &ast.KeyValueObject{
									Token: token.Token{Type: token.LBRACE, Literal: "{"},
									KV: []*ast.KeyValuePair{
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "var"},
												Value: "var",
											},
											Value: &ast.Symbol{
												Token: token.Token{Type: token.STRING, Literal: "$x"},
												Value: "$x",
											},
										},
										{
											Key: &ast.StringLiteral{
												Token: token.Token{Type: token.STRING, Literal: "val"},
												Value: "val",
											},
											Value: &ast.IntegerLiteral{
												Token: token.Token{Type: token.INT, Literal: "1"},
												Value: 1,
											},
										},
									},
								},
							},
						},
					},
					&ast.Symbol{
						Token: token.Token{Type: token.STRING, Literal: "$x"},
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
