package evaluator

import "github.com/JunNishimura/jsop/object"

var builtins = map[string]*object.Builtin{
	"+": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("number of arguments to '+' must be more than 0, got %d", len(args))
			}

			var result int64
			for _, arg := range args {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '+' must be INTEGER, got %s", arg.Type())
				}
				result += arg.(*object.Integer).Value
			}
			return &object.Integer{Value: result}
		},
	},
	"-": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return newError("number of arguments to '-' must be more than 1, got %d", len(args))
			}

			var result int64
			for i, arg := range args {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '-' must be INTEGER, got %s", arg.Type())
				}
				if i == 0 {
					result = arg.(*object.Integer).Value
				} else {
					result -= arg.(*object.Integer).Value
				}
			}

			return &object.Integer{Value: result}
		},
	},
	"*": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("number of arguments to '*' must be more than 0, got %d", len(args))
			}

			var result int64 = 1
			for _, arg := range args {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '*' must be INTEGER, got %s", arg.Type())
				}
				result *= arg.(*object.Integer).Value
			}

			return &object.Integer{Value: result}
		},
	},
	"/": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return newError("number of arguments to '/' must be more than 1, got %d", len(args))
			}

			var result int64
			for i, arg := range args {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '/' must be INTEGER, got %s", arg.Type())
				}
				if i == 0 {
					result = arg.(*object.Integer).Value
				} else {
					if arg.(*object.Integer).Value == 0 {
						return newError("division by zero")
					}
					result /= arg.(*object.Integer).Value
				}
			}

			return &object.Integer{Value: result}
		},
	},
}
