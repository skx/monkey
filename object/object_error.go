package object

// Error wraps string and implements Object interface.
type Error struct {
	// Message contains the error-message we're wrapping
	Message string
}

// Type returns the type of this object.
func (e *Error) Type() Type {
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
