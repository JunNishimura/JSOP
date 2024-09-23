package evaluator

import (
	"fmt"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
)

func Eval(exp ast.Expression) object.Object {
	switch expt := exp.(type) {
	case *ast.Program:
		return evalProgram(expt)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: expt.Value}
	case *ast.PrefixAtom:
		right := Eval(expt.Right)
		if isError(right) {
			return right
		}
		return evalPrefixAtom(expt.Operator, right)
	case *ast.Symbol:
		return evalSymbol(expt)
	case *ast.CommandObject:
		return evalCommandObject(expt)
	default:
		return newError("unknown expression type: %T", exp)
	}
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, exp := range program.Expressions {
		result = Eval(exp)
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

func evalSymbol(symbol *ast.Symbol) object.Object {
	builtintFunc, ok := builtins[symbol.Value]
	if ok {
		return builtintFunc
	}

	return newError("symbol not found: %s", symbol.Value)
}

func evalCommandObject(command *ast.CommandObject) object.Object {
	symbol := Eval(command.Symbol)
	if isError(symbol) {
		return symbol
	}

	args := evalArgs(command.Args)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return applyFunction(symbol, args)
}

func evalArgs(args []ast.Expression) []object.Object {
	var result []object.Object

	for _, arg := range args {
		evaluated := Eval(arg)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(function object.Object, args []object.Object) object.Object {
	switch funcType := function.(type) {
	case *object.Builtin:
		return funcType.Fn(args...)
	default:
		return newError("not a function: %s", function.Type())
	}
}
