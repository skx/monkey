package object

import "fmt"

// Integer wraps int64 and implements Object and Hashable interfaces.
type Integer struct {
	Value int64
	Const bool
}

func (i *Integer) Constant() bool {
	return i.Const
}
func (i *Integer) SetConstant(val bool) {
	i.Const = val
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
