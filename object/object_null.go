package object

// Null wraps nothing and implements Object interface.
type Null struct{}

func (n *Null) SetConstant(val bool) {}
func (n *Null) Constant() bool {
	return false
}
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}
func (n *Null) Inspect() string {
	return "null"
}
