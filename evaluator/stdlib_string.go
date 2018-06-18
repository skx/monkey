package evaluator

import (
	"monkey/object"
	"regexp"
	"strings"
)

// string = string.interpolate( string, hash );
func stringInterpolate(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("first argument must be a string, got=%s",
			args[0].Type())
	}
	if args[1].Type() != object.HASH_OBJ {
		return newError("second argument must be a string, got=%s",
			args[0].Type())
	}

	str := args[0].(*object.String).Value
	hashObject := args[1].(*object.Hash)

	re := regexp.MustCompile("\\$\\{([^\\}]+)\\}")
	out := re.ReplaceAllStringFunc(str, func(in string) string {

		in = strings.TrimPrefix(in, "${")
		in = strings.TrimSuffix(in, "}")

		tmp := &object.String{Value: in}
		key := tmp.HashKey()
		k, ok := hashObject.Pairs[key]
		if ok {
			return (k.Value.Inspect())
		} else {
			return "${" + in + "}"
		}
	})
	return &object.String{Value: out}
}

func stringToUpper(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument must be a string, got=%s",
			args[0].Type())
	}
	input := args[0].(*object.String).Value
	return &object.String{Value: strings.ToUpper(input)}
}

func stringToLower(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument must be a string, got=%s",
			args[0].Type())
	}
	input := args[0].(*object.String).Value
	return &object.String{Value: strings.ToLower(input)}
}

func stringTrim(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument must be a string, got=%s",
			args[0].Type())
	}
	input := args[0].(*object.String).Value
	return &object.String{Value: strings.TrimSpace(input)}
}

func stringSplit(args ...object.Object) object.Object {
	//
	// Default separator.
	//
	sep := " "

	if len(args) != 1 && len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=1 or 2.",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument to `split` must be a string, got=%s",
			args[0].Type())
	}
	if len(args) == 2 {
		if args[1].Type() != object.STRING_OBJ {
			return newError("argument to `split` must be a string, got=%s",
				args[0].Type())
		}
		sep = args[1].(*object.String).Value

	}

	// split by separator
	fields := strings.Split(args[0].(*object.String).Value, sep)

	// make results
	l := len(fields)
	result := make([]object.Object, l, l)
	for i, txt := range fields {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

func init() {

	registerBuiltin("string.interpolate",
		func(args ...object.Object) object.Object {
			return (stringInterpolate(args...))
		})
	registerBuiltin("string.toupper",
		func(args ...object.Object) object.Object {
			return (stringToUpper(args...))
		})
	registerBuiltin("string.tolower",
		func(args ...object.Object) object.Object {
			return (stringToLower(args...))
		})
	registerBuiltin("string.trim",
		func(args ...object.Object) object.Object {
			return (stringTrim(args...))
		})
	registerBuiltin("string.split",
		func(args ...object.Object) object.Object {
			return (stringSplit(args...))
		})
}
