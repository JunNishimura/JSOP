package evaluator

import (
	"fmt"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/object"
)

func DefineMacros(program ast.Expression, env *object.Environment) error {
	if arrayExp, ok := program.(*ast.Array); ok {
		definitions := make([]int, 0)

		for i, exp := range arrayExp.Elements {
			if macro, ok := isMacroDefinition(exp); ok {
				if err := addMacro(macro, env); err != nil {
					return err
				}
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
	}

	if macro, ok := isMacroDefinition(program); ok {
		if err := addMacro(macro, env); err != nil {
			return err
		}
	}

	return nil
}

func isMacroDefinition(exp ast.Expression) (*ast.KeyValueObject, bool) {
	kvObj, ok := exp.(*ast.KeyValueObject)
	if !ok {
		return nil, false
	}

	defmacroVal, ok := kvObj.KVPairs()["defmacro"]
	if !ok {
		return nil, false
	}
	defmacro, ok := defmacroVal.(*ast.KeyValueObject)
	if !ok {
		return nil, false
	}

	return defmacro, true
}

func addMacro(macro *ast.KeyValueObject, env *object.Environment) error {
	kvPairs := macro.KVPairs()

	nameVal, ok := kvPairs["name"]
	if !ok {
		return fmt.Errorf("macro expects 'name' key")
	}
	macroName, ok := nameVal.(*ast.StringLiteral)
	if !ok {
		return fmt.Errorf("macro expects 'name' key to be StringLiteral, got %t", nameVal)
	}

	keys := make([]*ast.StringLiteral, 0)
	keysValue, ok := kvPairs["keys"]
	if ok {
		switch keysValue := keysValue.(type) {
		case *ast.Array:
			for _, key := range keysValue.Elements {
				strLiteral, ok := key.(*ast.StringLiteral)
				if !ok {
					return fmt.Errorf("macro expects 'keys' to be Array of StringLiterals")
				}
				keys = append(keys, strLiteral)
			}
		case *ast.StringLiteral:
			keys = append(keys, keysValue)
		default:
			return fmt.Errorf("macro expects 'keys' to be Array or StringLiteral")
		}
	}

	bodyVal, ok := kvPairs["body"]
	if !ok {
		return fmt.Errorf("macro expects 'body' key")
	}

	macroObj := &object.Macro{
		Keys: keys,
		Body: bodyVal,
		Env:  env,
	}
	env.Set(macroName.Value, macroObj)

	return nil
}

func ExpandMacros(program ast.Expression, env *object.Environment) ast.Expression {
	return ast.Modify(program, func(exp ast.Expression) ast.Expression {
		kvObj, ok := exp.(*ast.KeyValueObject)
		if !ok {
			return exp
		}

		macroName, macroObj, ok := isMacroCall(kvObj, env)
		if !ok {
			return exp
		}

		body, ok := kvObj.KVPairs()[macroName]
		if !ok {
			return exp
		}
		bodyObj, ok := body.(*ast.KeyValueObject)
		if !ok {
			return exp
		}

		quotedKeys, ok := quoteKeys(macroObj, bodyObj)
		if !ok {
			return exp
		}

		macroEnv := extendMacroEnv(macroObj, quotedKeys)

		evaluated := Eval(macroObj.Body, macroEnv)

		quote, ok := evaluated.(*object.Quote)
		if !ok {
			return exp
		}

		return quote.Expression
	})
}

func quoteKeys(macroObj *object.Macro, body *ast.KeyValueObject) (map[string]*object.Quote, bool) {
	quotedKeys := make(map[string]*object.Quote)

	kvPairs := body.KVPairs()
	for _, key := range macroObj.Keys {
		value, ok := kvPairs[key.Value]
		if !ok {
			return nil, false
		}
		quotedKeys[key.Value] = &object.Quote{Expression: value}
	}

	return quotedKeys, true
}

func extendMacroEnv(macro *object.Macro, quotedKeys map[string]*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)
	for _, key := range macro.Keys {
		extended.Set(key.Value, quotedKeys[key.Value])
	}
	return extended
}

func isMacroCall(kvObj *ast.KeyValueObject, env *object.Environment) (string, *object.Macro, bool) {
	for key := range kvObj.KVPairs() {
		keyObj, ok := env.Get(key)
		if !ok {
			continue
		}

		macroObj, ok := keyObj.(*object.Macro)
		if !ok {
			continue
		}

		return key, macroObj, true
	}

	return "", nil, false
}
