package object

import "fmt"

// Boolean wraps bool and implements Object and Hashable interface.
type Boolean struct {
	// Value holds the boolean value we wrap.
	Value bool

	// Const determines whether this is a constant value.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (b *Boolean) SetConstant(val bool) {
	b.Const = val
}

// Constant returns true if an object is read-only or constant.
func (b *Boolean) Constant() bool {
	return b.Const
}

// Type returns the type of this object.
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect returns a string-representation of the given object.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// HashKey returns a hash key for the given object.
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (b *Boolean) InvokeMethod(method string, args ...Object) Object {
	if method == "methods" {
		names := []string{"methods", "string"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "string" {
		return &String{Value: b.Inspect()}
	}
	return nil
}
