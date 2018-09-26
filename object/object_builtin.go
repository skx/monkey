package object

// BuiltinFunction holds the type of a built-in function.
type BuiltinFunction func(args ...Object) Object

// Builtin wraps func and implements Object interface.
type Builtin struct {
	// Value holds the function we wrap.
	Fn BuiltinFunction

	// Const is true if this object is constant.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (b *Builtin) SetConstant(val bool) {
	b.Const = val
}

// Constant returns true if an object is read-only or constant.
func (b *Builtin) Constant() bool {
	return b.Const
}

// Type returns the type of this object.
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

// Inspect returns a string-representation of the given object.
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (b *Builtin) InvokeMethod(method string, args ...Object) Object {
	if method == "methods" {
		names := []string{"methods"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	//
	// There are no methods available upon a return-object.
	//
	// (The return-object is an implementation-detail.)
	//
	return nil
}
