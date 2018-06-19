package object

import (
	"math"
)

type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment create new environment by outer parameter
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Register default variables.
func (e *Environment) RegisterDefaults() {
	e.Set("PI", &Float{Value: math.Pi})
	e.Set("E", &Float{Value: math.E})
}

// Get object by name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set object by name
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
