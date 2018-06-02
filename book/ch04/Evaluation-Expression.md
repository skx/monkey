# 计算表达式
好了， 开始编写我们的 `Eval` 函数。我们现有拥有了抽象语法树和赞新的对象系统，这些让我们开始记录我们在执行Monkey代码的时候遇到的值，是时候考试计算抽象语法树。

这是我们 `Eval` 函数签名的第一版：
```go
func Eval(node ast.Node) object.Object
```
函数以一个 `ast.Node` 作为输入，返回一个 `object.Object`对象。要注意我们在 `ast`包中定义的每一个节点都实现了 `ast.Node` 接口，因此它们都可以传递给 `Eval` 函数。所以它允许我们在计算部分抽象语法树的时候递归地调用它自己。抽象语法树的节点需要不同形式的计算方式，而 `Eval`则决定如何去判别这些形式。举一个例子，我们传递一个 `*ast.Promgram`节点到`Eval`，那么 `Eval`应该做些做些什么去计算每一个 `*ast.Program.Statements`通过调用自身的时候每一个语句，返回值是我们在算计最后一个时候的返回值。

我们以实现自计算的表达式开始，也就是我们称之为字面计算值，简单来说就是布尔型和整型。它们是Monkey语言的基础也是非常容易去计算的，因为它们就是计算自身。如果我在`REPL`中输入5，那么5也是要输出的，如果我输入`true`，那么我将得到`true`。

听上去很简单？的确如此，让我们 “输入5， 得到5”变成现实。

**Integer Literals**
在开始写代码之前，想想它究竟意味着什么？ 我们将一个表达式语句作为一个输入，它只包含一个整型字面值，然后将其计算出来并返回。

转换为我们系统的语言就是，提供一个`*ast.IntegerLiteral`，我们的 `Eval` 函数应该返回一个`*object.Integer`对象，该对象包含一个 `Value` 字段，而且该值等于 `*ast.IntegerLiteral.Value`中的整型值。

我们很容易为新的`evaluator`包写出我们测试框架：
```go
// evaluator/evalutator_test.go

import (
    "monkey/lexer"
    "monkey/object"
    "monkey/parser"
    "testing"
)
func TestEvalIntegerExpression(t *testing.T){
    tests := [] struct {
        input string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
    }
    for _, tt:=range tests{
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}
func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program:=p.ParseProgram()
    return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool{
    result, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("object is not Integer, got=%T (%+v)", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("object has wrong value. got=%d, want=%d", result.VAlue, expected)
        return false
    }
}
```
一点也不奇怪，这就是我们刚刚说的内容，除了它目前还不能工作。测试依然是失败的因为 `Eval`函数返回的是 `nil` 而不是 `*object.Integer`
```
$ go test ./evaluator
--- FAIL: TestEvalIntegerExpression (0.00s)
    evalutor_test.go:36: object is not Integer. got=<nil>(<nil>)
    evalutor_test.go:36: object is not Integer. got=<nil>(<nil>)
FAIL
FAIL    monkey/evalutor     0.006s
```
失败的原因是我们从未遇到`*ast.IntegerLiteral`在`Eval`， 我们并没有遍历整个抽象语法树，我们应当从树的顶端开始，接受一个`*ast.Program`，然后递归的遍历每一个节点，但是事实上我们并没有这么做，而是仅仅等待一个 `*ast.IntegerLiteral`。 接下来的修改就是真正地遍历这颗树然后计算`*ast.Program`中的每一个语句。
```go
// evalutor/evalutor.go
func Eval(node ast.Node) object.Object {
    switch node := node.(type){

        //Statements
        case *ast.Program:
            return evalStatements(node.Statements)
        case *ast.ExpressionStatement:
            return Eval(node.Expression)
        
        //Expressions
        case *ast.IntegerLiteral:
            return &object.Integer{Value: node.Value}
    }
    return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
    var result object.Object
    for _, statements := range stmt {
        result = Eval(statements)
    }
    return result
}
```
在我们的`Monkey`程序中计算每一个语句，如果语句是一个`*ast.ExpressionStatement`，我们计算它的表达式，它反映的就是我们从一行输入如`5`的抽象语法树，它是一个只包含一个语句的程序，一个包含一个整型字面表达式表达式的语句（不是返回语句或者声明语句）。
```
$ go test ./evalutor
ok  monkey/evalutor 0.006s
```
好的，现在测试通过了，我们可以计算整型字面值了！*大家好，如果我么输入一个数字，通过几千行代码，我们就可以将其输出出来，测试也是同样如此*，但是这些与我们想象中的还是不是很像，不过就才是简单的开始，我们将看到如果进行计算工作并且如果去拓展我们的计算器。`Eval`的结构不会改变，我们仅仅是增加或者拓展它。

