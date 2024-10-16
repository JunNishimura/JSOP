package evaluator

import (
	"testing"

	"github.com/JunNishimura/jsop/object"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "quote integer",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": 1
					}
				}`,
			expected: "1",
		},
		{
			name: "quote string",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": "hello"
					}
				}`,
			expected: "\"hello\"",
		},
		{
			name: "quote boolean",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": true
					}
				}`,
			expected: "true",
		},
		{
			name: "quote prefix expression",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": -10
					}
				}`,
			expected: "-10",
		},
		{
			name: "quote array",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": [1, 2, 3]
					}
				}`,
			expected: "[1, 2, 3]",
		},
		{
			name: "quote key value object",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "+",
								"args": [1, 2]
							}
						}
					}
				}`,
			expected: "{\"command\": {\"symbol\": \"+\", \"args\": [1, 2]}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(t, tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected object.Quote. got=%T", evaluated)
			}
			if quote.Expression.String() != tt.expected {
				t.Errorf("expected=%q. got=%q", tt.expected, quote.Expression.String())
			}
		})
	}
}

func TestQuoteUnquote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "quote unquote integer",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": 1
							}
						}								
					}
				}`,
			expected: "1",
		},
		{
			name: "quote unquote string",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": "hello"
							}
						}
					}
				}`,
			expected: "\"hello\"",
		},
		{
			name: "quote unquote boolean",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": true
							}
						}
					}
				}`,
			expected: "true",
		},
		{
			name: "quote unquote prefix expression",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": -10
							}
						}
					}
				}`,
			expected: "-10",
		},
		{
			name: "quote unquote array",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": [1, 2, 3]
							}
						}
					}
				}`,
			expected: "[1, 2, 3]",
		},
		{
			name: "quote unquote key value object",
			input: `
				{
					"command": {
						"symbol": "quote",
						"args": {
							"command": {
								"symbol": "unquote",
								"args": {
									"command": {
										"symbol": "+",
										"args": [1, 2]
									}
								}
							}
						}
					}
				}`,
			expected: "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(t, tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected object.Quote. got=%T", evaluated)
			}
			if quote.Expression.String() != tt.expected {
				t.Errorf("expected=%q. got=%q", tt.expected, quote.Expression.String())
			}
		})
	}
}
