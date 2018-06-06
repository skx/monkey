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

为了测试`-`前缀表达式， 我选择拓展测试而不是重新编写一个测试方法主要是两个原因：一是整型是唯一支持`-`操作符的前缀操作数；第二是测试方法应该包含所有整型的运算方法以便达到清晰的目的。

我们已经提前拓展了 `evalPrefixExpression` 方法以便能让测试通过，只需要在switch语句下添加新的分支：
```go
// evaluator/evalutor.go
func evalPrefixExpression(operator string, right object.Object) object.Object {
    switch operator{
    case "!":
        return evalBangOperatorExpression(right)
    case "-":
        return evalMinusPrefixOperatorExpression(right)
    default:
        return NULL
    }
}
```
`evalMinusPrefixOperatorExpresion` 函数看上去像这样：
```go
// evalutor/evalutor.go
func evalMinusPrefixOperatorExpression(right object.Object)object.Object {
    if right.Type() != object.INTEGER_OBJ {
        return NULL
    }
    val := right.(*object.Integer).Value
    return &object.Integer{Value : -value}
}
```
首先先检查操作数是否为整数，如果不是，返回NULL；如果是，我们提取 `*object.Integer`中的值，然后重新分配一个对象来封装它的相反值。

是不是只需要做简单的事，但是他的确能够工作:
```
$ go test ./evaluator
ok  monkey/evalutor 0.0007s
```
棒极了，在继续中缀表达式之前，我们可以在REPL中给出我们前缀表达式的值：
```
$ go run main.go
Hello mrnugget! This is the Monkey programming language！
Feel free to tyoe in comamnds
>> -5
5
>> !true
false
>> !-5
false
>> !! -5
true
>> !!!! -5
true
>> -true
null
```

**中缀表达式**
作为一个新人，下面是八个Monkey语言支持的的中缀表达式
```go
5 + 5;
5 - 5;
5 * 5;
5 / 5;

5 > 5;
5 < 5;
5 == 5;
5 != 5;
```
这八个操作符可以被分为两组，一组的操作符生成布尔型值作为结果，另外一组生成。我们开始实现第二组操作符 `+, -, *, /`。刚开始我们只支持整数操作数，只要他能工作，我们就支持操作数两边是布尔型值。

测试框架已经准备就绪，我们仅仅只要拓展 `TestEvalIntegerExpression`测试方法来适应新的操作符。
```go
//evalutor/evalutor_test.go
func TestEvalIntegerExpression(t *testing.T){
    tests :=[] struct {
        input string
        expected int64
    }{
       {"5", 5},
		{"10", 10},
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5+5+5+5-10", 10},
		{"2*2*2*2*2", 32},
		{"-50+100+ -50", 0},
		{"5*2+10", 20},
		{"5+2*10", 25},
		{"20 + 2 * -10", 0},
		{"50/2 * 2 +10", 60},
		{"2*(5+10)", 30},
		{"3*3*3+10", 37},
		{"3*(3*3)+10", 37},
		{"(5+10*2+15/3)*2+-10", 50},
    }
//[...]
}
```
这里有些测试用例可以删除，因为他们与其他的很有些重复并且增加了一些新的。实话来讲，我是非常高兴当这些测试通过的时候，我明白了我们所做的工作完成了，忍不住问自己，是不是这么简单？实际情况是的确这么简单。

为了让这些测试用例通过，首先要做的是拓展在`Eval`函数中的`switch`语句：
```go
//evaluator/evalutor.go
func Eval(node ast.Node) objet.Object{
// [...]
    case *ast.InfixExpression:
        left := Eval(node.Left)
        right := Eval(node.Right)
        return evalInfixExpression(node.operator, left, right)
// [...]
}
```
就跟`*ast.PrefixExpression`一样，我们首先计算操作数。 现在我们有两个操作数，左边和右边各一个抽象语法树节点，我么已经知道，这可能是其他的表达式、一个函数调用、一个整型字面值、一个操作符表达式等等。我们不去关心这个，让`Eval`函数去关心这个。

