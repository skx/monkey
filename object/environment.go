package object

import (
	"math"
	"os"
	"strings"
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
	// Mathematical constants
	e.Set("PI", &Float{Value: math.Pi})
	e.Set("E", &Float{Value: math.E})

	// For read/write operations
	e.Set("STDIN", &Integer{Value: 0})
	e.Set("STDOUT", &Integer{Value: 1})
	e.Set("STDERR", &Integer{Value: 2})

	// Setup each environmental variable.
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		e.Set("$"+pair[0], &String{Value: pair[1]})
	}
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
