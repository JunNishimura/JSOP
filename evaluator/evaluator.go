package evaluator

import (
	"fmt"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
)

var (
	Null  = &object.Null{}
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
)

func Eval(exp ast.Expression, env *object.Environment) object.Object {
	switch expt := exp.(type) {
	case *ast.Array:
		return evalArray(expt, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: expt.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(expt.Value)
	case *ast.PrefixAtom:
		right := Eval(expt.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixAtom(expt.Operator, right)
	case *ast.Symbol:
		return evalSymbol(expt, env)
	case *ast.CommandObject:
		return evalCommandObject(expt, env)
	case *ast.IfExpression:
		return evalIfExpression(expt, env)
	case *ast.SetExpression:
		value := Eval(expt.Value, env)
		if isError(value) {
			return value
		}
		return env.Set(expt.Name.Value, value)
	case *ast.LoopExpression:
		return evalLoopExpression(expt, env)
	default:
		return newError("unknown expression type: %T", exp)
	}
}

func evalArray(array *ast.Array, env *object.Environment) *object.Array {
	result := &object.Array{
		Elements: []object.Object{},
	}

	for _, el := range array.Elements {
		evaluated := Eval(el, env)
		if isError(evaluated) {
			return &object.Array{Elements: []object.Object{evaluated}}
		}
		result.Elements = append(result.Elements, evaluated)
	}

	return result
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}

func evalPrefixAtom(operator string, right object.Object) object.Object {
	switch operator {
	case "-":
		return evalMinusPrefix(right)
	default:
		return newError("unknown operator: %s%s", operator, right)
	}
}

func evalMinusPrefix(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right)
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalSymbol(symbol *ast.Symbol, env *object.Environment) object.Object {
	builtintFunc, ok := builtins[symbol.Value]
	if ok {
		return builtintFunc
	}

	obj, ok := env.Get(symbol.Value)
	if ok {
		return obj
	}

	return newError("symbol not found: %s", symbol.Value)
}

func evalCommandObject(command *ast.CommandObject, env *object.Environment) object.Object {
	symbol := Eval(command.Symbol, env)
	if isError(symbol) {
		return symbol
	}

	if command.Args == nil {
		return applyFunction(symbol, Null)
	}

	args := Eval(command.Args, env)
	if isError(args) {
		return args
	}

	return applyFunction(symbol, args)
}

func applyFunction(function object.Object, args object.Object) object.Object {
	switch funcType := function.(type) {
	case *object.Builtin:
		return funcType.Fn(args)
	default:
		return newError("not a function: %s", function.Type())
	}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return False
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case True:
		return true
	case False:
		return false
	default:
		return true
	}
}

func evalLoopExpression(le *ast.LoopExpression, env *object.Environment) object.Object {
	from := Eval(le.From, env)
	if isError(from) {
		return from
	}
	fromValue, ok := from.(*object.Integer)
	if !ok {
		return newError("from value must be INTEGER, got %s", from.Type())
	}

	to := Eval(le.To, env)
	if isError(to) {
		return to
	}
	toValue, ok := to.(*object.Integer)
	if !ok {
		return newError("to value must be INTEGER, got %s", to.Type())
	}

	var result object.Object

	for i := fromValue.Value; i < toValue.Value; i++ {
		env.Set(le.Index.Value, &object.Integer{Value: i})
		evaluated := Eval(le.Body, env)
		if isError(evaluated) {
			return evaluated
		}

		result = evaluated
	}

	return result
}
