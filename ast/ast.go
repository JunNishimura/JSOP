package ast

import (
	"bytes"

	"github.com/JunNishimura/jsop/token"
)

type Expression interface {
	TokenLiteral() string
	String() string
}

type Program struct {
	Expressions []Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Expressions) > 0 {
		return p.Expressions[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, exp := range p.Expressions {
		out.WriteString(exp.String())
	}

	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixAtom struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pa *PrefixAtom) TokenLiteral() string { return pa.Token.Literal }
func (pa *PrefixAtom) String() string {
	var out bytes.Buffer

	out.WriteString(pa.Operator)
	out.WriteString(pa.Right.String())

	return out.String()
}

type Symbol struct {
	Token token.Token
	Value string
}

func (s *Symbol) TokenLiteral() string { return s.Token.Literal }
func (s *Symbol) String() string       { return s.Token.Literal }

type CommandObject struct {
	Token  token.Token
	Symbol *Symbol
	Args   []Expression
}

func (c *CommandObject) TokenLiteral() string { return c.Token.Literal }
func (c *CommandObject) String() string {
	var out bytes.Buffer

	out.WriteString("{\"command\": {\"symbol\": ")
	out.WriteString(c.Symbol.String())
	out.WriteString(", \"args\": [")
	for i, arg := range c.Args {
		out.WriteString(arg.String())
		if i != len(c.Args)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("]}}")

	return out.String()
}
