package evaluator

import "github.com/JunNishimura/jsop/object"

var builtins = map[string]*object.Builtin{
	"+": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '+' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) == 0 {
				return newError("number of arguments to '+' must be more than 0, got %d", len(arrayArg.Elements))
			}

			var result int64
			for _, arg := range arrayArg.Elements {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '+' must be INTEGER, got %s", arg.Type())
				}
				result += arg.(*object.Integer).Value
			}
			return &object.Integer{Value: result}
		},
	},
	"-": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '-' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '-' must be more than 1, got %d", len(arrayArg.Elements))
			}

			var result int64
			for i, arg := range arrayArg.Elements {
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
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '*' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) == 0 {
				return newError("number of arguments to '*' must be more than 0, got %d", len(arrayArg.Elements))
			}

			var result int64 = 1
			for _, arg := range arrayArg.Elements {
				if arg.Type() != object.INTEGER_OBJ {
					return newError("argument to '*' must be INTEGER, got %s", arg.Type())
				}
				result *= arg.(*object.Integer).Value
			}

			return &object.Integer{Value: result}
		},
	},
	"/": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '/' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '/' must be more than 1, got %d", len(arrayArg.Elements))
			}

			var result int64
			for i, arg := range arrayArg.Elements {
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
	"==": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '==' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '=' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				if arrayArg.Elements[i] != arrayArg.Elements[i+1] {
					return False
				}
			}

			return True
		},
	},
	"!=": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '!=' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '!=' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				if arrayArg.Elements[i] != arrayArg.Elements[i+1] {
					return True
				}
			}

			return False
		},
	},
	">": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '>' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '>' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				first, ok := arrayArg.Elements[i].(*object.Integer)
				if !ok {
					return newError("argument to '>' must be INTEGER, got %s", arrayArg.Elements[i].Type())
				}
				second, ok := arrayArg.Elements[i+1].(*object.Integer)
				if !ok {
					return newError("argument to '>' must be INTEGER, got %s", arrayArg.Elements[i+1].Type())
				}
				if first.Value <= second.Value {
					return False
				}
			}

			return True
		},
	},
	"<": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '<' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '<' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				first, ok := arrayArg.Elements[i].(*object.Integer)
				if !ok {
					return newError("argument to '<' must be INTEGER, got %s", arrayArg.Elements[i].Type())
				}
				second, ok := arrayArg.Elements[i+1].(*object.Integer)
				if !ok {
					return newError("argument to '<' must be INTEGER, got %s", arrayArg.Elements[i+1].Type())
				}
				if first.Value >= second.Value {
					return False
				}
			}

			return True
		},
	},
	">=": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '>=' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '>=' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				first, ok := arrayArg.Elements[i].(*object.Integer)
				if !ok {
					return newError("argument to '>=' must be INTEGER, got %s", arrayArg.Elements[i].Type())
				}
				second, ok := arrayArg.Elements[i+1].(*object.Integer)
				if !ok {
					return newError("argument to '>=' must be INTEGER, got %s", arrayArg.Elements[i+1].Type())
				}
				if first.Value < second.Value {
					return False
				}
			}

			return True
		},
	},
	"<=": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to '<=' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) <= 1 {
				return newError("number of arguments to '<=' must be more than 1, got %d", len(arrayArg.Elements))
			}

			for i := 0; i < len(arrayArg.Elements)-1; i++ {
				first, ok := arrayArg.Elements[i].(*object.Integer)
				if !ok {
					return newError("argument to '<=' must be INTEGER, got %s", arrayArg.Elements[i].Type())
				}
				second, ok := arrayArg.Elements[i+1].(*object.Integer)
				if !ok {
					return newError("argument to '<=' must be INTEGER, got %s", arrayArg.Elements[i+1].Type())
				}
				if first.Value > second.Value {
					return False
				}
			}

			return True
		},
	},
	"!": {
		Fn: func(args object.Object) object.Object {
			if args.Type() != object.BOOLEAN_OBJ {
				return newError("argument to '!' must be BOOLEAN, got %s", args.Type())
			}

			return &object.Boolean{Value: !args.(*object.Boolean).Value}
		},
	},
	"at": {
		Fn: func(args object.Object) object.Object {
			arrayArg, ok := args.(*object.Array)
			if !ok {
				return newError("argument to 'at' must be ARRAY, got %s", args.Type())
			}
			if len(arrayArg.Elements) != 2 {
				return newError("number of arguments to 'at' must be 2, got %d", len(arrayArg.Elements))
			}

			variable, ok := arrayArg.Elements[0].(*object.Array)
			if !ok {
				return newError("first argument to 'at' must be ARRAY, got %s", arrayArg.Elements[0].Type())
			}

			index, ok := arrayArg.Elements[1].(*object.Integer)
			if !ok {
				return newError("second argument to 'at' must be INTEGER, got %s", arrayArg.Elements[1].Type())
			}

			if index.Value < 0 || index.Value >= int64(len(variable.Elements)) {
				return newError("index out of range: %d", index.Value)
			}

			return variable.Elements[index.Value]
		},
	},
}
