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
	case *ast.KeyValueObject:
		return evalKeyValueObject(expt, env)
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

func evalKeyValueObject(kv *ast.KeyValueObject, env *object.Environment) object.Object {
	for key, value := range kv.KVPairs() {
		switch key {
		case "command":
			return evalCommandObject(value, env)
		case "if":
			return evalIfExpression(value, env)
		case "set":
			return evalSetExpression(value, env)
		case "loop":
			return evalLoopExpression(value, env)
		}
	}

	return newError("unknown key for object: %s", kv)
}

func evalCommandObject(exp ast.Expression, env *object.Environment) object.Object {
	keyValueObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return newError("invalid value for command: %s", exp)
	}
	kvPairs := keyValueObj.KVPairs()

	symbolValue, ok := kvPairs["symbol"]
	if !ok {
		return newError("symbol key not found in command: %s", keyValueObj)
	}
	symbol := Eval(symbolValue, env)
	if isError(symbol) {
		return symbol
	}

	argsValue, ok := kvPairs["args"]
	if !ok {
		return applyFunction(symbol, Null)
	}
	args := Eval(argsValue, env)
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

func evalIfExpression(exp ast.Expression, env *object.Environment) object.Object {
	keyValueObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return newError("invalid value for if: %s", exp)
	}
	kvPairs := keyValueObj.KVPairs()

	conditionValue, ok := kvPairs["cond"]
	if !ok {
		return newError("cond key not found in if: %s", keyValueObj)
	}
	condition := Eval(conditionValue, env)
	if isError(condition) {
		return condition
	}

	consequenceValue, ok := kvPairs["conseq"]
	if !ok {
		return newError("consequence key not found in if: %s", keyValueObj)
	}

	if isTruthy(condition) {
		return Eval(consequenceValue, env)
	}

	alternativeValue, ok := kvPairs["alt"]
	if !ok {
		return Null
	}

	return Eval(alternativeValue, env)
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

func evalSetExpression(exp ast.Expression, env *object.Environment) object.Object {
	keyValueObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return newError("invalid value for set: %s", exp)
	}
	kvPairs := keyValueObj.KVPairs()

	varValue, ok := kvPairs["var"]
	if !ok {
		return newError("var key not found in set: %s", keyValueObj)
	}
	variable, ok := varValue.(*ast.Symbol)
	if !ok {
		return newError("var key must be SYMBOL, got %s", varValue)
	}

	valueValue, ok := kvPairs["val"]
	if !ok {
		return newError("value key not found in set: %s", keyValueObj)
	}
	value := Eval(valueValue, env)
	if isError(value) {
		return value
	}

	return env.Set(variable.Value, value)
}

func evalLoopExpression(exp ast.Expression, env *object.Environment) object.Object {
	keyValueObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return newError("invalid value for loop: %s", exp)
	}
	kvPairs := keyValueObj.KVPairs()

	forValue, ok := kvPairs["for"]
	if !ok {
		return newError("for key not found in loop: %s", keyValueObj)
	}
	loopSymbol, ok := forValue.(*ast.Symbol)
	if !ok {
		return newError("for key must be SYMBOL, got %s", forValue)
	}

	fromValue, ok := kvPairs["from"]
	if !ok {
		return newError("from key not found in loop: %s", keyValueObj)
	}
	from := Eval(fromValue, env)
	if isError(from) {
		return from
	}
	fromInt, ok := from.(*object.Integer)
	if !ok {
		return newError("from value must be INTEGER, got %s", from.Type())
	}

	toValue, ok := kvPairs["to"]
	if !ok {
		return newError("to key not found in loop: %s", keyValueObj)
	}
	to := Eval(toValue, env)
	if isError(to) {
		return to
	}
	toInt, ok := to.(*object.Integer)
	if !ok {
		return newError("to value must be INTEGER, got %s", to.Type())
	}

	doValue, ok := kvPairs["do"]
	if !ok {
		return newError("do key not found in loop: %s", keyValueObj)
	}

	var result object.Object

	for i := fromInt.Value; i < toInt.Value; i++ {
		env.Set(loopSymbol.Value, &object.Integer{Value: i})
		evaluated := Eval(doValue, env)
		if isError(evaluated) {
			return evaluated
		}

		result = evaluated
	}

	return result
}
