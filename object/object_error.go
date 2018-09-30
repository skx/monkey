package object

// Error wraps string and implements Object interface.
type Error struct {
	// Message contains the error-message we're wrapping
	Message string

	// Const is true if this object is constant.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (e *Error) SetConstant(val bool) {
	e.Const = val
}

// Constant returns true if an object is read-only or constant.
func (e *Error) Constant() bool {
	return e.Const
}

// Type returns the type of this object.
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect returns a string-representation of the given object.
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (e *Error) InvokeMethod(method string, env Environment, args ...Object) Object {

	//
	// There are no methods available upon a return-object.
	//
	// (The error-object is an implementation-detail.)
	//
	return nil
}
