// The implementation of our string-object.

package object

import (
	"hash/fnv"
	"strings"
	"unicode/utf8"
)

// String wraps string and implements Object and Hashable interfaces.
type String struct {
	// Value holds the string value this object wraps.
	Value string

	// Const determines whether this is a constant value
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (s *String) SetConstant(val bool) {
	s.Const = val
}

// Constant returns true if an object is read-only or constant.
func (s *String) Constant() bool {
	return s.Const
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
func (s *String) InvokeMethod(method string, args ...Object) Object {
	if method == "methods" {
		names := []string{"len", "methods", "reverse", "split", "toupper", "tolower", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "len" {
		return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
	}
	if method == "reverse" {
		out := make([]rune, utf8.RuneCountInString(s.Value))
		i := len(out)
		for _, c := range s.Value {
			i--
			out[i] = c
		}
		return &String{Value: string(out)}
	}
	if method == "split" {

		// default seperator
		sep := " "

		if len(args) >= 1 {
			// may be changed.
			sep = args[0].(*String).Value
		}

		// do the split
		fields := strings.Split(s.Value, sep)

		// copy the results to the caller.
		l := len(fields)
		result := make([]Object, l, l)
		for i, txt := range fields {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}

	}
	if method == "trim" {
		return &String{Value: strings.TrimSpace(s.Value)}
	}
	if method == "tolower" {
		return &String{Value: strings.ToLower(s.Value)}
	}
	if method == "toupper" {
		return &String{Value: strings.ToUpper(s.Value)}
	}
	if method == "type" {
		return &String{Value: "string"}
	}
	return nil
}
