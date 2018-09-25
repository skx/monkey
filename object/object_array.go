package object

import (
	"bytes"
	"strings"
)

// Array wraps Object array and implements Object interface.
type Array struct {
	Elements []Object
	Const    bool
}

func (ao *Array) SetConstant(val bool) {
	ao.Const = val
}
func (ao *Array) Constant() bool {
	return ao.Const
}
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := make([]string, 0)
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
