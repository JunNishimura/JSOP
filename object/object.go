package object

import (
	"bytes"
	"fmt"

	"github.com/JunNishimura/jsop/ast"
)

const (
	ERROR_OBJ    = "ERROR"
	INTEGER_OBJ  = "INTEGER"
	STRING_OBJ   = "STRING"
	BOOLEAN_OBJ  = "BOOLEAN"
	ARRAY_OBJ    = "ARRAY"
	NULL_OBJ     = "NULL"
	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"
	QUOTE_OBJ    = "QUOTE"
	MACRO_OBJ    = "MACRO"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	out.WriteString("[")
	for i, el := range a.Elements {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(el.Inspect())
	}
	out.WriteString("]")

	return out.String()
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.StringLiteral
	Body       ast.Expression
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	out.WriteString("fn(")
	for i, p := range f.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(p.Value)
	}
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type BuiltinFunction func(args Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Quote struct {
	Expression ast.Expression
}

func (q *Quote) Type() ObjectType { return QUOTE_OBJ }
func (q *Quote) Inspect() string  { return q.Expression.String() }

type Macro struct {
	Keys []*ast.StringLiteral
	Body ast.Expression
	Env  *Environment
}

func (m *Macro) Type() ObjectType { return MACRO_OBJ }
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	out.WriteString("macro(")
	for i, k := range m.Keys {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(k.Value)
	}
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}
