package object

import (
	"hash/fnv"
	"sort"
	"strconv"
	"strings"
)

// Float wraps float64 and implements Object and Hashable interfaces.
type Float struct {
	// Value holds the float-value this object wraps.
	Value float64
}

// Inspect returns a string-representation of the given object.
func (f *Float) Inspect() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

// Type returns the type of this object.
func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

// HashKey returns a hash key for the given object.
func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(f.Inspect()))
	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (f *Float) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"methods", "string", "type"}
		dynamic := env.Names("float.")

		var names []string
		for _, e := range static {
			names = append(names, e)
		}
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "string" {
		return &String{Value: f.Inspect()}
	}
	if method == "type" {
		return &String{Value: "float"}
	}
	return nil
}
