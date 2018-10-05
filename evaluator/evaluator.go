package evaluator

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
	switch node := node.(type) {

	//Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	//Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.PostfixExpression:
		return evalPostfixExpression(env, node.Operator, node)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		res := evalInfixExpression(node.Operator, left, right)
		if isError(res) {
			fmt.Printf("Error: %s\n", res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
		}
		return (res)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ForLoopExpression:
		return evalForLoopExpression(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return val
	case *ast.ConstStatement:
		val := Eval(node.Value, env)
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
		res := evalObjectCallExpression(node, env)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling object-method %s\n", res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
		}
		return res
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpression(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		res := applyFunction(env, function, args)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling `%s` : %s\n", node.Function, res.Inspect())
			if PRAGMAS["strict"] == 1 {
				os.Exit(1)
			}
			return res
		}
		return res

	case *ast.ArrayLiteral:
		elements := evalExpression(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.BacktickLiteral:
		return backTickOperation(node.Value)
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *ast.AssignStatement:
		return evalAssignStatement(node, env)
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	}
	return nil
}

// eval block statement
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
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
	switch obj := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -obj.Value}
	case *object.Float:
		return &object.Float{Value: -obj.Value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
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
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "+=":
		return &object.Integer{Value: leftVal + rightVal}
	case "%":
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

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalAssignStatement(a *ast.AssignStatement, env *object.Environment) (val object.Object) {
	evaluated := Eval(a.Value, env)
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

		res := evalInfixExpression("+=", current, evaluated)
		if isError(res) {
			fmt.Printf("Error handling += %s\n", res.Inspect())
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

		res := evalInfixExpression("-=", current, evaluated)
		if isError(res) {
			fmt.Printf("Error handling -= %s\n", res.Inspect())
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

		res := evalInfixExpression("*=", current, evaluated)
		if isError(res) {
			fmt.Printf("Error handling *= %s\n", res.Inspect())
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

		res := evalInfixExpression("/=", current, evaluated)
		if isError(res) {
			fmt.Printf("Error handling /= %s\n", res.Inspect())
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

func evalForLoopExpression(fle *ast.ForLoopExpression, env *object.Environment) object.Object {
	rt := &object.Boolean{Value: true}
	for {
		condition := Eval(fle.Condition, env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			rt := Eval(fle.Consequence, env)
			if !isError(rt) && (rt.Type() == object.RETURN_VALUE_OBJ || rt.Type() == object.ERROR_OBJ) {
				return rt
			}
		} else {
			break
		}
	}
	return rt
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

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
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

func evalExpression(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

// Split a line of text into tokens, but keep anything "quoted"
// together..
//
// So this input:
//
//   /bin/sh -c "ls /etc"
//
// Would give output of the form:
//   /bin/sh
//   -c
//   ls /etc
//
func splitCommand(input string) []string {

	//
	// This does the split into an array
	//
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	res := r.FindAllString(input, -1)

	//
	// However the resulting pieces might be quoted.
	// So we have to remove them, if present.
	//
	var result []string
	for _, e := range res {
		result = append(result, trimQuotes(e, '"'))
	}
	return (result)
}

// Remove balanced characters around a string.
func trimQuotes(in string, c byte) string {
	if len(in) >= 2 {
		if in[0] == c && in[len(in)-1] == c {
			return in[1 : len(in)-1]
		}
	}
	return in
}

// Run a command and return a hash containing the result.
// `stderr`, `stdout`, and `error` will be the fields
func backTickOperation(command string) object.Object {

	// split the command
	toExec := splitCommand(command)
	cmd := exec.Command(toExec[0], toExec[1:]...)

	// get the result
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()

	// If the command exits with a non-zero exit-code it
	// is regarded as a failure.  Here we test for ExitError
	// to regard that as a non-failure.
	if err != nil && err != err.(*exec.ExitError) {
		fmt.Printf("Failed to run '%s' -> %s\n", command, err.Error())
		return NULL
	}

	//
	// The result-objects to store in our hash.
	//
	stdout := &object.String{Value: outb.String()}
	stderr := &object.String{Value: errb.String()}

	// Create keys
	stdoutKey := &object.String{Value: "stdout"}
	stdoutHash := object.HashPair{Key: stdoutKey, Value: stdout}

	stderrKey := &object.String{Value: "stderr"}
	stderrHash := object.HashPair{Key: stderrKey, Value: stderr}

	// Make a new hash, and populate it
	newHash := make(map[object.HashKey]object.HashPair)
	newHash[stdoutKey.HashKey()] = stdoutHash
	newHash[stderrKey.HashKey()] = stderrHash

	return &object.Hash{Pairs: newHash}
}

func evalIndexExpression(left, index object.Object) object.Object {
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

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}

	}
	return &object.Hash{Pairs: pairs}

}

func applyFunction(env *object.Environment, fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendEnv)
		return upwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(env, args...)
	default:
		return newError("not a function: %s", fn.Type())
	}

}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	// Set the defaults
	for key, val := range fn.Defaults {
		env.Set(key, Eval(val, env))
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
func evalObjectCallExpression(call *ast.ObjectCallExpression, env *object.Environment) object.Object {

	obj := Eval(call.Object, env)
	if method, ok := call.Call.(*ast.CallExpression); ok {

		//
		// Here we try to invoke the object.method() call which has
		// been implemented in go.
		//
		// We do this by forwarding the call to the appropriate
		// `invokeMethod` interface on the object.
		//
		args := evalExpression(call.Call.(*ast.CallExpression).Arguments, env)
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
				extendEnv := extendFunctionEnv(fn.(*object.Function), args)

				//
				// Now set "self" to be the implicit object, against
				// which the function-call will be operating.
				//
				extendEnv.Set("self", obj)

				//
				// Finally invoke & return.
				//
				evaluated := Eval(fn.(*object.Function).Body, extendEnv)
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
