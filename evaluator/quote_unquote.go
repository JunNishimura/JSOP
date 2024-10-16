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

		switch exp := exp.(type) {
		case *ast.KeyValueObject:
			return evalUnquoteCommand(exp, env)
		case *ast.StringLiteral:
			return evalUnquoteString(exp, env)
		}

		return exp
	})
}

func evalUnquoteCommand(kvObj *ast.KeyValueObject, env *object.Environment) ast.Expression {
	cmdVal, ok := kvObj.KVPairs()["command"]
	if !ok {
		return kvObj
	}
	cmdObj, ok := cmdVal.(*ast.KeyValueObject)
	if !ok {
		return kvObj
	}

	argsVal, ok := cmdObj.KVPairs()["args"]
	if !ok {
		return kvObj
	}

	unquoted := Eval(argsVal, env)
	return convertObjectToExpression(unquoted)
}

func evalUnquoteString(strLit *ast.StringLiteral, env *object.Environment) ast.Expression {
	if obj, isFound := env.Get(strLit.Value[1:]); isFound {
		return convertObjectToExpression(obj)
	}
	return strLit
}

func isUnquote(exp ast.Expression) bool {
	switch exp := exp.(type) {
	case *ast.KeyValueObject:
		return isUnquoteCommand(exp)
	case *ast.StringLiteral:
		return strings.HasPrefix(exp.Value, ",")
	}

	return false
}

func isUnquoteCommand(exp ast.Expression) bool {
	kvObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return false
	}

	cmdVal, ok := kvObj.KVPairs()["command"]
	if !ok {
		return false
	}
	cmdObj, ok := cmdVal.(*ast.KeyValueObject)
	if !ok {
		return false
	}

	symbolVal, ok := cmdObj.KVPairs()["symbol"]
	if !ok {
		return false
	}
	symbolStr, ok := symbolVal.(*ast.StringLiteral)
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