接下来要做的是布尔型字面值得计算，但是在开始做之前，我们应该先庆祝我们的一个计算的成功并且犒劳自己，让我们先把`REPL`中的`E`功能完成。

**完成REPL**
到现在为止，在我们的 `REPL`中的中的 `E`是缺失的，而且现在我们也仅仅只有 `REPL`（Read-Pare-Print-Loop）。现在我们有了`Eval`就可以构建一个真正的 `REPL`

在 `repl`包中使用计算器就跟你想象中的一样简单：
```go
//repl/repl.go
import (
// [...]
    "monkey/evaluator"
)
//[...]
func Start(in io.Reader, out io.Writer){
    scanner := buffio.NewScanner(in)
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)
        program := p.ParseProgram()
        if len(p.Erros()) != 0 {
            printParseErrors(out, p.Error()
            continue
        }
        evaluated := evalutor.Eval(program)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}
```
不是输出`program`(抽象语法树返回的值)，我么将`program`传递给 `Eval`函数，如果 `Eval` 函数返回一个非空的值，也就是 `*object.Object`对象，我们使用`Inspect()`方法将其输出。在`*object.Integer`中输出的结果就是其封装的整数值的字符串。

我们使用 `REPL`可以这样工作：
```
$ go run main.go
Hello mrnugget! This is the monkey programming language!
Feel free to type in commands
>>5
5
>>10
10
>>999
999
```
是不是感觉很不错？词法解析、语法解析、计算都在这里。


**布尔字面值**
布尔字面值就跟我刚刚遇到的整型一样，用来计算它们自己：`true`计算返回`true`,`false`计算返回`false`。在`Eval`中实现这个就跟实现整型字面值一样简单，接下来就是枯燥的测试。
```go
// evaluator/evaluator_test.go
func TestEvalBooleanExpression(t *testing.T) {
    tests := []struct {
        input string
        expected bool
    }{
        {"true", true},
        {"false", false},
    }
    for _, tt := range tests {
        evaluated := testEval(tt.intput)
        testBooleanObject(t, evaluated, tt.expected)
    }
}
func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
    result, ok := obj.(*object.Boolean)
    if !ok {
        t.Errorf("object has wrong value. got=%T(%+V)", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
        return false
    }
    return true
}
```
以后我们会拓展`tests`切片以便能够支持更多的表达式而不是布尔型。现在我们只需要确保当我们输入`true`和`false`能够得到正确的输出。这个测试目前当然是失败的：
```
$ go test ./evaluator
--- FAIL: TestEvalBooleanExpression (0.00s)
    evaluator_test.go:42: Eval didn't return BooleanObject. got=<nil>(nil)
    evaluator_test.go:42: Eval didn't return BooleanObject. got=<nil>(nil)
FAIL
FAIL mongkey/evalutor 0.006s
```
为了使测试通过也非常简单，只需要将`*ast.IntegerLiteral`分支拷贝过来并做一些简单的修改即可：
```go
//evaluator/evalutor.go
func Eval(node ast.Node) object.Object {
//[...]
    case *ast.Boolean:
        return &object.Boolean{Value: node.Value}
}
```
让我们看看在`REPL`如果进行工作的
```
$ go run main.go
Hello mrnugget! This is the monkey programming language!
Feel free to type in commands
>> true
true
>> false
false
>>
```
完美!但是等等，我们是不是在这里每一个`true`或者`false`的时候都创建一个`object.Boolean`对象是不是有点不对劲？在两个`true`或者`fasle`内部之间是并没有什么不同的，但是我们为什么每次都使用新的实例对象呢？在这里只有连个不同的值，因此在这里我们只需要引用即可，而不是创建一个新的。
```go
// evaluator/evalutor.go
var (
    TRUE = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)
func Eval(node ast.Node) object.Object {
// [...]
    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
// [...]
}
func nativeBoolToBooleanObject(input bool) *object.Boolean {
    if input{
        return TRUE
    }
    return FALSE
}
```
现在在我们包中只有两个`object.Boolean`对象实例：`TRUE`和`FALSE`，我们引用它们而不是申请空间去创建它们。这个对我们性能小小提升很具有意义，而这些并不需要很多的工作。我们在`null`类型中同样这么处理的。


