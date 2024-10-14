package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	// check if the variable already exists in the current environment
	// current environment has higher priority than outer environment
	if _, ok := e.store[name]; ok {
		// if it does, update the value
		e.store[name] = val
		return val
	}

	// check if the variable exists in the outer environment
	if e.outer != nil {
		if _, ok := e.outer.Get(name); ok {
			// if it does, update the value
			e.outer.Set(name, val)
			return val
		}
	}

	// create a new variable in the current environment
	e.store[name] = val
	return val
}
