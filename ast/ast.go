package ast

import (
	"bytes"
	"fmt"

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

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return fmt.Sprintf("\"%s\"", sl.Value) }

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
func (s *Symbol) String() string       { return fmt.Sprintf("\"%s\"", s.Value) }

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

type Object interface {
	Expression
	KVPairs() map[string]Expression
}

type KeyValueObject struct {
	Token token.Token
	KV    []*KeyValuePair
}

func (k *KeyValueObject) TokenLiteral() string { return k.Token.Literal }
func (k *KeyValueObject) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	for i, kv := range k.KV {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(kv.String())
	}
	out.WriteString("}")

	return out.String()
}
func (k *KeyValueObject) KVPairs() map[string]Expression {
	kvPairs := make(map[string]Expression)
	for _, kv := range k.KV {
		kvPairs[kv.Key.Value] = kv.Value
	}
	return kvPairs
}

type KeyValuePair struct {
	Key   *StringLiteral
	Value Expression
}

func (k *KeyValuePair) String() string {
	return fmt.Sprintf("%s: %s", k.Key.String(), k.Value.String())
}
