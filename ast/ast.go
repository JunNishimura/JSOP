package ast

import (
	"bytes"

	"github.com/JunNishimura/jsop/token"
)

type Expression interface {
	TokenLiteral() string
	String() string
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

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

type Array struct {
	Token    token.Token
	Elements []Expression
}

func (a *Array) TokenLiteral() string { return a.Token.Literal }
func (a *Array) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	for i, el := range a.Elements {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(el.String())
	}
	out.WriteString("]")

	return out.String()
}

type CommandObject struct {
	Token  token.Token
	Symbol Expression
	Args   Expression
}

func (c *CommandObject) TokenLiteral() string { return c.Token.Literal }
func (c *CommandObject) String() string {
	var out bytes.Buffer

	out.WriteString("{\"command\": {\"symbol\": ")
	out.WriteString(c.Symbol.String())

	if c.Args == nil {
		// no args
		out.WriteString("}}")
	} else {
		out.WriteString(", \"args\": ")
		out.WriteString(c.Args.String())
		out.WriteString("}}")
	}

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("{\"if\": {\"cond\": ")
	out.WriteString(ie.Condition.String())
	out.WriteString(", \"conseq\": ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString(", \"alt\": ")
		out.WriteString(ie.Alternative.String())
	}
	out.WriteString("}}")

	return out.String()
}

type SetExpression struct {
	Token token.Token
	Name  *Symbol
	Value Expression
}

func (se *SetExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SetExpression) String() string {
	var out bytes.Buffer

	out.WriteString("{\"set\": {\"var\": ")
	out.WriteString(se.Name.String())
	out.WriteString(", \"val\": ")
	out.WriteString(se.Value.String())
	out.WriteString("}}")

	return out.String()
}
