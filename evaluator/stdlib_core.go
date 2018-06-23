package evaluator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/skx/monkey/object"
)

// exit a program.
func exitFun(args ...object.Object) object.Object {

	code := 0

	// Optionally an exit-code might be supplied as an argument
	if len(args) > 0 {
		switch arg := args[0].(type) {
		case *object.Integer:
			code = int(arg.Value)
		case *object.Float:
			code = int(arg.Value)
		}
	}

	os.Exit(code)
	return NULL
}

// convert a double/string to an int
func intFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch args[0].(type) {
	case *object.String:
		input := args[0].(*object.String).Value
		i, err := strconv.Atoi(input)
		if err == nil {
			return &object.Integer{Value: int64(i)}
		} else {
			return newError("Converting string '%s' to int failed %s", input, err.Error())
		}
	case *object.Boolean:
		input := args[0].(*object.Boolean).Value
		if input {
			return &object.Integer{Value: 1}

		}
		return &object.Integer{Value: 0}
	case *object.Integer:
		// nop
		return args[0]
	case *object.Float:
		input := args[0].(*object.Float).Value
		return &object.Integer{Value: int64(input)}
	default:
		return newError("argument to `int` not supported, got=%s",
			args[0].Type())
	}
}

// the first item in an array.
func firstFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[0]
	}
	return NULL
}

// last item in an array
func lastFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}
	return NULL
}

// length of item
func lenFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(utf8.RuneCountInString(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got=%s",
			args[0].Type())
	}
}

// regular expression match
func matchFun(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	if args[0].Type() != object.STRING_OBJ {
		return newError("argument to `match` must be STRING, got %s",
			args[0].Type())
	}
	if args[1].Type() != object.STRING_OBJ {
		return newError("argument to `match` must be STRING, got %s",
			args[1].Type())
	}

	//
	// Compile and match
	//
	reg := regexp.MustCompile(args[0].(*object.String).Value)
	res := reg.FindStringSubmatch(args[1].(*object.String).Value)

	if len(res) > 0 {

		newHash := make(map[object.HashKey]object.HashPair)

		//
		// If we get a match then the output is an array
		// First entry is the match, any additional parts
		// are the capture-groups.
		//
		if len(res) > 1 {
			for i := 1; i < len(res); i++ {

				// Capture groups start at index 0.
				k := &object.Integer{Value: int64(i - 1)}
				v := &object.String{Value: res[i]}

				newHashPair := object.HashPair{Key: k, Value: v}
				newHash[k.HashKey()] = newHashPair

			}
		}

		return &object.Hash{Pairs: newHash}
	}

	// No match
	return NULL
}

// set a global pragam
func pragmaFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch args[0].(type) {
	case *object.String:
		input := args[0].(*object.String).Value
		PRAGMAS[input] = 1
	default:
		return newError("argument to `pragma` not supported, got=%s",
			args[0].Type())
	}

	return &object.Boolean{Value: true}
}

// push something onto an array
func pushFun(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got=%s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]
	return &object.Array{Elements: newElements}
}

// output a string to stdout
func putsFun(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Print(arg.Inspect())
	}
	return NULL
}

// rest of an array.
func restFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got=%s",
			args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}
	return NULL

}

// Get hash keys
func hashKeys(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.HASH_OBJ {
		return newError("argument to `keys` must be HASH, got=%s",
			args[0].Type())
	}

	// The object we're working with
	hash := args[0].(*object.Hash)
	ents := len(hash.Pairs)

	// Create a new array for the results.
	array := make([]object.Object, ents, ents)

	// Now copy the keys into it.
	i := 0
	for _, ent := range hash.Pairs {
		array[i] = ent.Key
		i++
	}

	// Return the array.
	return &object.Array{Elements: array}
}

// Delete a given hash-key
func hashDelete(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.HASH_OBJ {
		return newError("argument to `delete` must be HASH, got=%s",
			args[0].Type())
	}

	// The object we're working with
	hash := args[0].(*object.Hash)

	// The key we're going to delete
	key, ok := args[1].(object.Hashable)
	if !ok {
		return newError("key `delete` into HASH must be Hashable, got=%s",
			args[1].Type())
	}

	// Make a new hash
	newHash := make(map[object.HashKey]object.HashPair)

	// Copy the values EXCEPT the one we have.
	for k, v := range hash.Pairs {
		if k != key.HashKey() {
			newHash[k] = v
		}
	}
	return &object.Hash{Pairs: newHash}
}

// set a hash-field
func setFun(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.HASH_OBJ {
		return newError("argument to `set` must be HASH, got=%s",
			args[0].Type())
	}
	key, ok := args[1].(object.Hashable)
	if !ok {
		return newError("key `set` into HASH must be Hashable, got=%s",
			args[1].Type())
	}
	newHash := make(map[object.HashKey]object.HashPair)
	hash := args[0].(*object.Hash)
	for k, v := range hash.Pairs {
		newHash[k] = v
	}
	newHashKey := key.HashKey()
	newHashPair := object.HashPair{Key: args[1], Value: args[2]}
	newHash[newHashKey] = newHashPair
	return &object.Hash{Pairs: newHash}
}

func strFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	out := args[0].Inspect()
	return &object.String{Value: out}
}

// type of an item
func typeFun(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch args[0].(type) {
	case *object.String:
		return &object.String{Value: "string"}
	case *object.Boolean:
		return &object.String{Value: "bool"}
	case *object.Array:
		return &object.String{Value: "array"}
	case *object.Function:
		return &object.String{Value: "function"}
	case *object.Integer:
		return &object.String{Value: "integer"}
	case *object.Float:
		return &object.String{Value: "float"}
	case *object.Hash:
		return &object.String{Value: "hash"}
	default:
		return newError("argument to `type` not supported, got=%s",
			args[0].Type())
	}
}
func init() {
	RegisterBuiltin("delete",
		func(args ...object.Object) object.Object {
			return (hashDelete(args...))
		})
	RegisterBuiltin("exit",
		func(args ...object.Object) object.Object {
			return (exitFun(args...))
		})
	RegisterBuiltin("int",
		func(args ...object.Object) object.Object {
			return (intFun(args...))
		})
	RegisterBuiltin("keys",
		func(args ...object.Object) object.Object {
			return (hashKeys(args...))
		})
	RegisterBuiltin("last",
		func(args ...object.Object) object.Object {
			return (lastFun(args...))
		})
	RegisterBuiltin("len",
		func(args ...object.Object) object.Object {
			return (lenFun(args...))
		})
	RegisterBuiltin("match",
		func(args ...object.Object) object.Object {
			return (matchFun(args...))
		})
	RegisterBuiltin("first",
		func(args ...object.Object) object.Object {
			return (firstFun(args...))
		})
	RegisterBuiltin("pragma",
		func(args ...object.Object) object.Object {
			return (pragmaFun(args...))
		})
	RegisterBuiltin("push",
		func(args ...object.Object) object.Object {
			return (pushFun(args...))
		})
	RegisterBuiltin("puts",
		func(args ...object.Object) object.Object {
			return (putsFun(args...))
		})
	RegisterBuiltin("rest",
		func(args ...object.Object) object.Object {
			return (restFun(args...))
		})
	RegisterBuiltin("set",
		func(args ...object.Object) object.Object {
			return (setFun(args...))
		})
	RegisterBuiltin("string",
		func(args ...object.Object) object.Object {
			return (strFun(args...))
		})
	RegisterBuiltin("type",
		func(args ...object.Object) object.Object {
			return (typeFun(args...))
		})
}
