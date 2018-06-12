**用GO语言编写解析器**

# 翻译计划
由于本人在互联网公司工作，工作时间较长，翻译工作只能在有限的下班时间展开。现在邀请一起翻译，<del>目前所有章节翻译工作尚未有其他人认领<-del>，需要认领章节的，邮箱联系本人，邮件内容包含认领的小节和个人GitHub地址。

**翻译表**  

翻译人员 | 章节
---|---
[gaufung](https://github.com/gaufung) | 1.1 - 1.3
[Jehu Lu](https://github.com/lwhile)  | 2.1 - 2.5
[momaek](https://github.com/momaek)  | 3.1 - 3.9
  

- [1 前言](#ch01-introduction)
    - [1.1 Monkey 编程语言和解释器](#ch01-the-monkey-programming-language-and-interpreter)
    - [1.2 为什么使用 Go 语言](#ch01-why-go)
    - [1.3 如何使用这本书](#ch01-How-to-Use-this-Book)
- [2 词法分析器](#ch02-Lexing)
    - [2.1 词法分析](#ch02-Lexical-Analysis)
    - [2.2 定义Token](#ch02-Defining-Our-Tokens)
    - [2.3 词法分析器](#ch02-The-Lexer)
    - [2.4 拓展Token集和词法分析器](#ch02-Extending-Our-Token-Set-and-Lexer)
    - [2.5 REPL编写](#ch02-Start-of-a-REPL)
- [3 语法解析](#ch03-Parsing)
    - [3.1 语法解析器](#ch03-Parsers)
    - [3.2 为何不采用语法生成器](#ch03-Why-Not-a-Parser-Generator)
    - [3.3 为Monkey编程语言编写语法解析器](#ch04-Writing-a-Parser-for-the-Monkey-Programming-Language)
    - [3.4 解析Let语言](#ch03-Parsing-Let-Statement)
    - [3.5 解析Return语句](#ch03-Parsing-Retrun-Statement)
    - [3.6 解析表达式](#ch03-Parsing-Expression)
    - [3.7 Pratt解析法如何工作](#ch03-How-Pratt-Parsing-Works)
    - [3.8 拓展解析器](#ch03-Extending-The-Parser)
    - [3.9 REPL](#ch03-Read-Parse-Print-Loop)
- [4 计算](#ch04-Evaluation)
    - [4.1 符号赋值](#ch04-Giving-Meaning-to-Symbols)
    - [4.2 计算策略](#ch04-Strategies-of-Evaluation)
    - [4.3 树遍历计算](#ch04-A-Tree-Walking-Interpreter)
    - [4.4 表达对象](#ch04-Representing-Objects)
    - [4.5 表达式计算](#ch04-Evaluaiton-Expression)
    - [4.6 条件语句](#ch04-Conditionals)
    - [4.7 返回语句](#ch04-Return-Statement)
    - [4.8 错误处理](#ch04-Error-Handling)
    - [4.9 绑定和环境](#ch04-Binding-and-Environment)
    - [4.10 函数和函数调用](#ch04-Function-and-Function-Call)
    - [4.11 垃圾回收](#ch04-Trash-Out)
- [5 拓展解释器](#ch05-Extending-the-Interpreter)
    - [5.1 数据类型和函数](#ch05-Data-Type-and-Functions)
    - [5.2 字符串](#ch05-Strings)
    - [5.3 内置函数](#ch05-Built-in-Functions)
    - [5.4 数组](#ch05-Array)
    - [5.5 哈希表](#ch05-Hashes)
    - [5.6 完结](#ch05-the-Grand-Finale)
- [6 资源](#ch06-Resources)
- [7 反馈](#ch07-Feedback)


<h1 id="ch01-introduction">1 前言</h1>
首先第一句话要说的应该是“解释器是具有魔法的”， 一个不愿意透露姓名的早期阅读者说道：”这听上去好像有点傻“。但是我并没有这样认为，我始终坚持解释器非常有魔力。让我一点点告诉你为什么。

表面上来看，解释器看上去误以为很简单：文本写入，得到一些东西出来。他们就是一个程序把其他的程序代码作为输入，并且生成一些东西。是不是很简单， 对吗？但是你越考虑这个问题，你就越觉得这个更加迷人。看上去随机的字符，包括字母、数字或者特殊的符号被输送到解释器后就变得有意义，这些都是解释器赋予的意义。它从无意义中发现意义，电脑只是一个建立在只能理解0和1上的机器，但是却能够理解我们输送的字符并且做出相应的操作，这些都是解释器在读取的过程中进行的翻译。

我曾经不停的问自己：解释器到底是如何工作的？当问题第一次在我脑海中形成的时候，我已近知道只有我自己写一个解释器我才能明白问题的答案。所以我就开始着手进行这件事。

有好多书籍、文章、博客或者教程是关于解释器，但是它们绝大多数涉及两个风格中的其中一个。一是涉及的主题非常宏大，难以置信的理论知识，面向那些已经非常理解这些主题的读者；另外就是非常简短，仅仅提供了简单的介绍，将外部工作当做一个黑盒子并且以玩具版的解释器为关心的重点。

其中一个基础来源就是本书后面的资源，因为解释器仅仅说明了语法简的解释型编程语言。 我并不想走捷径，我确实想知道解释器如何工作的并且理解词法分析器和句法解析器是如何工作的。尤其是类C一样带花括号和分号的编程语言，当我还不知道如何开始解析它们，那些学术上的书籍包含着我要寻找的答案。当然对我来说从哪些冗长的，理论化解释和数学符号中，我很难得到我想要的答案。

我想要的东西是介于一个900页的关于编译器的书和用50行ruby代码写一个Lisp解释器的博客之间的内容。

为了你也包括我，写了这本书。我希望这本书是为了那些喜欢一探究竟的人，亦或者是那些喜欢通过了解一些如何工作的而学习的人。

在这本书中我们将从零开始为我们自己的编程语言写一本解释器。我们将不会使用任何第三方的工作或者库。这些将不会再生产实际中使用，也不会对性能测试上做一些工作。当然，这个解释器支持的编程语言会缺失一些功能，但是我们能够从中学到很多。

由于解释器中的种类繁多并且没有很相像，所以用通用的语句描述解释器非常困难。我们能够说的就是它们共有的基础属性就是读取源代码并且计算它，没有产生一些后续执行的可视化、即可结果。编译器确实恰恰相关的，它读取源代码并且生成背后机器理解的其他代码。

有些解释器非常短小简单，甚至没有涉及解析的的步骤。它们仅仅是立马解析输入，看看类似Brainfuck中的一个就明白我说的意思。

在其他精心设计的解释器中，包含了大量高度优化和使用了先进的解析和计算技术。其中一些不去计算其中的输入，将其中编译成叫做字节码的中间表达代码，然后计算这些字节码。更先进的就是叫做JIT解释器，它将输入编译成本地机器码，然后执行。

但是，在上述表示的两种类别中，有一种解释器能够解析源代码，然后构建一颗抽象语法树(AST)然后计算这棵树。这种类型的解释器有时叫做"tree-walking"解释器，因为它就像在抽象语法树上行走然后解释它。

在这本书中，我们将构建一个"tree-walking"解释器。

我们将会构建我们的词法解析器，语法解析器，树表达式和最后的计算器。我们将会看到什么是token, 什么是抽象语法树，如何去构建这一颗树，如何计算这个树和如何去拓展我们的编程语言和一些内置函数。

<h2 id="ch01-the-monkey-programming-language-and-interpreter">1.1 Monkey 编程语言和解释器</h2>
每一个解释器都是为了解释一种特定的编程语言，这也是你如何实现一门编程语言。如果没有编译器或者解释器，任何一种一种编程语言仅仅就是一些特定的符号而已。

我们将要解析和计算的语言叫 Monkey。 这是专门为本书设计的编程语言，也是本书中的解释器实现的编程语言。

它的所有语言特性列表如下：
- C语言类似语法
- 变量绑定
- 整型和布尔型
- 算术表达式
- 内置函数
- 第一类和高阶函数
- 闭包
- 字符串类型
- 数组类型
- 散列类型

接下来我们将要具体的查看如何去实现上述的每一个功能，再次之前我们先看看 Monkey 是什么样子的。

在Monkey中我们如何将值绑定到名称上

```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

除了整型、布尔型和字符串类型，我们构建的Monkey解释器同样也支持数组和哈希，接下来将展示整型数组是怎样的：
```
let myArray=[1, 2, 3, 4, 5];
```
接下来就是哈希，每个值对应于相应的键：
```
let thorsten = {"name": "Thorsten", "age": 28}
```
可以通过索引表达式来访问数组和哈希表中的元素
```
myArray[0]       // => 1
thorsten["name"] // => "Thorsten"
```

`let` 语句也可以用来将函数绑定变量上，接下来就是一个将连个数字相加的小函数
```
let add = fn(a, b) { return a + b; };
```
但是 Monkey 不但支持 `return` 语句， 而且支持隐式返回值，也就是说我们不使用 `return` 直接返回。
```
let add = fn(a, b) { a + b; };
```
调用函数也非常简单
```
add (1, 2);
```
展示一个更加复杂的函数，比如 `fibonacci` 函数可以返回第N个斐波那契数字。
```
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```
可以看到可以递归地调用 `fibonacci` 函数！

Monkey 也支持一种特殊类型的函数，叫做高阶函数。这些函数能够将其他函数作为参数，接下来就是例子
```
let twice = fn (f, x) {
    return f(f(x));
};
let addTwo = fn(x) {
    return x +  2;
};
twice(addTwo, 2);
```
在这里 `twice` 函数接受两个参数：一个叫做 `addTwo` 的函数，另一个是整数2。 它调用 `addTwo` 两次，第一次用2作为参数，第二次用第一次调用的返回值作为参数，最终生成返回6。
所以，我们可以将函数作为函数调用的参数，函数在Monkey中也就是一个值，就像整数和字符串。这种功能叫做函数是一等公民。

我们在本书中将要构建的解析器将会实现上述所有功能，在REPL中，它首先读取源代码中的token，然后解析它。然后构建一颗内部表达的抽象语法树，然后计算这颗树。它将拥有如下几个主要部分：
- 词法解析器
- 语法解析器
- 抽象语法树
- 内部对象体系
- 计算器

我们将从低到上按照这个顺序进行构建。也许可以这样中，从源代码到。最终的结果输出，但是这种方法的缺点就是在第一章结束后它不能生成一个 `Hello World`。但是它的优势就是可以很简单的理解这个部分如何组合在一起，数据是如何在整个程序中流动的。

但是我们为什么叫它 `Monkey` 呢？ 因为猴子是一个漂亮、优雅、迷人和好玩的生物，这一点跟就跟解释器一样的，也是这本书的名称的原因。


<h2 id="ch01-why-go">1.2 为何选择Go语言</h2>
如果你到目前为止还没有注意到标题的中的词Go，先第一层意思首恭喜你，非常有纪念意义；第二层意思就是我们将要使用Go语言编写解析器。那为什么我们使用Go语言呢？

我喜欢用Go写代码，我喜欢用这门语言和它提供的标准库以及工具，另外考虑的就是Go拥有一些特性对这本中的内容非常适合。

Go非常容易阅读和理解，你不需要理解这本书中编写的Go语言代码，甚至如果你不是有经验的Go语言的程序员，我打赌你能够完全理解这本书中哪怕你从来没有写过一行关于Go语言的代码。

另一个原因就是Go语言提供非常棒的工具，本书中我们编写的解释器关注的重点背后的想法的和概念。通过Go语言的格式化命令 `gofmt` 和内置的测试框架，我们可以专注与我们的解释器不用去关心第三方库、工具和依赖。在这本书中我们将不会使用任何其他的工具，仅仅使用Go语言提供的。

但是我认为更重要的是本书提供的Go语言代码与那些更底层的代码非常类似，比如C，C++和Rust。或者这跟Go语言的本身相关，更注重简洁性,<del>its stripped-down charm and lack of programming language constructs that are absednt in other languages and hard to translate</del>。 或许这也是我为本书选择Go语言的原因。同样原因，这本书中将不会有任何元编程的相关技巧，虽然那样做何以走一些捷径，但是两周之后将没有人能够看懂。没有强大的面向对象设计和模式</del>that need pen, paper and the sentence "actually, it's pretty easy" to explain</del>。

所有的原因都是书中的代码能够让你更好的理解（概念上的和技术层面上的）和重复使用它们。如果你在读完这本书后，打算用其他语言写自己的解释器，那将会是非常容易上手的。通过这本书，我想给你提供一个理解和构造一个解释器的起点。


<h2 id="ch01-How-to-Use-this-Book">1.3 如何使用这本书</h2>
这本书既不是一本参考手册，也不是关于描述实现解释器相关概念的论文的集合。这本书用来从头到尾，按照我推荐的顺序阅读，同时输入和修改提供的代码。

每一个章节是建立在先前章节上的，主要包括代码和内容。在每一章节中我们一点一点地构建我们的解释器。为了使它更容易理解，本书提供了一个叫`code`的文件夹，如果你购买的本书没有改文件夹，你可以从下面的地址下载到

[https://interpreterbook.com/waiig_cod_1.1.zip](https://interpreterbook.com/waiig_cod_1.1.zip)

`code` 文件夹被分成几个子文件夹，每一章节分为一个文件夹，其中包含了相应章节的内容。

有时我仅仅偶尔想起一些代码，但是并没有书中显示这些代码（因为他不仅仅占用了太多的空间，因为它们是一些测试文件的中的测试用例，或者仅仅是一些细节）。你能够在相应章节中找到这些代码。
接下来你需要哪些功能呢？不多，一个文本编辑器和Go编程语言，任何Go语言版本大于1.0即可工作。但是为了将来版本，我进行一些免责申明：我在编写的时候使用的是Go1.7。

我同样推荐使用[direnv](https://direnv.net)，它能根据你的 `.envrc` 文件改变你的shell环境。本书的`code`文件夹中的每一个子文件中都有一个 `.envrc`文件，它用来将`GOPATH` 添加到其子文件下中，它将允许我们不同章节下的代码都能够工作。
让我们开始行动吧！

<h1 id="ch02-Lexing">2 词法分析器</h1>
<h2 id="ch02-Lexical-Analysis">2.1 词法分析</h2>
<h2 id="ch02-Defining-Our-Tokens">2.2 定义Token</h2>
<h2 id="ch02-The-Lexer">2.3 词法分析器</h2>
<h2 id="#ch02-Extending-Our-Token-Set-and-Lexer">2.4 拓展Token集和词法分析器</h2>
<h2 id="#ch02-Start-of-a-REPL">2.5 REPL编写</h2>

<h1 id="ch03-Parsing">3 语法解析</h1>
<h2 id="ch03-Parsers">3.1 语法解析器</h2>
<h2 id="ch03-Why-Not-a-Parser-Generator">3.2 为何不采用语法生成器</h2>
<h2 id="h04-Writing-a-Parser-for-the-Monkey-Programming-Language">3.3 为Monkey编程语言编写语法解析器</h2>
<h2 id="ch03-Parsing-Let-Statement">3.4 解析Let语言</h2>
<h2 id="ch03-Parsing-Retrun-Statement">3.5 解析Return语句</h2>
<h2 id="ch03-Parsing-Expression)">3.6 解析表达式</h2>
<h2 id="ch03-How-Pratt-Parsing-Works">3.7 Pratt解析法如何工作</h2>
<h2 id="ch03-Extending-The-Parser">3.8 拓展解析器</h2>
<h2 id="ch03-Read-Parse-Print-Loop">3.9 REPL</h2>

<h1 id="ch04-Evaluation">4 计算</h1>
<h2 id="ch04-Giving-Meaning-to-Symbols">4.1 符号赋值</h2>
<h2 id="ch04-Strategies-of-Evaluation">4.2 计算策略</h2>
<h2 id="ch04-A-Tree-Walking-Interpreter">4.3 树遍历计算</h2>
<h2 id="ch04-Representing-Objects">4.4 表达对象</h2>
<h2 id="ch04-Evaluaiton-Expression">4.5 表达式计算</h2>
<h2 id="ch04-Conditionals">4.6 条件语句</h2>
<h2 id="ch04-Return-Statement">4.7 返回语句</h2>
<h2 id="ch04-Error-Handling">4.8 错误处理</h2>
<h2 id="ch04-Binding-and-Environment">4.9 绑定和环境</h2>
<h2 id="h04-Function-and-Function-Call">4.10 函数和函数调用</h2>
<h2 id="ch04-Trash-Out">4.11 垃圾回收</h2>

<h1 id="ch05-Extending-the-Interpreter">5 拓展解释器</h1>
<h2 id="ch05-Data-Type-and-Functions">5.1 数据类型和函数</h2>
<h2 id="ch05-Strings">5.2 字符串</h2>
<h2 id="ch05-Built-in-Functions">5.3 内置函数</h2>
<h2 id="ch05-Array">5.4 数组</h2>
<h2 id="ch05-Hashes">5.5 哈希表</h2>
<h2 id="ch05-the-Grand-Finale">5.6 完结</h2>

<h1 id="ch06-Resource">6 资源</h1>
<h1 id="ch07-Feedback">7 反馈</h1>

