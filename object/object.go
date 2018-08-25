package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"

	"github.com/skx/monkey/ast"
)

type ObjectType string

// pre-defined constant ObjectType
const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

type Object interface {
	Type() ObjectType
	Inspect() string
	SetConstant(val bool)
	Constant() bool
}

// Hashable type can be hashed
type Hashable interface {
	HashKey() HashKey
}

// Integer wraps int64 and implements Object and Hashable interfaces.
type Integer struct {
	Value int64
	Const bool
}

func (i *Integer) Constant() bool       { return i.Const }
func (i *Integer) SetConstant(val bool) { i.Const = val }
func (i *Integer) Inspect() string      { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType     { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// Float wraps float64 and implements Object and Hashable interfaces.
type Float struct {
	Value float64
	Const bool
}

func (i *Float) SetConstant(val bool) { i.Const = val }
func (i *Float) Constant() bool       { return i.Const }
func (f *Float) Inspect() string      { return strconv.FormatFloat(f.Value, 'f', -1, 64) }
func (f *Float) Type() ObjectType     { return FLOAT_OBJ }
func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(f.Inspect()))
	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

// Boolean wraps bool and implements Object and Hashable interface.
type Boolean struct {
	Value bool
	Const bool
}

func (b *Boolean) SetConstant(val bool) { b.Const = val }
func (b *Boolean) Constant() bool       { return b.Const }
func (b *Boolean) Type() ObjectType     { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string      { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// Null wraps nothing and implements Object interface.
type Null struct{}

func (n *Null) SetConstant(val bool) {}
func (n *Null) Constant() bool       { return false }
func (n *Null) Type() ObjectType     { return NULL_OBJ }
func (n *Null) Inspect() string      { return "null" }

// ReturnValue wraps Object and implements Object interface.
type ReturnValue struct {
	Value Object
	Const bool
}

func (rv *ReturnValue) SetConstant(val bool) { rv.Const = val }
func (rv *ReturnValue) Constant() bool       { return rv.Const }
func (rv *ReturnValue) Type() ObjectType     { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string      { return rv.Value.Inspect() }

// Error wraps string and implements Object interface.
type Error struct {
	Message string
	Const   bool
}

func (e *Error) SetConstant(val bool) { e.Const = val }
func (e *Error) Constant() bool       { return e.Const }
func (e *Error) Type() ObjectType     { return ERROR_OBJ }
func (e *Error) Inspect() string      { return "ERROR: " + e.Message }

// Function wraps ast.Identifier array, ast.BlockStatement and Environment and implements Object interface.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Defaults   map[string]ast.Expression
	Env        *Environment
	Const      bool
}

func (f *Function) SetConstant(val bool) { f.Const = val }
func (f *Function) Constant() bool       { return f.Const }
func (f *Function) Type() ObjectType     { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	parameters := make([]string, 0)
	for _, p := range f.Parameters {
		parameters = append(parameters, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(parameters, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// String wraps string and implements Object and Hashable interfaces.
type String struct {
	Value string
	Const bool
}

func (s *String) SetConstant(val bool) { s.Const = val }
func (s *String) Constant() bool       { return s.Const }
func (s *String) Type() ObjectType     { return STRING_OBJ }
func (s *String) Inspect() string      { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type BuiltinFunction func(args ...Object) Object

// Builtin wraps func and implements Object interface.

type Builtin struct {
	Fn    BuiltinFunction
	Const bool
}

func (b *Builtin) SetConstant(val bool) { b.Const = val }
func (b *Builtin) Constant() bool       { return b.Const }
func (b *Builtin) Type() ObjectType     { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string      { return "builtin function" }

// Array wraps Object array and implements Object interface.
type Array struct {
	Elements []Object
	Const    bool
}

func (ao *Array) SetConstant(val bool) { ao.Const = val }
func (ao *Array) Constant() bool       { return ao.Const }
func (ao *Array) Type() ObjectType     { return ARRAY_OBJ }
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

func (h *Hash) SetConstant(val bool) { h.Const = val }
func (h *Hash) Constant() bool       { return h.Const }
func (h *Hash) Type() ObjectType     { return HASH_OBJ }
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
