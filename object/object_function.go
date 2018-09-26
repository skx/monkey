package object

import (
	"bytes"
	"strings"

	"github.com/skx/monkey/ast"
)

// Function wraps ast.Identifier array, ast.BlockStatement and Environment and implements Object interface.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Defaults   map[string]ast.Expression
	Env        *Environment
	Const      bool
}

// SetConstant allows an object to be marked as read-only, or constant.
func (f *Function) SetConstant(val bool) {
	f.Const = val
}

// Constant returns true if an object is read-only or constant.
func (f *Function) Constant() bool {
	return f.Const
}

// Type returns the type of this object.
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect returns a string-representation of the given object.
func (f *Function) Inspect() string {
	var out bytes.Buffer
	parameters := make([]string, 0)
	for _, p := range f.Parameters {
		parameters = append(parameters, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(parameters, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (f *Function) InvokeMethod(method string, args ...Object) Object {
	return nil
}
