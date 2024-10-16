package evaluator

import (
	"testing"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/object"
	"github.com/JunNishimura/jsop/parser"
)

func testParseProgram(input string) (ast.Expression, error) {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func TestDefineMacro(t *testing.T) {
	input := `
        {
            "defmacro": {
                "name": "myMacro",
                "keys": ["foo", "bar"],
                "body": {
                    "command": {
                        "symbol": "quote",
                        "args": {
                            "command": {
                                "symbol": "+",
                                "args": [",foo", ",bar"]
                            }
                        }
                    }
                }
            }
        }`

	env := object.NewEnvironment()
	exp, err := testParseProgram(input)
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}

	if err := DefineMacros(exp, env); err != nil {
		t.Fatalf("define macro error: %s", err)
	}

	obj, ok := env.Get("myMacro")
	if !ok {
		t.Fatalf("macro not in environment")
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		t.Fatalf("object is not Macro. got=%T (%+v)", obj, obj)
	}

	if len(macro.Keys) != 2 {
		t.Fatalf("wrong number of macro Keys. want=2, got=%d", len(macro.Keys))
	}

	if macro.Keys[0].Value != "foo" {
		t.Fatalf("key is not foo. got=%q", macro.Keys[0])
	}
	if macro.Keys[1].Value != "bar" {
		t.Fatalf("key is not bar. got=%q", macro.Keys[1])
	}

	expectedBody := "{\"command\": {\"symbol\": \"quote\", \"args\": {\"command\": {\"symbol\": \"+\", \"args\": [\",foo\", \",bar\"]}}}}"
	if macro.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, macro.Body.String())
	}
}

func TestExpandMacro(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `
                [
                    {
                        "defmacro": {
                            "name": "hoge",
                            "body": {
                                "command": {
                                    "symbol": "quote",
                                    "args": {
                                        "command": {
                                            "symbol": "+",
                                            "args": [1, 2]
                                        }
                                    }
                                }
                            }
                        }
                    },
                    {
                        "hoge": {}
                    }
                ]`,
			expected: "[{\"command\": {\"symbol\": \"+\", \"args\": [1, 2]}}]",
		},
		{
			input: `
                [
                    {
                        "defmacro": {
                            "name": "reverse",
                            "keys": ["first", "second"],
                            "body": {
                                "command": {
                                    "symbol": "quote",
                                    "args": {
                                        "command": {
                                            "symbol": "-",
                                            "args": [",second", ",first"]
                                        }
                                    }
                                }
                            }
                        }
                    },
                    {
                        "reverse": {
                            "first": 2,
                            "second": 3
                        }
                    }
                ]`,
			expected: "[{\"command\": {\"symbol\": \"-\", \"args\": [3, 2]}}]",
		},
		{
			input: `
                [
                    {
                        "defmacro": {
                            "name": "unless",
                            "keys": ["cond", "conseq", "alt"],
                            "body": {
                                "command": {
                                    "symbol": "quote",
                                    "args": {
                                        "if": {
                                            "cond": {"command": {"symbol": "!", "args": ",cond"}},
                                            "conseq": ",conseq",
                                            "alt": ",alt"
                                        }
                                    }
                                } 
                            }
                        }
                    },
                    {
                        "unless": {
                            "cond": {
                                "command": {
                                    "symbol": ">",
                                    "args": [2, 1]
                                }
                            },
                            "conseq": {
                                "command": {
                                    "symbol": "print",
                                    "args": "not greater"
                                }
                            },
                            "alt": {
                                "command": {
                                    "symbol": "print",
                                    "args": "greater"
                                }
                            }
                        }
                    }
                ]`,
			expected: "[{\"if\": {\"cond\": {\"command\": {\"symbol\": \"!\", \"args\": {\"command\": {\"symbol\": \">\", \"args\": [2, 1]}}}}, \"conseq\": {\"command\": {\"symbol\": \"print\", \"args\": \"not greater\"}}, \"alt\": {\"command\": {\"symbol\": \"print\", \"args\": \"greater\"}}}}]",
		},
	}

	for _, tt := range tests {
		expected, err := testParseProgram(tt.expected)
		if err != nil {
			t.Fatalf("parse error for expected: %s", err)
		}
		program, err := testParseProgram(tt.input)
		if err != nil {
			t.Fatalf("parse error for input: %s", err)
		}

		env := object.NewEnvironment()
		if err := DefineMacros(program, env); err != nil {
			t.Fatalf("define macro error: %s", err)
		}
		expanded := ExpandMacros(program, env)

		if expanded.String() != expected.String() {
			t.Errorf("not equal. got=%q, want=%q", expanded.String(), expected.String())
		}
	}
}
