//
// Driver for monkey.
//
// If no argument is given will read from stdin, otherwise from the
// named file.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

var VERSION = "0.2"

// Implemention of "version()" function.
func versionFun(args ...object.Object) object.Object {
	return &object.String{Value: VERSION}
}

// Implemention of "args()" function.
func argsFun(args ...object.Object) object.Object {
	l := len(os.Args[1:])
	result := make([]object.Object, l, l)
	for i, txt := range os.Args[1:] {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

// Execute the program in the given file.
func Execute(input string) int {

	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}

	// Register a function called version()
	// that the script can call.
	evaluator.RegisterBuiltin("version",
		func(args ...object.Object) object.Object {
			return (versionFun(args...))
		})

	// Access to the command-line arguments
	evaluator.RegisterBuiltin("args",
		func(args ...object.Object) object.Object {
			return (argsFun(args...))
		})

	evaluator.Eval(program, env)
	return 0
}

func main() {

	var input []byte
	var err error

	if len(os.Args) > 1 {
		input, err = ioutil.ReadFile(flag.Arg(1))
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
	}

	Execute(string(input))
}
