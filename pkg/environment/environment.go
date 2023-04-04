package environment

// Environment is a linked list of environments
type Environment struct {
	values map[string]interface{}
	parent *Environment
}

// NewEnvironment creates a new environment with an empty values map
func NewEnvironment() *Environment {
	return &Environment{values: make(map[string]interface{}), parent: nil}
}

// Extend creates a new child environment for the current environment
func (env *Environment) Extend() *Environment {
	child := &Environment{values: make(map[string]interface{}), parent: env}
	return child
}

// Set sets the value of a variable in the current environment
func (env *Environment) Set_assign(name string, value interface{}) {
	_, ok := env.values[name]
	if ok {
		env.values[name] = value
		return
	}
	if env.parent != nil {
		env.parent.Set_assign(name, value)
		return
	}
	panic("Symbol not found!")
}

func (env *Environment) Set_declare(name string, value interface{}) {
	env.values[name] = value
}

// Get returns the value of a variable in the current environment or its parent environments
func (env *Environment) Get(name string) (interface{}, bool) {
	value, ok := env.values[name]
	if ok {
		return value, true
	}

	if env.parent != nil {
		return env.parent.Get(name)
	}

	return nil, false
}

var Unassigned = struct {
	tag string
}{
	tag: "unassigned",
}
