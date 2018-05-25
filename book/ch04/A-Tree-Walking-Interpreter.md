# 遍历树解释器
我们将要构建一个遍历树解释器，以前面解析步骤完成构建好的抽象语法树，边遍历边进行解释，跳过将预处理和编译的步骤。

我们的解释器将会和经典的Lisp解释器一样，我们采用的的设计受《计算机程序结构和描述》（The Structure and Interpretation of Computer Program）中描述的解释器的影响很大，尤其是关于环境的使用。但是这并不意味着我们想要去复制一个特定的解释器，如果你足够了解的话，我们是使用一个你在其他很多解释器中看得到的蓝图。这也是为什么这种特定的设计很流行的原因。它很容易的着手启动起来，也很容易理解和后期的拓展。

我们只需要执行两部分工作：一是遍历整个树计算；另一个是用我们的宿主语言Go来表达Monkey中的值。计算听上去好像很强大、很宏大。但是它仅仅是一个 `eval` 函数的调用，它这函数的唯一作用就是计算AST的值。接下来就是伪代码的版本来表来在解释器的上下文中什么是计算和树遍历。
```
function eval(astNode){
    if (astNode is integerliteral){
        return astNode.integerValue
    }else if (astNode is booleanLiteral){
        return astNode.booleanValue
    }else if (astNode is infixExpression) {
        leftEvaluated = eval(astNode.Left)
        rightEvaluated = eval(astNode.Right)
        if astNode.Operator == "+" {
            return leftEvaluated + rightEvalueated
        }esle if astNode.Operator == "-" {
            return leftEvaluated - rightEvalueted
        }
    }
}
```

正如你所见到的，`eval` 函数是递归的，如果 `astNode` 是 `infixExpression`，那么调用 `eval` 本身两次来分别计算左边操作数和右边操作数，它将会导致计算另外的中缀表达式计算或者整型计算或者布尔型计算亦或者一个变量。在构建和测试抽象语法树的过程中我们见到过这种递归的形式，在这里我们也是应用同样的概念，唯一区别是我们是计算而不是构建这颗抽象语法树。

从伪代码的片段我们可以大致想象对这函数进行扩展是非常简单的，这也是我们非常擅长的工作。我们将会一点一点地构建我们自己的 `eval` 函数，随着拓展我们的解释器，我们一点点增加新的分支和功能。

在这代码片段里，最好玩的是返回语句。它们到底是返回什么？下面两行是将函数调用的返回值绑定至相应的变量中。
```
leftEvaluated = eval(astNode.Left)
rightEvaluated = eval(astNode.Right)
```
那么它们返回的是什么？返回值的类型又是什么？这个问题的答案就是：在我们的解释器中拥有哪些内部的对象系统？