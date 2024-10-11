package evaluator

import (
	"fmt"
	"strings"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
	"github.com/JunNishimura/jsop/token"
)

func quote(exp ast.Expression, env *object.Environment) object.Object {
	unquotedExp := evalUnquote(exp, env)
	return &object.Quote{Expression: unquotedExp}
}

func evalUnquote(quoted ast.Expression, env *object.Environment) ast.Expression {
	return ast.Modify(quoted, func(exp ast.Expression) ast.Expression {
		if !isUnquote(exp) {
			return exp
		}

		kvObject, ok := exp.(*ast.KeyValueObject)
		if ok {
			return evalUnquoteCommand(kvObject, env)
		}

		strLit, ok := exp.(*ast.StringLiteral)
		if !ok {
			return exp
		}

		if obj, isFound := env.Get(strLit.Value[1:]); isFound {
			return convertObjectToExpression(obj)
		}

		return exp
	})
}

func evalUnquoteCommand(kvObject *ast.KeyValueObject, env *object.Environment) ast.Expression {
	commandValue, ok := kvObject.KVPairs()["command"]
	if !ok {
		return kvObject
	}
	commandObject, ok := commandValue.(*ast.KeyValueObject)
	if !ok {
		return kvObject
	}

	argsValue, ok := commandObject.KVPairs()["args"]
	if !ok {
		return kvObject
	}

	unquoted := Eval(argsValue, env)
	return convertObjectToExpression(unquoted)
}

func isUnquote(exp ast.Expression) bool {
	kvObject, ok := exp.(*ast.KeyValueObject)
	if ok {
		return isUnquoteCommand(kvObject)
	}

	strLit, ok := exp.(*ast.StringLiteral)
	if ok {
		return strings.HasPrefix(strLit.Value, ",")
	}

	return false
}

func isUnquoteCommand(exp ast.Expression) bool {
	kvObject, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return false
	}

	commandValue, ok := kvObject.KVPairs()["command"]
	if !ok {
		return false
	}
	commandObject, ok := commandValue.(*ast.KeyValueObject)
	if !ok {
		return false
	}

	symbol, ok := commandObject.KVPairs()["symbol"]
	if !ok {
		return false
	}

	symbolStr, ok := symbol.(*ast.StringLiteral)
	if !ok {
		return false
	}

	return symbolStr.Value == "unquote"
}

func convertObjectToExpression(obj object.Object) ast.Expression {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}
	case *object.Boolean:
		if obj.Value {
			return &ast.Boolean{
				Token: token.Token{Type: token.TRUE, Literal: "true"},
				Value: true,
			}
		}
		return &ast.Boolean{
			Token: token.Token{Type: token.FALSE, Literal: "false"},
			Value: false,
		}
	case *object.String:
		return &ast.StringLiteral{Token: token.Token{Type: token.STRING, Literal: obj.Value}, Value: obj.Value}
	case *object.Array:
		elements := make([]ast.Expression, len(obj.Elements))
		for i, el := range obj.Elements {
			elements[i] = convertObjectToExpression(el)
		}
		return &ast.Array{
			Token:    token.Token{Type: token.LBRACKET, Literal: "["},
			Elements: elements,
		}
	case *object.Quote:
		return obj.Expression
	default:
		return nil
	}
}
