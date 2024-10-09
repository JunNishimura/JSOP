package evaluator

import (
	"fmt"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
	"github.com/JunNishimura/jsop/token"
)

func quote(exp ast.Expression, env *object.Environment) object.Object {
	unquotedExp := evalUnquoteCommand(exp, env)
	return &object.Quote{Expression: unquotedExp}
}

func evalUnquoteCommand(quoted ast.Expression, env *object.Environment) ast.Expression {
	return ast.Modify(quoted, func(exp ast.Expression) ast.Expression {
		if !isUnquoteCommand(exp) {
			return exp
		}

		kvObject, ok := exp.(*ast.KeyValueObject)
		if !ok {
			return exp
		}

		commandValue, ok := kvObject.KVPairs()["command"]
		if !ok {
			return exp
		}
		commandObject, ok := commandValue.(*ast.KeyValueObject)
		if !ok {
			return exp
		}

		argsValue, ok := commandObject.KVPairs()["args"]
		if !ok {
			return exp
		}

		unquoted := Eval(argsValue, env)
		return convertObjectToExpression(unquoted)
	})
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

	symbolStr, ok := symbol.(*ast.Symbol)
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
