package object

import "fmt"

// Boolean wraps bool and implements Object and Hashable interface.
type Boolean struct {
	Value bool
	Const bool
}

func (b *Boolean) SetConstant(val bool) {
	b.Const = val
}
func (b *Boolean) Constant() bool {
	return b.Const
}
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}
