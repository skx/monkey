package object

// Error wraps string and implements Object interface.
type Error struct {
	Message string
	Const   bool
}

func (e *Error) SetConstant(val bool) {
	e.Const = val
}
func (e *Error) Constant() bool {
	return e.Const
}
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
