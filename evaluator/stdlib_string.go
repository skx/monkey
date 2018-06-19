package evaluator

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/skx/monkey/object"
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

func stringReverse(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("argument must be a string, got=%s",
			args[0].Type())
	}
	input := args[0].(*object.String).Value

	out := make([]rune, utf8.RuneCountInString(input))
	i := len(out)
	for _, c := range input {
		i--
		out[i] = c
	}
	return &object.String{Value: string(out)}
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

	RegisterBuiltin("string.interpolate",
		func(args ...object.Object) object.Object {
			return (stringInterpolate(args...))
		})
	RegisterBuiltin("string.toupper",
		func(args ...object.Object) object.Object {
			return (stringToUpper(args...))
		})
	RegisterBuiltin("string.tolower",
		func(args ...object.Object) object.Object {
			return (stringToLower(args...))
		})
	RegisterBuiltin("string.trim",
		func(args ...object.Object) object.Object {
			return (stringTrim(args...))
		})
	RegisterBuiltin("string.reverse",
		func(args ...object.Object) object.Object {
			return (stringReverse(args...))
		})
	RegisterBuiltin("string.split",
		func(args ...object.Object) object.Object {
			return (stringSplit(args...))
		})
}
