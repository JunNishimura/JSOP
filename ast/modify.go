package ast

type ModifierFun func(Expression) Expression

func Modify(exp Expression, modifier ModifierFun) Expression {
	switch e := exp.(type) {
	case *PrefixAtom:
		e.Right = Modify(e.Right, modifier)
	case *Array:
		for i, el := range e.Elements {
			e.Elements[i] = Modify(el, modifier)
		}
	case *KeyValueObject:
		newKV := make([]*KeyValuePair, len(e.KV))
		for i, kv := range e.KV {
			newKeyExp := Modify(kv.Key, modifier)
			newKey, ok := newKeyExp.(*StringLiteral)
			if !ok {
				panic("key must be StringLiteral")
			}
			newValue := Modify(kv.Value, modifier)
			newKV[i] = &KeyValuePair{Key: newKey, Value: newValue}
		}
		e.KV = newKV
	}

	return modifier(exp)
}
