package object

import (
	"fmt"
	"sort"
	"strings"
)

// Integer wraps int64 and implements Object and Hashable interfaces.
type Integer struct {
	// Value holds the integer value this object wraps
	Value int64
}

// Inspect returns a string-representation of the given object.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the type of this object.
func (i *Integer) Type() Type {
	return INTEGER_OBJ
}

// HashKey returns a hash key for the given object.
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// GetMethod returns a method against the object.
// (Built-in methods only.)
func (i *Integer) GetMethod(method string) BuiltinFunction {
	switch method {
	case "chr":
		return func(env *Environment, args ...Object) Object {
			return &String{Value: string(rune(i.Value))}
		}
	case "methods":
		return func(env *Environment, args ...Object) Object {
			static := []string{"chr", "methods"}
			dynamic := env.Names("integer.")

			var names []string
			names = append(names, static...)

			for _, e := range dynamic {
				bits := strings.Split(e, ".")
				names = append(names, bits[1])
			}
			sort.Strings(names)

			result := make([]Object, len(names))
			for i, txt := range names {
				result[i] = &String{Value: txt}
			}
			return &Array{Elements: result}
		}
	}
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (i *Integer) ToInterface() interface{} {
	return i.Value
}
