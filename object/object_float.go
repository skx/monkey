package object

import (
	"hash/fnv"
	"strconv"
)

// Float wraps float64 and implements Object and Hashable interfaces.
type Float struct {
	Value float64
	Const bool
}

func (i *Float) SetConstant(val bool) {
	i.Const = val
}
func (i *Float) Constant() bool {
	return i.Const
}
func (f *Float) Inspect() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}
func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}
func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(f.Inspect()))
	return HashKey{Type: f.Type(), Value: h.Sum64()}
}
