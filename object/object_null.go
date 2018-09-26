package object

// Null wraps nothing and implements our Object interface.
type Null struct{}

// SetConstant allows an object to be marked as read-only, or constant.
func (n *Null) SetConstant(val bool) {
	// NULL cannot be set to constant
}

// Constant returns true if an object is read-only or constant.
func (n *Null) Constant() bool {
	// NULL cannot be set to constant
	return false
}

// Type returns the type of this object.
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect returns a string-representation of the given object.
func (n *Null) Inspect() string {
	return "null"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (n *Null) InvokeMethod(method string, args ...Object) Object {
	return nil
}
