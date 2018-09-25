package object

type BuiltinFunction func(args ...Object) Object

// Builtin wraps func and implements Object interface.

type Builtin struct {
	Fn    BuiltinFunction
	Const bool
}

func (b *Builtin) SetConstant(val bool) {
	b.Const = val
}
func (b *Builtin) Constant() bool {
	return b.Const
}
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
func (b *Builtin) Inspect() string {
	return "builtin function"
}
