package object

import (
	"bytes"
	"fmt"
	"strings"
)

// HashKey is the structure used for hash-keys
type HashKey struct {
	// Type holds the type of the object.
	Type ObjectType

	// Value holds the actual hash-key value.
	Value uint64
}

// HashPair is a structure which is used to store hash-entries
type HashPair struct {
	// Key holds our hash-key key.
	Key Object

	// Value holds our hash-key value.
	Value Object
}

// Hash wrap map[HashKey]HashPair and implements Object interface.
type Hash struct {
	// Pairs holds the key/value pairs of the hash we wrap
	Pairs map[HashKey]HashPair

	// Const determines whether this is a constant value.
	Const bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (h *Hash) SetConstant(val bool) {
	h.Const = val
}

// Constant returns true if an object is read-only or constant.
func (h *Hash) Constant() bool {
	return h.Const
}

// Type returns the type of this object.
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

// Inspect returns a string-representation of the given object.
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

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (h *Hash) InvokeMethod(method string, args ...Object) Object {
	if method == "methods" {
		names := []string{"keys", "methods", "string", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "keys" {
		ents := len(h.Pairs)
		array := make([]Object, ents, ents)

		// Now copy the keys into it.
		i := 0
		for _, ent := range h.Pairs {
			array[i] = ent.Key
			i++
		}

		return &Array{Elements: array}
	}
	if method == "string" {
		return &String{Value: h.Inspect()}
	}
	if method == "type" {
		return &String{Value: "hash"}
	}
	return nil
}
