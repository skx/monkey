package object

import (
	"bytes"
	"strings"
)

// Array wraps Object array and implements Object interface.
type Array struct {
	// Elements holds the individual members of the array we're wrapping.
	Elements []Object

	// Const is true if this object is constant.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (ao *Array) SetConstant(val bool) {
	ao.Const = val
}

// Constant returns true if an object is read-only or constant.
func (ao *Array) Constant() bool {
	return ao.Const
}

// Type returns the type of this object.
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect returns a string-representation of the given object.
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := make([]string, 0)
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (ao *Array) InvokeMethod(method string, args ...Object) Object {
	if method == "len" {
		return &Integer{Value: int64(len(ao.Elements))}
	}
	if method == "methods" {
		names := []string{"len", "methods", "string"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "string" {
		return &String{Value: ao.Inspect()}
	}
	if method == "type" {
		return &String{Value: "array"}
	}
	return nil
}