在计算完操作数之后，我们将两个返回值和操作符传递到 `evalIntegerInfixExpressions`函数中去，函数是这样子的：
```go
func evalInfixExpression(
    operator string,
    left, right object.Object,
)object.Object {
    switch {
    case left.Type()==object.INTEGER_OBJ && right.Type()==object.INTEGER_OBJ:
        return evalIntegerInfixExpression(operator,left,right)
    default:
        return NULL
    }
}
```
正如我刚刚保证的，一旦两边的操作数都不是整数的时候，我们就返回`NULL`, 当然我们后面将要拓展我们的函数，但是为了测试通过，这就足够了。重点在于`evalIntegerInfixExpression`函数中，在该函数中，我们封装的`*objet.Integers`的操作有加、减、乘和除。
```go
// evalutor/evalutor.go
func evalIntegerInfixExpression(
    operator string
    left, right object.Object,
)object.Object{
    leftVal := left.(*object.Integer).Value
    rightVal := right.(*object.Integer).Value
    switch operator{
    case "+":
        return &object.Integer{Value:leftVal+rightVal}
    case "-":
        return &object.Integer{Value:leftVal-rightVal}
    case "*":
        return &object.Integer{Value:leftVal*rightVal}
    case "/":
        return &object.Integer{Value:leftVal/rightVal}
    default:
        return NULL
    }
}
```
现在信不信由你，测试通过了：
```
$ go test ./evalutor
ok monkey/evalutor 0.007s

好的，我们继续前进，我们过会儿会再回到这边，以便支持哪些能够生成布尔值的操作符`==`,`!=`,`<`和`>`。

我们可以拓展我们的`TestEvalBooleanExpression`方法，为上述的操作符增加测试用例，*因为它们都能生成布尔型值。
```go
//evaluator/evaluator_test.go
func TestEvalBooleanExpression(t *testing.T){
    tests := []struct {
        input string
        expected bool
    }{
        {"true", true},
		{"false", false},
        {"1<2", true},
		{"1>2", false},
		{"1<1", false},
		{"1>1", false},
		{"1==1", true},
		{"1!=1", false},
		{"1==2", false},
		{"1!=2", true},
    }
}
```
除此之外，我们还需要增加一些代码在 `evalIntegerInfixExpression`函数中，它们能够保证测试能够通过：
```go
// evaluator/evaluator.go
func evalIntegerInfixExpression(
    operator string
    left, right object.Object,
) object.Object {
    leftVal := right.(*object.Integer).Value
    rightVal := right.(*obejct.Integer).Value
    switch operator {
// [...]
    case "<":
        return nativeBoolToBooleanObject(leftVal < rightVal)
    case ">":
        return nativeBoolToBooleanObject(leftVal > rightVal)
    case "==":
        return nativeBoolToBooleanObject(leftVal == rightVal)
    case "!=":
        return nativeBoolToBooleanObject(leftVal != rightVal)
    default:
        return NULL
    }
}
`nativeBoolToBoolean`方法是我们用在布尔字面值的时候，现在我们在比较未封装的值比较的过程中又重新使用了它们。

至少对于整数来说，我们现在已经完全支持八种中缀操作符，剩下的工作就是支持布尔型操作数。

Monkey语言支持的布尔操作数是相等判别符`==`和`!=`。它不支持布尔型数值的加减乘除。检查`true`是否比`false`大，`<`和`>`能够做或运算都是不支持的，这些都减少了我们的工作量。

首先我么你需要做的事情，你知道的，就是增加测试内容，就跟以前一样，我们拓展已有的测试方法，我们使用`TestEvalBooleanExpression`函数并增加`==`和`!=`操作符的测试用例。

```go
func TestEvalBooleanExpression(t *testing.T){
    tests := []struct{
        input string
        expected bool
    }{
// [...]
        {"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"(1<2)==true", true},
		{"(1<2) == false", false},
		{"(1>2) == true", false},
        {"(1>2)==false", true},
    }
// [...]
}
```
严格来讲，只需要五个测试用例就足够了，但是我们又增加了四个测试用例来检查生成布尔型值得比较。

到目前为止，没有任何惊奇的地方，仅仅又是一系列失败的测试用例
```
$ go test ./evalutor
--- FAIL: TestEvalBooleanExpression (0.00s)
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
evalutor_test.go:121: object is not Boolean. got=*object.Null {&{}}
FAIL
FAIL monkey/evalutor 0.007s
```
接下来就是增加一些简单的东西让测试通过：
```go
//evalutor/evalutor.go
func evalInfixExpression(
    opertor string 
    left, right object.Obejct,
)object.Object {
    switch {
// [...]
    case operator == "==":
        return nativeBoolToBooleanObject(left == right)
    case oeprator == "!=":
        return nativeBoolToBooleanObject(left != right)
    default:
        return NULL
    }
}
```
是的，我们只是在已经存在的`evalInfixExpression`函数中增加四行代码，测试就通过了，我们通过指针的比较来检查两个布尔型值之间的相等。这样做的原因是我们指向布尔型的指针只有两个`TRUE`和`FALSE`，如果有其他值也是`TRUE`，也就是内存地址一样，它就是`true`。对于`NULL`也是同样的道理。

但是对于整型或者其他数据类型并不奏效，因为对于`*object.Integer`我们每次都分配内存来生成`object.Integer`实例因而我们就能有新的指针。我们无法比较指向不同实例的指针。我们无法比较指向不同实例的指针，否则像`5==5`就会是false，而这个并不是我们想要的。在这种情况下，我们要明确的指出是比较值而不是封装值得对象。

这也是为什么我们在`switch`语句中首先检查整型部分的分支，它们将会首先进行匹配，只要我们留心其他啊操作数的类型是在指针比较之前，我们的成果就能很好的工作。

想象一下，如果十年之后，Monkey语言变得出名了，许多人开始来研究讨论。忽视了这个半吊子地设计这门语言，那么我们就变得非常出名了。 有人就是去StackOverflow上去问，为什么在Monkey语言中，整型比较会比布尔型比较慢得多。你和我，或者其他人就会回答到，因为在Monkey语言中，不允许使用指针比较整型数据，在比较之前，需要拆封它们的值，然后进行比较。相对而言，布尔型操作比较就比较快。 我们将会加上答案的最后加上这么一句”因为源代码是我写的“。

有点跑题了，回到正题。我们做到了，相当高兴，差不多都可以开始庆祝了。是时候开香槟庆祝了吗？对的，看看我们的解释器现在能做什么！
```
$go run main.go
Hello mrnugget! This is the monkey programming language!
Feel free to type in commands
>> 5 * 5 + 10
35
>> 3 + 4 * 5 == 3 * 1 + 4 * 5
true
>> 5 * 10 > 40 + 5
true
>> (5 > 5 == true) != false
false
>> 500 / 2 != 250
false
```
到目前为止，我们已经完成了一个函数计算器， 接下来让我继续增加，使它变得更像一个编程语言。