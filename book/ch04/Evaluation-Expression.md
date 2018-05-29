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