// Package object contains our core-definitions for objects.
package object

// ObjectType describes the type of an object.
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

// Object is the interface that all of our various object-types must implmenet.
type Object interface {

	// Type returns the type of this object.
	Type() ObjectType

	// Inspect returns a string-representation of the given object.
	Inspect() string

	// SetConstant allows an object to be marked as read-only, or constant.
	SetConstant(val bool)

	// Constant returns true if an object is read-only or constant.
	Constant() bool

	// InvokeMethod invokes a method against the object.
	// (Built-in methods only.)
	InvokeMethod(method string, args ...Object) Object
}

// Hashable type can be hashed
type Hashable interface {

	// HashKey returns a hash key for the given object.
	HashKey() HashKey
}
