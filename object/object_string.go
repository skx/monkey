package object

import "hash/fnv"

// String wraps string and implements Object and Hashable interfaces.
type String struct {
	Value string
	Const bool
}

func (s *String) SetConstant(val bool) {
	s.Const = val
}
func (s *String) Constant() bool {
	return s.Const
}
func (s *String) Type() ObjectType {
	return STRING_OBJ
}
func (s *String) Inspect() string {
	return s.Value
}
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
