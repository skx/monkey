package object

// Null wraps nothing and implements our Object interface.
type Null struct{}

// Type returns the type of this object.
func (n *Null) Type() Type {
	return NULL_OBJ
}

// Inspect returns a string-representation of the given object.
func (n *Null) Inspect() string {
	return "null"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (n *Null) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}
