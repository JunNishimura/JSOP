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
