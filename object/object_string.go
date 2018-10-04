// The implementation of our string-object.

package object

import (
	"hash/fnv"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// String wraps string and implements Object and Hashable interfaces.
type String struct {
	// Value holds the string value this object wraps.
	Value string
}

// Type returns the type of this object.
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect returns a string-representation of the given object.
func (s *String) Inspect() string {
	return s.Value
}

// HashKey returns a hash key for the given object.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "len" {
		return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
	}
	if method == "methods" {
		static := []string{"len", "methods", "ord", "to_i", "to_f"}
		dynamic := env.Names("string.")

		var names []string
		for _, e := range static {
			names = append(names, e)
		}
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "ord" {
		return &Integer{Value: int64(s.Value[0])}
	}
	if method == "to_i" {
		i, err := strconv.ParseInt(s.Value, 0, 64)
		if err != nil {
			i = 0
		}
		return &Integer{Value: int64(i)}
	}
	if method == "to_f" {
		i, err := strconv.ParseFloat(s.Value, 64)
		if err != nil {
			i = 0
		}
		return &Float{Value: i}
	}
	return nil
}
