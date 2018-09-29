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
	if method == "count" {
		if len(args) < 1 {
			return &Error{Message: "Missing argument to count()!"}
		}
		// Note that this coerces into a string :)
		arg := args[0].Inspect()
		return &Integer{Value: int64(strings.Count(s.Value, arg))}
	}
	if method == "find" {
		if len(args) < 1 {
			return &Error{Message: "Missing argument to find()!"}
		}

		// Note that this coerces into a string :)
		arg := args[0].Inspect()
		return &Integer{Value: int64(strings.Index(s.Value, arg))}
	}
	if method == "len" {
		return &Integer{Value: int64(utf8.RuneCountInString(s.Value))}
	}
	if method == "methods" {
		names := []string{"count", "find", "len", "methods", "replace", "split", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "ord" {
		return &Integer{Value: int64(s.Value[0])}
	}
	if method == "replace" {
		if len(args) < 2 {
			return &Error{Message: "Missing arguments to replace()!"}
		}
		// Note that this coerces into strings :)
		oldS := args[0].Inspect()
		newS := args[1].Inspect()
		return &String{Value: strings.Replace(s.Value, oldS, newS, -1)}
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
	if method == "type" {
		return &String{Value: "string"}
	}
	return nil
}
