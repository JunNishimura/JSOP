package evaluator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
)

var (
	Null     = &object.Null{}
	Break    = &object.Break{}
	Continue = &object.Continue{}
	True     = &object.Boolean{Value: true}
	False    = &object.Boolean{Value: false}
)

const identEmbedPattern = `\{\s*\$\w+\s*\}`

func Eval(exp ast.Expression, env *object.Environment) object.Object {
	switch expt := exp.(type) {
	case *ast.Array:
		return evalArray(expt, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: expt.Value}
	case *ast.StringLiteral:
		if strings.HasPrefix(expt.Value, "$") {
			return evalSymbol(expt, env)
		}

		re := regexp.MustCompile(identEmbedPattern)
		matches := re.FindAllString(expt.Value, -1)

		if len(matches) == 0 {
			return &object.String{Value: expt.Value}
		}

		// expand embedded identifiers
		return evalEmbeddedIdentifiers(expt, env, matches)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(expt.Value)
	case *ast.PrefixAtom:
		right := Eval(expt.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixAtom(expt.Operator, right)
	case *ast.KeyValueObject:
		return evalKeyValueObject(expt, env)
	default:
		return newError("unknown expression type: %T", exp)
	}
}

func evalArray(array *ast.Array, env *object.Environment) object.Object {
	result := &object.Array{
		Elements: []object.Object{},
	}

	for _, el := range array.Elements {
		evaluated := Eval(el, env)
		if isError(evaluated) {
			return &object.Array{Elements: []object.Object{evaluated}}
		}
		if returnValue, ok := evaluated.(*object.ReturnValue); ok {
			return returnValue
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

func evalSymbol(symbol *ast.StringLiteral, env *object.Environment) object.Object {
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
		case "lambda":
			return evalLambdaExpression(value, env)
		case "break":
			return Break
		case "continue":
			return Continue
		case "return":
			evaluated := Eval(value, env)
			if isError(evaluated) {
				return evaluated
			}
			return &object.ReturnValue{Value: evaluated}
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

	var symbol object.Object
	if symbolStr, ok := symbolValue.(*ast.StringLiteral); ok {
		if symbolStr.Value == "quote" {
			argsValue, ok := kvPairs["args"]
			if !ok {
				return newError("quote command requires args key: %s", keyValueObj)
			}
			return quote(argsValue, env)
		}

		symbol = evalSymbol(symbolStr, env)
		if isError(symbol) {
			return symbol
		}
	} else {
		symbol = Eval(symbolValue, env)
		if isError(symbol) {
			return symbol
		}
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

func extendFunctionEnv(fn *object.Function, args object.Object) (*object.Environment, error) {
	switch args := args.(type) {
	case *object.Array:
		if len(fn.Parameters) != len(args.Elements) {
			return nil, fmt.Errorf("wrong number of arguments. want=%d, got=%d", len(fn.Parameters), len(args.Elements))
		}

		extendedEnv := object.NewEnclosedEnvironment(fn.Env)
		for i, param := range fn.Parameters {
			extendedEnv.Set(param.Value, args.Elements[i])
		}

		return extendedEnv, nil
	case *object.Integer, *object.Boolean:
		if len(fn.Parameters) != 1 {
			return nil, fmt.Errorf("wrong number of arguments. want=%d, got=1", len(fn.Parameters))
		}

		extendedEnv := object.NewEnclosedEnvironment(fn.Env)
		extendedEnv.Set(fn.Parameters[0].Value, args)

		return extendedEnv, nil
	case *object.Null:
		if len(fn.Parameters) != 0 {
			return nil, fmt.Errorf("wrong number of arguments. want=%d, got=0", len(fn.Parameters))
		}

		return object.NewEnclosedEnvironment(fn.Env), nil
	default:
		return nil, fmt.Errorf("unhandled argument type: %s", args.Type())
	}
}

func applyFunction(function object.Object, args object.Object) object.Object {
	switch funcType := function.(type) {
	case *object.Builtin:
		return funcType.Fn(args)
	case *object.Function:
		extendedEnv, err := extendFunctionEnv(funcType, args)
		if err != nil {
			return newError("failed to apply function: %s", err)
		}

		evaluated := Eval(funcType.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	default:
		return newError("not a function: %s", function.Type())
	}
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		fmt.Println("returnValue.Value", returnValue.Value)
		return returnValue.Value
	}

	return obj
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
	variable, ok := varValue.(*ast.StringLiteral)
	if !ok {
		return newError("var key must be SYMBOL, got %s", varValue)
	}
	if !strings.HasPrefix(variable.Value, "$") {
		return newError("var key must not start with $: %s", variable.Value)
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

	_, isFromKeyFound := kvPairs["from"]
	_, isUntilKeyFound := kvPairs["until"]
	if isFromKeyFound && isUntilKeyFound {
		return evalFromUntilLoop(keyValueObj, env)
	} else if (isFromKeyFound && !isUntilKeyFound) || (!isFromKeyFound && isUntilKeyFound) {
		return newError("both from and until keys must be found in loop: %s", keyValueObj)
	}

	if _, ok := kvPairs["in"]; ok {
		return evalInLoop(keyValueObj, env)
	}

	return newError("unknown loop type: %s", keyValueObj)
}

func evalFromUntilLoop(keyValueObj *ast.KeyValueObject, env *object.Environment) object.Object {
	extendedEnv := object.NewEnclosedEnvironment(env)
	kvPairs := keyValueObj.KVPairs()

	forValue, ok := kvPairs["for"]
	if !ok {
		return newError("for key not found in loop: %s", keyValueObj)
	}
	loopSymbol, ok := forValue.(*ast.StringLiteral)
	if !ok {
		return newError("for key must be String, got %s", forValue)
	}
	if !strings.HasPrefix(loopSymbol.Value, "$") {
		return newError("for key must not start with $: %s", loopSymbol.Value)
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

	untilValue, ok := kvPairs["until"]
	if !ok {
		return newError("until key not found in loop: %s", keyValueObj)
	}
	until := Eval(untilValue, env)
	if isError(until) {
		return until
	}
	untilInt, ok := until.(*object.Integer)
	if !ok {
		return newError("until value must be INTEGER, got %s", until.Type())
	}

	doValue, ok := kvPairs["do"]
	if !ok {
		return newError("do key not found in loop: %s", keyValueObj)
	}

	var result object.Object

	for i := fromInt.Value; i < untilInt.Value; i++ {
		extendedEnv.Set(loopSymbol.Value, &object.Integer{Value: i})
		evaluated := Eval(doValue, extendedEnv)
		if isError(evaluated) {
			return evaluated
		}

		if evaluated == Break {
			break
		} else if evaluated == Continue {
			continue
		} else if returnValue, ok := evaluated.(*object.ReturnValue); ok {
			return returnValue
		}

		result = evaluated
	}

	return result
}

func evalInLoop(keyValueObj *ast.KeyValueObject, env *object.Environment) object.Object {
	extendedEnv := object.NewEnclosedEnvironment(env)
	kvPairs := keyValueObj.KVPairs()

	forValue, ok := kvPairs["for"]
	if !ok {
		return newError("for key not found in loop: %s", keyValueObj)
	}
	loopSymbol, ok := forValue.(*ast.StringLiteral)
	if !ok {
		return newError("for key must be String, got %s", forValue)
	}
	if !strings.HasPrefix(loopSymbol.Value, "$") {
		return newError("for key must not start with $: %s", loopSymbol.Value)
	}

	doValue, ok := kvPairs["do"]
	if !ok {
		return newError("do key not found in loop: %s", keyValueObj)
	}

	var result object.Object

	inValue, ok := kvPairs["in"]
	if !ok {
		return newError("in key not found in loop: %s", keyValueObj)
	}
	inArray, ok := inValue.(*ast.Array)
	if ok {
		for _, el := range inArray.Elements {
			evaluatedElement := Eval(el, env)
			if isError(evaluatedElement) {
				return evaluatedElement
			}

			extendedEnv.Set(loopSymbol.Value, evaluatedElement)
			evaluated := Eval(doValue, extendedEnv)
			if isError(evaluated) {
				return evaluated
			}

			result = evaluated
		}

		return result
	}

	symbol, ok := inValue.(*ast.StringLiteral)
	if !ok {
		return newError("in key must be ARRAY or SYMBOL, got %s", inValue)
	}
	if !strings.HasPrefix(symbol.Value, "$") {
		return newError("in key must not start with $: %s", symbol.Value)
	}

	evaluatedInSymbol := Eval(symbol, env)
	if isError(evaluatedInSymbol) {
		return evaluatedInSymbol
	}

	arrayObj, ok := evaluatedInSymbol.(*object.Array)
	if !ok {
		return newError("in key must be ARRAY or SYMBOL, got %s", evaluatedInSymbol.Type())
	}

	for _, el := range arrayObj.Elements {
		extendedEnv.Set(loopSymbol.Value, el)
		evaluated := Eval(doValue, extendedEnv)
		if isError(evaluated) {
			return evaluated
		}

		if evaluated == Break {
			break
		} else if evaluated == Continue {
			continue
		} else if returnValue, ok := evaluated.(*object.ReturnValue); ok {
			return returnValue
		}

		result = evaluated
	}

	return result
}

func evalLambdaExpression(exp ast.Expression, env *object.Environment) object.Object {
	keyValueObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return newError("invalid value for defun: %s", exp)
	}
	kvPairs := keyValueObj.KVPairs()

	params := make([]*ast.StringLiteral, 0)
	paramsValue, ok := kvPairs["params"]
	if ok {
		if paramsArray, ok := paramsValue.(*ast.Array); ok {
			for _, param := range paramsArray.Elements {
				paramSymbol, ok := param.(*ast.StringLiteral)
				if !ok {
					return newError("params key must be ARRAY of SYMBOL, got %s", param)
				}
				if !strings.HasPrefix(paramSymbol.Value, "$") {
					return newError("params key must not start with $: %s", paramSymbol.Value)
				}
				params = append(params, paramSymbol)
			}
		} else {
			paramsSymbol, ok := paramsValue.(*ast.StringLiteral)
			if !ok {
				return newError("params key must be ARRAY or SYMBOL, got %s", paramsValue)
			}
			if !strings.HasPrefix(paramsSymbol.Value, "$") {
				return newError("params key must not start with $: %s", paramsSymbol.Value)
			}
			params = append(params, paramsSymbol)
		}
	}

	body, ok := kvPairs["body"]
	if !ok {
		return newError("body key not found in defun: %s", keyValueObj)
	}

	return &object.Function{
		Parameters: params,
		Body:       body,
		Env:        env,
	}
}

func evalEmbeddedIdentifiers(strLiteral *ast.StringLiteral, env *object.Environment, matches []string) object.Object {
	evaluatedIdents := make([]object.Object, 0)
	for _, match := range matches {
		identName := strings.TrimSpace(strings.Trim(match, "{}")) // remove leading/trailing spaces and {}
		ident, isFound := env.Get(identName)
		if !isFound {
			return newError("identifier not found: %s", identName)
		}
		evaluatedIdents = append(evaluatedIdents, ident)
	}

	newStrValue := strLiteral.Value
	for i, match := range matches {
		newStrValue = strings.Replace(newStrValue, match, evaluatedIdents[i].Inspect(), 1)
	}

	return &object.String{Value: newStrValue}
}
