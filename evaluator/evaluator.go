// Package evaluator contains the core of our interpreter, which walks
// the AST produced by the parser and evaluates the user-submitted program.
package evaluator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/skx/monkey/ast"
	"github.com/skx/monkey/object"
)

// pre-defined object including Null, True and False
var (
	NULL    = &object.Null{}
	TRUE    = &object.Boolean{Value: true}
	FALSE   = &object.Boolean{Value: false}
	PRAGMAS = make(map[string]int)
)

// The built-in functions / standard-library methods are stored here.
var builtins = map[string]*object.Builtin{}

// Eval is our core function for evaluating nodes.
func Eval(node ast.Node, env *object.Environment) object.Object {
	return EvalContext(context.Background(), node, env)
}

// EvalContext is our core function for evaluating nodes.
// The context.Context provided can be used to cancel a running script instance.
func EvalContext(ctx context.Context, node ast.Node, env *object.Environment) object.Object {

	//
	// We test our context at every iteration of our main-loop.
	//
	select {
	case <-ctx.Done():
		return &object.Error{Message: ctx.Err().Error()}
	default:
		// nop
	}

	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalProgram(ctx, node, env)
	case *ast.ExpressionStatement:
		return EvalContext(ctx, node.Expression, env)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.NullLiteral:
		return NULL
	case *ast.PrefixExpression:
		right := EvalContext(ctx, node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.PostfixExpression:
		return evalPostfixExpression(env, node.Operator, node)
	case *ast.InfixExpression:
		left := EvalContext(ctx, node.Left, env)
		if isError(left) {
			return left
		}
		right := EvalContext(ctx, node.Right, env)
		if isError(right) {
			return right
		}
		res := evalInfixExpression(node.Operator, left, right, env)
		if isError(res) {
			fmt.Printf("Error: %s\n", res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
		}
		return (res)

	case *ast.BlockStatement:
		return evalBlockStatement(ctx, node, env)
	case *ast.IfExpression:
		return evalIfExpression(ctx, node, env)
	case *ast.TernaryExpression:
		return evalTernaryExpression(ctx, node, env)
	case *ast.ForLoopExpression:
		return evalForLoopExpression(ctx, node, env)
	case *ast.ForeachStatement:
		return evalForeachExpression(ctx, node, env)
	case *ast.ReturnStatement:
		val := EvalContext(ctx, node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		val := EvalContext(ctx, node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return val
	case *ast.ConstStatement:
		val := EvalContext(ctx, node.Value, env)
		if isError(val) {
			return val
		}
		env.SetConst(node.Name.Value, val)
		return val
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		defaults := node.Defaults
		return &object.Function{Parameters: params, Env: env, Body: body, Defaults: defaults}
	case *ast.FunctionDefineLiteral:
		params := node.Parameters
		body := node.Body
		defaults := node.Defaults
		env.Set(node.TokenLiteral(), &object.Function{Parameters: params, Env: env, Body: body, Defaults: defaults})
		return NULL
	case *ast.ObjectCallExpression:
		res := evalObjectCallExpression(ctx, node, env)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling object-method %s\n", res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
		}
		return res
	case *ast.CallExpression:
		function := EvalContext(ctx, node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpression(ctx, node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		res := applyFunction(ctx, env, function, args)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling `%s` : %s\n", node.Function, res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
			return res
		}
		return res

	case *ast.ArrayLiteral:
		elements := evalExpression(ctx, node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.RegexpLiteral:
		return &object.Regexp{Value: node.Value, Flags: node.Flags}
	case *ast.BacktickLiteral:
		return backTickOperation(node.Value)
	case *ast.IndexExpression:
		left := EvalContext(ctx, node.Left, env)
		if isError(left) {
			return left
		}
		index := EvalContext(ctx, node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *ast.AssignStatement:
		return evalAssignStatement(ctx, node, env)
	case *ast.HashLiteral:
		return evalHashLiteral(ctx, node, env)
	case *ast.SwitchExpression:
		return evalSwitchStatement(ctx, node, env)
	}
	return nil
}

// eval block statement
func evalBlockStatement(ctx context.Context, block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = EvalContext(ctx, statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

// for performance, using single instance of boolean
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// eval prefix expression
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalPostfixExpression(env *object.Environment, operator string, node *ast.PostfixExpression) object.Object {
	switch operator {
	case "++":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return newError("%s is unknown", node.Token.Literal)
		}

		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value
			env.Set(node.Token.Literal, &object.Integer{Value: v + 1})
			return arg
		default:
			return newError("%s is not an int", node.Token.Literal)

		}
	case "--":
		val, ok := env.Get(node.Token.Literal)
		if !ok {
			return newError("%s is unknown", node.Token.Literal)
		}

		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value
			env.Set(node.Token.Literal, &object.Integer{Value: v - 1})
			return arg
		default:
			return newError("%s is not an int", node.Token.Literal)
		}
	default:
		return newError("unknown operator: %s", operator)
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	// Found by fuzzing
	if right == nil {
		return newError("null operand %v", right)
	}

	switch obj := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -obj.Value}
	case *object.Float:
		return &object.Float{Value: -obj.Value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object, env *object.Environment) object.Object {

	// Found by fuzzing
	if left == nil || right == nil {
		return newError("null operand %v %v", left, right)
	}

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalFloatIntegerInfixExpression(operator, left, right)
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalIntegerFloatInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "&&":
		return nativeBoolToBooleanObject(objectToNativeBoolean(left) && objectToNativeBoolean(right))
	case operator == "||":
		return nativeBoolToBooleanObject(objectToNativeBoolean(left) || objectToNativeBoolean(right))
	case operator == "!~":
		return notMatches(left, right)
	case operator == "~=":
		return matches(left, right, env)

	case operator == "==":
		return nativeBoolToBooleanObject(left == right)

	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func matches(left, right object.Object, env *object.Environment) object.Object {

	str := left.Inspect()

	if right.Type() != object.REGEXP_OBJ {
		return newError("regexp required for regexp-match, given %s", right.Type())
	}

	val := right.(*object.Regexp).Value
	if right.(*object.Regexp).Flags != "" {
		val = "(?" + right.(*object.Regexp).Flags + ")" + val
	}

	// Compile the regular expression.
	r, err := regexp.Compile(val)

	// Ensure it compiled
	if err != nil {
		return newError("error compiling regexp '%s': %s", right.Inspect(), err)
	}

	res := r.FindStringSubmatch(str)

	// Do we have any captures?
	if len(res) > 1 {
		for i := 1; i < len(res); i++ {
			env.Set(fmt.Sprintf("$%d", i), &object.String{Value: res[i]})
		}
	}

	// Test if it matched
	if len(res) > 0 {
		return TRUE
	}

	return FALSE
}

func notMatches(left, right object.Object) object.Object {
	str := left.Inspect()

	if right.Type() != object.REGEXP_OBJ {
		return newError("regexp required for regexp-match, given %s", right.Type())
	}

	val := right.(*object.Regexp).Value
	if right.(*object.Regexp).Flags != "" {
		val = "(?" + right.(*object.Regexp).Flags + ")" + val
	}

	// Compile the regular expression.
	r, err := regexp.Compile(val)

	// Ensure it compiled
	if err != nil {
		return newError("error compiling regexp '%s': %s", right.Inspect(), err)
	}

	// Test if it matched
	if r.MatchString(str) {
		return FALSE
	}

	return TRUE
}

// boolean operations
func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	// convert the bools to strings.
	l := &object.String{Value: string(left.Inspect())}
	r := &object.String{Value: string(right.Inspect())}

	switch operator {
	case "<":
		return evalStringInfixExpression(operator, l, r)
	case "<=":
		return evalStringInfixExpression(operator, l, r)
	case ">":
		return evalStringInfixExpression(operator, l, r)
	case ">=":
		return evalStringInfixExpression(operator, l, r)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	// Found by fuzzing
	if left == nil || right == nil {
		return newError("null operand %v %v", left, right)
	}

	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "+=":
		return &object.Integer{Value: leftVal + rightVal}
	case "%":
		// Found by fuzzing
		if rightVal == 0 {
			return newError("divide by zero")
		}

		return &object.Integer{Value: leftVal % rightVal}
	case "**":
		return &object.Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "-=":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "*=":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		// Found by fuzzing
		if rightVal == 0 {
			return newError("divide by zero")
		}
		return &object.Integer{Value: leftVal / rightVal}
	case "/=":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "..":
		// The start and end might not be ascending, so the size
		// will be the span
		diff := float64(rightVal - leftVal)
		len := int(math.Abs(diff)) + 1

		// Step is generally +1, but if we're going to
		// express the range "10..0" it will be -1 to allow
		// us to count down via subtraction
		var step int64
		step = 1.0

		if rightVal < leftVal {
			step = -1.0
		}

		// Found by fuzzing
		if len > 2048 {
			return newError("impossible large range for .. operator")
		}

		// Make an array to hold the return value
		array := make([]object.Object, len)

		// Now make the range of integers, counting via the step.
		i := 0
		for i < len {
			array[i] = &object.Integer{Value: leftVal}
			leftVal += step
			i++
		}
		return &object.Array{Elements: array}
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "+=":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "-=":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "*=":
		return &object.Float{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		// Found by fuzzing
		if rightVal == 0 {
			return newError("divide by zero")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "/=":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalFloatIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := float64(right.(*object.Integer).Value)
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "+=":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "-=":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "*=":
		return &object.Float{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		// Found by fuzzing
		if rightVal == 0 {
			return newError("divide by zero")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "/=":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalIntegerFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := float64(left.(*object.Integer).Value)
	rightVal := right.(*object.Float).Value
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "+=":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "-=":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "*=":
		return &object.Float{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		// Found by fuzzing
		if rightVal == 0 {
			return newError("divide by zero")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "/=":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	l := left.(*object.String)
	r := right.(*object.String)

	switch operator {
	case "==":
		return nativeBoolToBooleanObject(l.Value == r.Value)
	case "!=":
		return nativeBoolToBooleanObject(l.Value != r.Value)
	case ">=":
		return nativeBoolToBooleanObject(l.Value >= r.Value)
	case ">":
		return nativeBoolToBooleanObject(l.Value > r.Value)
	case "<=":
		return nativeBoolToBooleanObject(l.Value <= r.Value)
	case "<":
		return nativeBoolToBooleanObject(l.Value < r.Value)
	case "+":
		return &object.String{Value: l.Value + r.Value}
	case "+=":
		return &object.String{Value: l.Value + r.Value}
	}

	return newError("unknown operator: %s %s %s",
		left.Type(), operator, right.Type())
}

// evalIfExpression handles an `if` expression, running the block
// if the condition matches, and running any optional else block
// otherwise.
func evalIfExpression(ctx context.Context, ie *ast.IfExpression, env *object.Environment) object.Object {
	//
	// Create an environment for handling regexps
	//
	var permit []string
	i := 1
	for i < 32 {
		permit = append(permit, fmt.Sprintf("$%d", i))
		i++
	}
	nEnv := object.NewTemporaryScope(env, permit)
	condition := EvalContext(ctx, ie.Condition, nEnv)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return EvalContext(ctx, ie.Consequence, nEnv)
	} else if ie.Alternative != nil {
		return EvalContext(ctx, ie.Alternative, nEnv)
	} else {
		return NULL
	}
}

// evalTernaryExpression handles a ternary-expression.  If the condition
// is true we return the contents of evaluating the true-branch, otherwise
// the false-branch.  (Unlike an `if` statement we know that we always have
// an alternative/false branch.)
func evalTernaryExpression(ctx context.Context, te *ast.TernaryExpression, env *object.Environment) object.Object {

	condition := EvalContext(ctx, te.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return EvalContext(ctx, te.IfTrue, env)
	}
	return EvalContext(ctx, te.IfFalse, env)
}

func evalAssignStatement(ctx context.Context, a *ast.AssignStatement, env *object.Environment) (val object.Object) {
	evaluated := EvalContext(ctx, a.Value, env)
	if isError(evaluated) {
		return evaluated
	}

	//
	// An assignment is generally:
	//
	//    variable = value
	//
	// But we cheat and reuse the implementation for:
	//
	//    i += 4
	//
	// In this case we record the "operator" as "+="
	//
	switch a.Operator {
	case "+=":
		// Get the current value
		current, ok := env.Get(a.Name.String())
		if !ok {
			return newError("%s is unknown", a.Name.String())
		}

		res := evalInfixExpression("+=", current, evaluated, env)
		if isError(res) {
			return res
		}

		env.Set(a.Name.String(), res)
		return res

	case "-=":

		// Get the current value
		current, ok := env.Get(a.Name.String())
		if !ok {
			return newError("%s is unknown", a.Name.String())
		}

		res := evalInfixExpression("-=", current, evaluated, env)
		if isError(res) {
			return res
		}

		env.Set(a.Name.String(), res)
		return res

	case "*=":
		// Get the current value
		current, ok := env.Get(a.Name.String())
		if !ok {
			return newError("%s is unknown", a.Name.String())
		}

		res := evalInfixExpression("*=", current, evaluated, env)
		if isError(res) {
			return res
		}

		env.Set(a.Name.String(), res)
		return res

	case "/=":

		// Get the current value
		current, ok := env.Get(a.Name.String())
		if !ok {
			return newError("%s is unknown", a.Name.String())
		}

		res := evalInfixExpression("/=", current, evaluated, env)
		if isError(res) {
			return res
		}

		env.Set(a.Name.String(), res)
		return res

	case "=":
		// If we're running with the strict-pragma it is
		// a bug to set a variable which wasn't declared (via let).
		if PRAGMAS["strict"] == 1 {
			_, ok := env.Get(a.Name.String())
			if !ok {
				fmt.Printf("Setting unknown variable '%s' is a bug under strict-pragma!\n", a.Name.String())
				os.Exit(1)
			}
		}

		env.Set(a.Name.String(), evaluated)
	}
	return evaluated
}

func evalSwitchStatement(ctx context.Context, se *ast.SwitchExpression, env *object.Environment) object.Object {

	// Get the value.
	obj := EvalContext(ctx, se.Value, env)

	// Try all the choices
	for _, opt := range se.Choices {

		// skipping the default-case, which we'll
		// handle later.
		if opt.Default {
			continue
		}

		// Look at any expression we've got in this case.
		for _, val := range opt.Expr {

			// Get the value of the case
			out := EvalContext(ctx, val, env)

			// Is it a literal match?
			if obj.Type() == out.Type() &&
				(obj.Inspect() == out.Inspect()) {

				// Evaluate the block and return the value
				blockOut := evalBlockStatement(ctx, opt.Block, env)
				return blockOut
			}

			// Is it a regexp-match?
			if out.Type() == object.REGEXP_OBJ {

				m := matches(obj, out, env)
				if m == TRUE {

					// Evaluate the block and return the value
					out := evalBlockStatement(ctx, opt.Block, env)
					return out

				}
			}
		}
	}

	// No match?  Handle default if present
	for _, opt := range se.Choices {

		// skip default
		if opt.Default {

			out := evalBlockStatement(ctx, opt.Block, env)
			return out
		}
	}

	return nil
}

func evalForLoopExpression(ctx context.Context, fle *ast.ForLoopExpression, env *object.Environment) object.Object {
	rt := &object.Boolean{Value: true}
	for {
		condition := EvalContext(ctx, fle.Condition, env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			rt := EvalContext(ctx, fle.Consequence, env)
			if !isError(rt) && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
				return rt
			}
		} else {
			break
		}
	}
	return rt
}

// handle "for x [,y] in .."
func evalForeachExpression(ctx context.Context, fle *ast.ForeachStatement, env *object.Environment) object.Object {

	// expression
	val := EvalContext(ctx, fle.Value, env)

	helper, ok := val.(object.Iterable)
	if !ok {
		return newError("%s object doesn't implement the Iterable interface", val.Type())
	}

	// The one/two values we're going to permit
	var permit []string
	permit = append(permit, fle.Ident)
	if fle.Index != "" {
		permit = append(permit, fle.Index)
	}

	// Create a new environment for the block
	//
	// This will allow writing EVERYTHING to the parent scope,
	// except the two variables named in the permit-array
	child := object.NewTemporaryScope(env, permit)

	// Reset the state of any previous iteration.
	helper.Reset()

	// Get the initial values.
	ret, idx, ok := helper.Next()

	for ok {

		// Set the index + name
		child.Set(fle.Ident, ret)

		idxName := fle.Index
		if idxName != "" {
			child.Set(fle.Index, idx)
		}

		// Eval the block
		rt := EvalContext(ctx, fle.Body, child)

		//
		// If we got an error/return then we handle it.
		//
		if !isError(rt) && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
			return rt
		}

		// Loop again
		ret, idx, ok = helper.Next()
	}

	return &object.Null{}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalProgram(ctx context.Context, program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = EvalContext(ctx, statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	fmt.Fprintf(os.Stderr, "identifier not found: %s\n", node.Value)
	if PRAGMAS["strict"] == 1 {
		os.Exit(1)
	}
	return newError("identifier not found: " + node.Value)
}

func evalExpression(ctx context.Context, exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := EvalContext(ctx, e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

// parseCommandLine takes a command string and splits it into individual arguments,
// respecting quotes and escaping within the command.
// So this input:
//
//	/bin/sh -c "ls /etc"
//
// Would give output of the form:
//
//	/bin/sh
//	-c
//	ls /etc
func parseCommandLine(command string) ([]string, error) {
	var args []string
	var current strings.Builder
	inQuotes := false
	var quoteChar rune

	// flush appends the current argument to the args slice and resets the current string builder.
	flush := func() {
		if current.Len() > 0 {
			args = append(args, current.String())
			current.Reset()
		}
	}

	// Iterate through each character in the command string.
	for _, c := range command {
		switch {
		case unicode.IsSpace(c) && !inQuotes:
			// If a space is encountered outside of quotes, flush the current argument.
			flush()
		case (c == '\'' || c == '"') && !inQuotes:
			// If a single or double quote is encountered and we're not inside quotes,
			// mark the start of quoted text and record the quote character.
			inQuotes = true
			quoteChar = c
		case c == quoteChar && inQuotes:
			// If the matching closing quote is found while inside quotes, mark the end of quoted text
			// and flush the current argument.
			inQuotes = false
			flush()
		default:
			// Otherwise, append the character to the current argument.
			current.WriteRune(c)
		}
	}

	// If still inside quotes at the end of parsing, return an error for unclosed quotes.
	if inQuotes {
		return nil, fmt.Errorf("unclosed quote in command line: %s", command)
	}

	// Flush any remaining argument and return the parsed arguments.
	flush()
	return args, nil
}

// backTickOperation executes a shell command and returns a hash object containing the result.
// The hash includes 'stdout', 'stderr', and 'code' fields.
// If the command is empty or parsing fails, an error hash is returned.
func backTickOperation(command string) object.Object {
	var (
		args []string
		err  error
	)

	// Trim leading and trailing whitespace from the command.
	if command = strings.TrimSpace(command); command != "" {
		// Split the command into arguments.
		if args, err = parseCommandLine(command); err != nil {
			// Return an error hash for parsing failure.
			return createCommandExecHash(&object.String{Value: ""}, &object.String{Value: "parse error: " + err.Error()},
				&object.Integer{Value: -1})
		}
	}

	// Check if the command is empty after parsing.
	if len(args) == 0 {
		// Return an error hash for an empty command.
		return createCommandExecHash(&object.String{Value: ""}, &object.String{Value: "no command"},
			&object.Integer{Value: -1})
	}

	// Run the command.
	cmd := exec.Command(filepath.Clean(args[0]), args[1:]...)

	// Capture the command's stdout and stderr.
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	var exitCode int64 = 0

	// Execute the command and handle errors.
	err = cmd.Run()
	if err != nil {
		var exitError *exec.ExitError
		if !errors.As(err, &exitError) {
			// Handle non-ExitError errors (e.g., command not found).
			return createCommandExecHash(&object.String{Value: ""}, &object.String{Value: fmt.Sprintf("Failed to run '%s' -> %s\n", command, err.Error())},
				&object.Integer{Value: -1})
		}
		exitCode = int64(exitError.ExitCode())
	}

	// Create a hash with 'stdout', 'stderr', and 'code' fields.
	return createCommandExecHash(&object.String{Value: stdout.String()}, &object.String{Value: stderr.String()},
		&object.Integer{Value: exitCode})
}

// createCommandExecHash Create a hash with 'stdout', 'stderr', and 'code' fields.
func createCommandExecHash(stdoutObj, stderrObj, errorObj object.Object) object.Object {
	// Create keys for the hash.
	stdoutKey := &object.String{Value: "stdout"}
	stderrKey := &object.String{Value: "stderr"}
	exitCodeKey := &object.String{Value: "exitCode"}

	// Populate the hash with key-value pairs.
	hashPairs := map[object.HashKey]object.HashPair{
		stdoutKey.HashKey():   {Key: stdoutKey, Value: stdoutObj},
		stderrKey.HashKey():   {Key: stderrKey, Value: stderrObj},
		exitCodeKey.HashKey(): {Key: exitCodeKey, Value: errorObj},
	}

	// Create and return the hash object.
	return &object.Hash{Pairs: hashPairs}
}

func evalIndexExpression(left, index object.Object) object.Object {

	// Found by fuzzing
	if left == nil || index == nil {
		return newError("null operand %v[%v]", left, index)
	}

	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	case left.Type() == object.STRING_OBJ:
		return evalStringIndexExpression(left, index)
	default:
		return newError("index operator not support:%s", left.Type())

	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > max {
		return NULL
	}
	return arrayObject.Elements[idx]
}
func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}

func evalStringIndexExpression(input, index object.Object) object.Object {
	str := input.(*object.String).Value
	idx := index.(*object.Integer).Value
	max := int64(len(str))
	if idx < 0 || idx > max {
		return NULL
	}

	// Get the characters as an array of runes
	chars := []rune(str)

	// Now index
	ret := chars[idx]

	// And return as a string.
	return &object.String{Value: string(ret)}
}

func evalHashLiteral(ctx context.Context, node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := EvalContext(ctx, keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}
		value := EvalContext(ctx, valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}

	}
	return &object.Hash{Pairs: pairs}

}

func applyFunction(ctx context.Context, env *object.Environment, fn object.Object, args []object.Object) object.Object {

	// Found by fuzzing
	if fn == nil {
		return newError("impossible empty body on function-call")
	}
	switch fn := fn.(type) {
	case *object.Function:
		extendEnv := extendFunctionEnv(ctx, fn, args)
		evaluated := EvalContext(ctx, fn.Body, extendEnv)
		return upwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(env, args...)
	default:
		return newError("not a function: %s", fn.Type())
	}

}

func extendFunctionEnv(ctx context.Context, fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	// Set the defaults
	for key, val := range fn.Defaults {
		env.Set(key, EvalContext(ctx, val, env))
	}
	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		}
	}
	return env
}

func upwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

// RegisterBuiltin registers a built-in function.  This is used to register
// our "standard library" functions.
func RegisterBuiltin(name string, fun object.BuiltinFunction) {
	builtins[name] = &object.Builtin{Fn: fun}
}

// evalObjectCallExpression invokes methods against objects.
func evalObjectCallExpression(ctx context.Context, call *ast.ObjectCallExpression, env *object.Environment) object.Object {

	obj := EvalContext(ctx, call.Object, env)

	if obj == nil {
		return newError("impossible object-call on an empty object")
	}

	if method, ok := call.Call.(*ast.CallExpression); ok {

		//
		// Here we try to invoke the object.method() call which has
		// been implemented in go.
		//
		// We do this by forwarding the call to the appropriate
		// `invokeMethod` interface on the object.
		//
		args := evalExpression(ctx, call.Call.(*ast.CallExpression).Arguments, env)
		ret := obj.InvokeMethod(method.Function.String(), *env, args...)
		if ret != nil {
			return ret
		}

		//
		// If we reach this point then the invokation didn't
		// succeed, that probably means that the function wasn't
		// implemented in go.
		//
		// So now we want to look for it in monkey, and we have
		// enough details to find the appropriate function.
		//
		//  * We have the object involved.
		//
		//  * We have the type of that object.
		//
		//  * We have the name of the function.
		//
		//  * We have the arguments.
		//
		// We'll use the type + name to lookup the (global) function
		// to invoke.  For example in this case we'll invoke
		// `string.len()` - because the type of the object we're
		// invoking-against is string:
		//
		//  "steve".len();
		//
		// For this case we'll be looking for `array.foo()`.
		//
		//   let a = [ 1, 2, 3 ];
		//   puts( a.foo() );
		//
		// As a final fall-back we'll look for "object.foo()"
		// if "array.foo()" isn't defined.
		//
		//
		//
		attempts := []string{}
		attempts = append(attempts, strings.ToLower(string(obj.Type())))
		attempts = append(attempts, "object")

		//
		// Look for "$type.name", or "object.name"
		//
		for _, prefix := range attempts {

			//
			// What we're attempting to execute.
			//
			name := prefix + "." + method.Function.String()

			//
			// Try to find that function in our environment.
			//
			if fn, ok := env.Get(name); ok {

				//
				// Extend our environment with the functional-args.
				//
				extendEnv := extendFunctionEnv(ctx, fn.(*object.Function), args)

				//
				// Now set "self" to be the implicit object, against
				// which the function-call will be operating.
				//
				extendEnv.Set("self", obj)

				//
				// Finally invoke & return.
				//
				evaluated := EvalContext(ctx, fn.(*object.Function).Body, extendEnv)
				obj = upwrapReturnValue(evaluated)
				return obj
			}
		}

	}

	//
	// If we hit this point we have had a method invoked which
	// was neither defined in go nor monkey.
	//
	// e.g. "steve".md5sum()
	//
	// So we've got no choice but to return an error.
	//
	return newError("Failed to invoke method: %s", call.Call.(*ast.CallExpression).Function.String())
}

func objectToNativeBoolean(o object.Object) bool {
	if r, ok := o.(*object.ReturnValue); ok {
		o = r.Value
	}
	switch obj := o.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.String:
		return obj.Value != ""
	case *object.Regexp:
		return obj.Value != ""
	case *object.Null:
		return false
	case *object.Integer:
		if obj.Value == 0 {
			return false
		}
		return true
	case *object.Float:
		if obj.Value == 0.0 {
			return false
		}
		return true
	case *object.Array:
		if len(obj.Elements) == 0 {
			return false
		}
		return true
	case *object.Hash:
		if len(obj.Pairs) == 0 {
			return false
		}
		return true
	default:
		return true
	}
}
