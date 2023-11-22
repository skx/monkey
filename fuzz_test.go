//go:build go1.18
// +build go1.18

package main

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/skx/monkey/evaluator"
	"github.com/skx/monkey/lexer"
	"github.com/skx/monkey/object"
	"github.com/skx/monkey/parser"
)

// FuzzMonkey runs the fuzz-testing against our parser and interpreter.
func FuzzMonkey(f *testing.F) {

	// Known errors we might see
	known := []string{
		"as integer",
		"divide by zero",
		"null operand",
		"could not parse",
		"exceeded",
		"expected assign",
		"expected next token",
		"impossible",
		"nested ternary expressions are illegal",
		"no prefix parse function",
	}

	f.Fuzz(func(t *testing.T, input []byte) {

		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		env := object.NewEnvironment()
		l := lexer.New(string(input))
		p := parser.New(l)

		program := p.ParseProgram()
		falsePositive := false

		// No errors?  Then execute
		if len(p.Errors()) == 0 {

			evaluator.EvalContext(ctx, program, env)
			return
		}

		for _, msg := range p.Errors() {
			for _, ignored := range known {
				if strings.Contains(msg, ignored) {
					falsePositive = true
				}
			}

		}

		if !falsePositive {
			t.Fatalf("error running input: '%s': %v", input, p.Errors())
		}
	})
}
