package object

// BuiltinFunction holds the type of a built-in function.
type BuiltinFunction func(env *Environment, args ...Object) Object

// Builtin wraps func and implements Object interface.
type Builtin struct {
	// Value holds the function we wrap.
	Fn BuiltinFunction
}

// Type returns the type of this object.
func (b *Builtin) Type() Type {
	return BUILTIN_OBJ
}

// Inspect returns a string-representation of the given object.
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// GetMethod returns a method against the object.
// (Built-in methods only.)
func (b *Builtin) GetMethod(method string) BuiltinFunction {
	if method == "methods" {
		return func(env *Environment, args ...Object) Object {
			names := []string{"methods"}

			result := make([]Object, len(names))
			for i, txt := range names {
				result[i] = &String{Value: txt}
			}
			return &Array{Elements: result}
		}
	}
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (b *Builtin) ToInterface() interface{} {
	return "<BUILTIN>"
}
