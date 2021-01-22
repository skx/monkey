package main

import (
	"fmt"
	"os"

	"github.com/skx/monkey/evaluator"
	"github.com/skx/monkey/lexer"
	"github.com/skx/monkey/object"
	"github.com/skx/monkey/parser"
)

func main() {
	env := object.NewEnvironment()
	l := lexer.New(`{"b":fn(){return {"x":2};}}.b().x`)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(1)
	}


	//
	//  Now evaluate the code the user wanted to load.
	//
	//  Note that here our environment will still contain
	// the code we just loaded from our data-resource
	//
	//  (i.e. Our monkey-based standard library.)
	//
	fmt.Println(evaluator.Eval(program, env))
}
