# 条件语句
你将会惊奇在我么的计算器中实现条件语句如此简单，实现他们唯一的难点就是知道在何时实现他们，整个条件语句最关键点就是在于条件判断，而且计算过程与条件判断息息相关，考虑到这种情况：
```go
if (x>10){
    puts("everything okay!")
}else{
    put("x is too low!")
    shutdownSystem()
}
```
在计算`if-else`表达式中最重要的事是计算正确的分支，如果条件满足为true，我们不必计算`else`分支，只计算`if`分支即可，如果不满足，只需要计算`else`分支即可。

换句话说，也就是我们计算`else`分支是当计算条件`x>10`为...? 好像并不是非常准确，难道我们要计算`everyhing okay!`分支吗，当这个条件表达式为`true`或者它能推导出类似`true`的值吗，也就是非假或者非空？

这是最难的部分，因为这是设计选择，语言的选择部分必须要正确，正确处理代码序列。

在Monkey语言例子中，条件语句将会被执行，当条件为类`true`的时候：
```go
let x = 10;
if(x){
    puts("everything okay");
}else{
    puts("x is too high");
    shutdownSystem();
}
```
在上述的例子中，`everything okay`将会被答应出来。为什么呢？ 因为`x`将会被绑定到10， 计算10并且10不是空，也不是假值。 这就是条件语句在Monkey语言中如何工作的。

跟先前讨论一样，我们先增加一些测试用例：
```go
func TestIfElseExpression(t *testing.T){
    tests := []struct {
        input string
        expected interface{}
    }{
      	{"if (true) {10}", 10},
		{"if (false) {10}", nil},
		{"if (1) {10}", 10},
		{"if (1<2) {10}", 10},
		{"if (1<2) { 10} else {20}", 10},
		{"if (1>2) {10} else {20}", 20},
    }
    for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testDecimalObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
```
这个测试函数的作用我们还没有讨论，当条件没有没有被执行，它的返回值就是NULL:
```go
if (false) { 10 }
```
因为`else`分支丢失，因此条件语句将会生成`NULL`。

为了做一些类型判断和转换，我们在`field`字段中修改了部分内容。但是测试用例非常简单易读，很清楚地反应了我们想要做些什么。当然测试肯定是失败的，因为我们并没有返回`*object.Integer`或者`NULL`：
```shell
$ go test ./evalutor
--- FAIL: TestIfElseExpression(0.00s)
....
FAIL
FAIL monkey/evalutor 0.007s
```
先前我告诉你，你将会非常吃惊，支持条件语句实现过程如此简单，现在信不信我了？现在只需要增加一些代码，以便测试通过：
```go
//evalutor/evalutor.go
func Eval(node ast.Node) object.Object {
// [...]
    case *ast.BlockStatement:
        return evalStatments(node.Statements)
    case *ast.IfExpression:
        return evalIfexpression(node)
// [...]
}
func evalIfExpression(ie *ast.IfExpression) object.Object{
    condition := Eval(ie.Condition)
    if isTruthy(condition){
        return Eval(is.Consequeces)
    }else if ie.Alternaive != nil {
        return Eval(is.Alternative)
    }else{
        return NULL
    }
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
```
正如我先前所说的，难点就在于决定计算那个分支，所有的决定分支部分在封装到逻辑步骤非常清楚的`evalIfExpresion`函数中，`isTruthy`是相等判断式。上述两个函数也增加了`*ast.BlockStatement`条件分支，因为`*ast.IfExpression`中的`.Consequences`和`.Alternative`都是语句块。

我们增加了两个新的具体函数来表达Monkey语言的相关的语法。重用了已经存在的函数，仅仅增加一些以便支持条件语句，然后我们的测试通过。现在我们的解释器支持`if-else`表达式。我们现在离开简单的计算器领域，开始像编程语言进军：
```
Hello mrnugget! This is the monkey programming language!
Feel free to type in commands
>> if (5*5+10>34) { 99 } else { 100 }
99
>> if ((100/2) + 250 * 2==1000){9999}
9999
```