package object

// ReturnValue wraps Object and implements Object interface.
type ReturnValue struct {
	Value Object
	Const bool
}

func (rv *ReturnValue) SetConstant(val bool) {
	rv.Const = val
}
func (rv *ReturnValue) Constant() bool {
	return rv.Const
}
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