**Null**

就跟布尔型`true`和`false`各只有一个一样，对于 `null` 类型也应该只有一个，没有其他空类型的变种，没有任何其他乱七八糟的空类型，只有一个空类型。一个对象要么为空，要么不为空。所以我们首先创建一个`NULL`对象，以便我们在整个计算过程中都能引用它。
```go
// evaluator/evaluator.go
var (
    NULL = &object.NUll{}
    TRUE = &object.Boolean{Value:true}
    FALSE = &object.Boolean{Value:false}
)
```
因此我们只有一份`NULL`引用。

有了整型字面值和三个`NULL`, `TRUE` 和 `FALSE`，我们可以开始准备计算我么你操作表达式。

**前缀表达式**

在Monkey中中最简单的操作数表达式就是就前缀表达式，即单操作数表达式，也就是操作数跟在操作符之后。在我们先前解析过程中提到过，很多语言构造喜欢采用前缀表达式，因为这样解析它很简单。但是在本小节中的前缀表达式就是由一个操作数和一个操作符组成的操作符表达式。Monkey语言支持两种前缀操作数:`!`和`-`。

计算操作符表达式（尤其是前缀操作和操作数）并不难，我们一点点实现构建我们设计出来的行为。但是我们要特别注意的是，我们我们想要达成的结果远超我们预期。要记住，在计算的过程中处理输入语言的意义，我们定义了Monkey语言的文法，一个微小的改变在计算操作符表达式的时候会导致一系列无法预知的问题，测试能帮助我们判断是我们想要的结果。

首先我们开始实现支持`!`操作符，测试展示了这个操作符它的操作符变成一个布尔值，然后取反。
```go
// evaluator/evalutor_test.go

func TestBangOperator(t *testing.T){
    tests := []struct{
        input string
        expected bool
    }{
        {"!true", false},
        {"!false", true},
        {"!5", false},
        {"!!true", true},
        {"!!false", false},
        {"!!5", true}
    }
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testBooleanObjec(t, "evaluted", tt.expected)
    }
}
```

正如我说的，这就是我们让着语言如何工作方式，`!true`和`!false`表达式和它们的期望值看上去很合理，但是`!5`其他语言设计者觉得应该是返回一个错误，但是我们想要说的是`5`行为上表现就是`truthy`.

这个测试肯定不能通过，因为 `Eval`返回一个`nil`而不是`TRUE`或者`FALSE`. 计算前缀表达式第一个就是计算他的操作数，然后用它的操作符来计算结果。
```go
// evalutor/evalutor.go
func Eval(node ast.Node) object.Object {
//[...]
    case *ast.PrefixExpression:
        right :=Eval(node.Right)
        return evalPrefix(node.Operator, right)
}
```
在第一步调用后，右边的值可一个是 `*object.Integer` 或者是 `*object.Boolean`甚至可能是一个 `NULL`。 我们按着右边的操作数然后传递到`evalPrefixExpression`函数中，它用来检查这个操作符是否支持。
```go
// evalutor/evaluator.go
func evalPrefixExpression(operator string, right object.Object)object.Object{
    switch operator{
        case "!":
            return evalBangOperatorExpression(right)
        default:
            return NULL
    }
}
```
如果操作符不支持就返回NULL，这是最好的选择吗？或许是，或许也不是。但是到目前为止，这显然是最好的选择，因为我们还没有实现其他任何错误方式。

`evalBangOperatorExpression` 函数及时 `!` 操作数具体的操作。
```go
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
```
当然，我们的测试全部通过
```
$ go test ./evalutor
ok  monkey/evaluator 0.007s
```

让我们继续 `-` 前缀操作符，我们可以拓展我们的 `TestEvalIntegerExpression` 测试函数以纳入一下用例：
```go
//evalutor/evalutor_test.go
func TestEvalIntegerExpression(t *testing.T){
    tests := []struct {
        input string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
        {"-5", -5},
        {"-10", -10},
    }
// [...]
}
```

为了测试`-`前缀表达式， 我选择拓展测试而不是重新编写一个