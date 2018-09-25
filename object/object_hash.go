package object

import (
	"bytes"
	"fmt"
	"strings"
)

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

// Hash wrap map[HashKey]HashPair and implements Object interface.
type Hash struct {
	Pairs map[HashKey]HashPair
	Const bool
}

func (h *Hash) SetConstant(val bool) {
	h.Const = val
}
func (h *Hash) Constant() bool {
	return h.Const
}
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := make([]string, 0)
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
