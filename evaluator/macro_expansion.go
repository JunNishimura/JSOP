package evaluator

import (
	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
)

func DefineMacros(program ast.Expression, env *object.Environment) ast.Expression {
	if arrayExp, ok := program.(*ast.Array); ok {
		definitions := make([]int, 0)

		for i, exp := range arrayExp.Elements {
			if isMacroDefinition(exp) {
				addMacro(exp, env)
				definitions = append(definitions, i)
			}
		}

		for i := len(definitions) - 1; i >= 0; i = i - 1 {
			definitionIndex := definitions[i]
			arrayExp.Elements = append(
				arrayExp.Elements[:definitionIndex],
				arrayExp.Elements[definitionIndex+1:]...,
			)
		}

		return arrayExp
	}

	if isMacroDefinition(program) {
		addMacro(program, env)
	}

	return nil
}

func isMacroDefinition(exp ast.Expression) bool {
	kvObject, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return false
	}

	_, ok = kvObject.KVPairs()["defmacro"]
	return ok
}

func addMacro(exp ast.Expression, env *object.Environment) {
	kvObject, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return
	}

	defmacro, ok := kvObject.KVPairs()["defmacro"]
	if !ok {
		return
	}
	defmacroValue, ok := defmacro.(*ast.KeyValueObject)
	if !ok {
		return
	}
	kvPairs := defmacroValue.KVPairs()

	nameValue, ok := kvPairs["name"]
	if !ok {
		return
	}
	macroName, ok := nameValue.(*ast.StringLiteral)
	if !ok {
		return
	}

	keys := make([]*ast.StringLiteral, 0)
	keysValue, ok := kvPairs["keys"]
	if ok {
		if keysArray, ok := keysValue.(*ast.Array); ok {
			for _, key := range keysArray.Elements {
				strLiteral, ok := key.(*ast.StringLiteral)
				if !ok {
					return
				}
				keys = append(keys, strLiteral)
			}
		} else {
			strLiteral, ok := keysValue.(*ast.StringLiteral)
			if !ok {
				return
			}
			keys = append(keys, strLiteral)
		}
	}

	bodyValue, ok := kvPairs["body"]
	if !ok {
		return
	}

	macroObj := &object.Macro{
		Keys: keys,
		Body: bodyValue,
		Env:  env,
	}

	env.Set(macroName.Value, macroObj)
}

func ExpandMacros(program ast.Expression, env *object.Environment) ast.Expression {
	return ast.Modify(program, func(exp ast.Expression) ast.Expression {
		kvObject, ok := exp.(*ast.KeyValueObject)
		if !ok {
			return exp
		}

		macroName, ok := isMacroCall(kvObject, env)
		if !ok {
			return exp
		}
		macroObj, ok := env.Get(macroName)
		if !ok {
			return exp
		}
		macro, ok := macroObj.(*object.Macro)
		if !ok {
			return exp
		}

		macroContentValue, ok := kvObject.KVPairs()[macroName]
		if !ok {
			return exp
		}
		macroContent, ok := macroContentValue.(*ast.KeyValueObject)
		if !ok {
			return exp
		}
		macroContentKV := macroContent.KVPairs()

		newKV := map[string]*object.Quote{}
		for _, key := range macro.Keys {
			keyValue, ok := macroContentKV[key.Value]
			if !ok {
				return exp
			}
			newKV[key.Value] = &object.Quote{Expression: keyValue}
		}

		macroEnv := extendMacroEnv(macro, newKV)

		evaluated := Eval(macro.Body, macroEnv)

		quote, ok := evaluated.(*object.Quote)
		if !ok {
			return exp
		}

		return quote.Expression
	})
}

func extendMacroEnv(macro *object.Macro, kv map[string]*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)
	for _, key := range macro.Keys {
		extended.Set(key.Value, kv[key.Value])
	}
	return extended
}

func isMacroCall(kvObj *ast.KeyValueObject, env *object.Environment) (string, bool) {
	for key := range kvObj.KVPairs() {
		obj, ok := env.Get(key)
		if !ok {
			continue
		}

		_, ok = obj.(*object.Macro)
		if !ok {
			continue
		}

		return key, true
	}

	return "", false
}
