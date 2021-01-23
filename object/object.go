// Package object contains our core-definitions for objects.
package object

// Type describes the type of an object.
type Type string

// pre-defined constant Type
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
	FILE_OBJ         = "FILE"
	REGEXP_OBJ       = "REGEXP"
)

var (
	SystemTypes = []Type{
		INTEGER_OBJ,
		FLOAT_OBJ,
		BOOLEAN_OBJ,
		NULL_OBJ,
		RETURN_VALUE_OBJ,
		ERROR_OBJ,
		FUNCTION_OBJ,
		STRING_OBJ,
		BUILTIN_OBJ,
		ARRAY_OBJ,
		HASH_OBJ,
		FILE_OBJ,
		REGEXP_OBJ,
	}

	SystemTypesMap = map[Type]Object{
		INTEGER_OBJ:      &Integer{},
		FLOAT_OBJ:        &Float{},
		BOOLEAN_OBJ:      &Boolean{},
		NULL_OBJ:         &Null{},
		RETURN_VALUE_OBJ: &ReturnValue{},
		ERROR_OBJ:        &Error{},
		FUNCTION_OBJ:     &Function{},
		STRING_OBJ:       &String{},
		BUILTIN_OBJ:      &Builtin{},
		ARRAY_OBJ:        &Array{},
		HASH_OBJ:         &Hash{},
		FILE_OBJ:         &File{},
		REGEXP_OBJ:       &Regexp{},
	}
)

// Object is the interface that all of our various object-types must implmenet.
type Object interface {

	// Type returns the type of this object.
	Type() Type

	// Inspect returns a string-representation of the given object.
	Inspect() string

	// GetMethod returns a method against the object.
	// (Built-in methods only.)
	GetMethod(method string) BuiltinFunction

	// ToInterface converts the given object to a "native" golang value,
	// which is required to ensure that we can use the object in our
	// `sprintf` or `printf` primitives.
	ToInterface() interface{}
}

// Hashable type can be hashed
type Hashable interface {

	// HashKey returns a hash key for the given object.
	HashKey() HashKey
}

// Iterable is an interface that some objects might support.
//
// If this interface is implemented then it will be possible to
// use the `foreach` function to iterate over the object.  If
// the interface is not implemented then a run-time error will
// be generated instead.
type Iterable interface {

	// Reset the state of any previous iteration.
	Reset()

	// Get the next "thing" from the object being iterated
	// over.
	//
	// The return values are the item which is to be returned
	// next, the index of that object, and finally a boolean
	// to say whether the function succeeded.
	//
	// If the boolean value returned is false then that
	// means the iteration has completed and no further
	// items are available.
	Next() (Object, Object, bool)
}
