package object

// ReturnValue wraps Object and implements Object interface.
type ReturnValue struct {
	// Value is the object that is to be returned
	Value Object

	// Const is true if this object is constant.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (rv *ReturnValue) SetConstant(val bool) {
	rv.Const = val
}

// Constant returns true if an object is read-only or constant.
func (rv *ReturnValue) Constant() bool {
	return rv.Const
}

// Type returns the type of this object.
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect returns a string-representation of the given object.
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (rv *ReturnValue) InvokeMethod(method string, env Environment, args ...Object) Object {

	//
	// There are no methods available upon a return-object.
	//
	// (The return-object is an implementation-detail.)
	//
	return nil
}
