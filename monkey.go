//
// Driver for monkey.
//
// If no argument is given will read from stdin, otherwise from the
// named file.
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/skx/monkey/evaluator"
	"github.com/skx/monkey/lexer"
	"github.com/skx/monkey/object"
	"github.com/skx/monkey/parser"
)

// This version-string will be updated via travis for generated binaries.
var version = "master/unreleased"

//
// Implemention of "version()" function.
//
func versionFun(args ...object.Object) object.Object {
	return &object.String{Value: version}
}

//
// Implemention of "args()" function.
//
func argsFun(args ...object.Object) object.Object {
	l := len(os.Args[1:])
	result := make([]object.Object, l, l)
	for i, txt := range os.Args[1:] {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

//
// Execute the supplied string as a program.
//
func Execute(input string) int {

	env := object.NewEnvironment()
	env.RegisterDefaults()
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

	//
	// Our standard-library is mostly written in C,
	// and can be found implemented in evaluator/stdlib*.go
	//
	// However there is also data/stdlib.mon, and here we
	// load that
	//
	tmpl, err := getResource("data/stdlib.mon")
	if err != nil {
		fmt.Printf("Failed to load our standard-library: %s",
			err.Error())
		os.Exit(33)
	}

	//
	//  Parse and evaluate our standard-library.
	//
	init_l := lexer.New(string(tmpl))
	init_p := parser.New(init_l)
	init_prog := init_p.ParseProgram()
	evaluator.Eval(init_prog, env)

	//
	//  Now evaluate the code the user wanted to load.
	//
	//  Note that here our environment will still contain
	// the code we just loaded from our data-resource
	//
	//  (i.e. Our monkey-based standard library.)
	//
	evaluator.Eval(program, env)
	return 0
}

func main() {

	var input []byte
	var err error

	if len(os.Args) > 1 {
		input, err = ioutil.ReadFile(os.Args[1])
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
	}

	Execute(string(input))
}
