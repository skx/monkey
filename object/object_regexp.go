// The implementation of our regular-expression object.

package object

// Regexp wraps regular-expressions and implements Object and Hashable interfaces.
type Regexp struct {
	// Value holds the string value this object wraps.
	Value string
}

// Type returns the type of this object.
func (s *Regexp) Type() Type {
	return REGEXP_OBJ
}

// Inspect returns a string-representation of the given object.
func (r *Regexp) Inspect() string {
	return r.Value
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (s *Regexp) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}
