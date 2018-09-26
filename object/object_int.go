package object

import "fmt"

// Integer wraps int64 and implements Object and Hashable interfaces.
type Integer struct {
	// Value holds the integer value this object wraps
	Value int64

	// Const determines whether this is a constant value
	Const bool
}

// Constant returns true if an object is read-only or constant.
func (i *Integer) Constant() bool {
	return i.Const
}

// SetConstant allows an object to be marked as read-only, or constant.
func (i *Integer) SetConstant(val bool) {
	i.Const = val
}

// Inspect returns a string-representation of the given object.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the type of this object.
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// HashKey returns a hash key for the given object.
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (i *Integer) InvokeMethod(method string, args ...Object) Object {

	if method == "methods" {
		names := []string{"methods", "string", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "string" {
		return &String{Value: i.Inspect()}
	}
	if method == "type" {
		return &String{Value: "integer"}
	}
	return nil
}
