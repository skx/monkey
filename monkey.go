package main

import (
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

func Execute(filename string) int {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return 1
	}

	env := object.NewEnvironment()
	l := lexer.New(string(body))
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

	evaluator.Eval(program, env)
	return 0
}

func main() {
	for _, file := range os.Args[1:] {
		Execute(file)
	}
}
