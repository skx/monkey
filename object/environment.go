package object

import (
	"fmt"
	"os"
)

// Environment stores our functions, variables, constants, etc.
type Environment struct {
	// store holds variables, including functions.
	store map[string]Object

	// outer holds any parent environment.  Our env. allows
	// nesting to implement scope.
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

	// If this is a constant - we do nothing - then return
	// the value existing.
	if const_val, ok := e.Get(name); ok {
		if const_val.Constant() {
			fmt.Printf("Attempting to modify constant denied - %s\n", name)
			os.Exit(3)
		}
		return const_val
	}
	e.store[name] = val
	return val
}

// SetConst sets the value of a constant by name.
func (e *Environment) SetConst(name string, val Object) Object {

	// store the value
	e.store[name] = val

	// Mark this as a constant
	e.store[name].SetConstant(true)
	return val
}
