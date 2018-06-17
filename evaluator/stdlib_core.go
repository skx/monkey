package evaluator

import (
	"fmt"
	"monkey/object"
	"strconv"
)

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
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got=%s",
			args[0].Type())
	}
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
	registerBuiltin("int",
		func(args ...object.Object) object.Object {
			return (intFun(args...))
		})
	registerBuiltin("last",
		func(args ...object.Object) object.Object {
			return (lastFun(args...))
		})
	registerBuiltin("len",
		func(args ...object.Object) object.Object {
			return (lenFun(args...))
		})
	registerBuiltin("first",
		func(args ...object.Object) object.Object {
			return (firstFun(args...))
		})
	registerBuiltin("push",
		func(args ...object.Object) object.Object {
			return (pushFun(args...))
		})
	registerBuiltin("puts",
		func(args ...object.Object) object.Object {
			return (putsFun(args...))
		})
	registerBuiltin("rest",
		func(args ...object.Object) object.Object {
			return (restFun(args...))
		})
	registerBuiltin("set",
		func(args ...object.Object) object.Object {
			return (setFun(args...))
		})
	registerBuiltin("string",
		func(args ...object.Object) object.Object {
			return (strFun(args...))
		})
	registerBuiltin("type",
		func(args ...object.Object) object.Object {
			return (typeFun(args...))
		})
}
