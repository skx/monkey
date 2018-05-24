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