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
终于到达这一章：计算。在 `REPL`中的 `E` 就是解释器在处理源代码的的过程中的最后一步。在这一步，源代码将会变得有意义。没有计算，诸如 `1+2`这样的表达式仅仅是一连串的字符，token或者代表表达式的树状结构，而并没有意味着任何事情，当然，我们都知道 `1+2` 变成 `3`, `3.5 > 1` 是 `true`, `5<1` 是 `false`, `put("Hello World!")` 将会输出友好的消息。

计算的过程就是定义解释器如何将编程语言编程解释后的结果。
```
let num = 5;
if (num) {
    return a;
} else {
    return b;
}
```
是否返回 `a` 或者 `b` 将去取决于解释器在计算整数 `5` 是否为真。在一些语言中，它是真值，但是在其他语言中，我们需要用一个表达式来生成布尔值，如 `5 != 0`。

考虑如下
```
let one = fn() {
    printLine("one");
    return 1;
};

let two = fn() {
    printLine("two");
    return 2;
};
add(one(), two())
```
输出的结果是第一个 `one` 还是 `two` 函数亦或者是其他的？ 这个取决于特定的编程语言，最终是取决于解释器的实现，它是如何在一个函数调用的表达式中计算参数的顺序。

这样微小的选择方式在本章中还有很多，我们将决定在 Monkey 编程语言中如何进行工作，并且我们的解释器如何计算源代码。

或许你非常怀疑我告诉你写一个解析器非常好玩，但是请相信我，这事最重要的部分，在这部分Monkey编程语言将会变得有生命力起来，尤其是源代码变得运动起来，仿佛开始呼吸一样。
<h2 id="ch04-Strategies-of-Evaluation">4.2 计算策略</h2>
计算这一部分也是实现解释器过程中变化最多的部分，无论哪一种语言的实现。当计算源代码的时候，有很多不同的策略可以供选择。在本书简单介绍解释器架构的时候，我已经提示到这一点，现在我们手上已经有了抽象语法树（AST），现在问题来了，接下来我们需要做些什么，如何去计算这一棵树？我们接下来看看不同的观点。

在开始之前，值得注意的是解释器和编译器的界限是非常模糊的。常常来讲，解释器通常不离开可执行的环境（与此相关，编译器则脱离可执行环境）它并没有哪些现实中高优化的编程语言那么快。

就跟之前说的，显而易见经典的处理抽象语法树的方法就是解释它，遍历整个语法树，访问树种的每一个节点，并且计算出每一个节点的意义如：输出字符串，加两个数字，执行一个函数内部，等等如此。解释器这样工作的方式也叫“树的遍历”，这也是解释器的原型。有时在计算的过程中也做一些进一步优化处理，来重写抽象语法树（比如移除未使用的变量）或者将其改变成另一种中间形式的表达，对于递归或者循环等计算非常适用。

其他解释器同样遍历整个抽象语法树，与解释抽象语法树本身不同，它们首先将其转换成字节码。字节码同样也是抽象语法树的中间表达形式。这个特定的格式代码各式各样，主要取决于前段和宿主的编程语言。这种中间代码与汇编语言非常相似，可以打赌每一个字节码的定义中包含的 `push` 和 `pop` 等相关的栈操作数。但是字节码并不是原生的机器码，或者汇编语言。它不能被操作系统执行，CPU 也不是解释运行。它仅仅能够被虚拟机解释，这也是一个解释器，就跟 VMWare 和 VirtualBox 一样模拟真正的机器和CPU。虚拟机模式一个能理解这个特定格式的字节码的机器，这种处理方式能够提供一个很好的性能上的优势。

这种策略的变化一点也不影响抽象语法树，除了构建一个语法树，然后解析器直接生成字节码，那么是否现在我们仍然在讨论解释器或者编译器？难道生成字节码然后解释它（或许应该叫做执行）不是编译的一种形式？我想说的是：解释器和编译器之间的界线非常模糊，甚至令人感到糊涂。这样去想：一些编程语言的实现过程是这样的解析源代码，构建抽象语法树，然后将抽象语法树转变成字节码。但是不是在虚拟机上执行相关字节码的相关的操作，而是虚拟机将字节码编译成原生的机器码，仅仅是在执行之前，这种方式叫做 `JIT`解释器或者编译器。

其他编译器或者解释器跳过编程成字节码，它们递归地遍历整个抽象语法树，但是在执行一个特定的分支的所有节点之前将其编译成原生机器码，然后执行，同样也叫做 `JIT`。

一个微小的变种就是一种混合模式，解释器递归的遍历抽象语法树，只有当某一个特定的分支在计算很多次后才将其分支编译成机器码。

是不是很神奇？有很多种方法来完成计算这部分任务，其中有很多交叉和变种。

选择哪一种策略方式很大一部分取决于性能和可移植型的需要，以及你想要这个语言如何被解释。遍历整个语法树和递归的去计算可能是所有方法中最慢的一种，但是很容易去构建、拓展、思考以及可适配性。

将代码转换的字节码的的编译器在处理字节码上是非常快的，但是创建这种编译器也变得非常复杂和困难。将JIT转换成混合模式也需要你同时支持不同的计算机体系结构，一旦如果你想要解释器同时通过在ARM和x86架构的CPU上。

以上所有的实现方法在现实的编程语言中都能够被找到，大部分时候，实现的方式随着编程语言的发展而改变，Ruby就是很好的例子。在1.8以及之前的版本，Ruby的解释器是树遍历的，边遍历抽象语法树边执行，但是从1.9版本开始变成虚拟机架构，现在Ruby解释器解析源代码，构建抽象语法树然后将抽象语法树编译成字节码，并在虚拟机中执行，这一点咋性能上有了很大的提高。

WebKit Javascript的引擎 JavaScriptCore 和它的叫Squirrelfish同样采用抽象语法树遍历和执行的方式，在2008年之后转向了虚拟机和字节码解释。现在该引擎拥有4各部不同阶段的JIT编译，为了取得很好性能上的表现，在不同的阶段对程序进行不同形式的解释。

另一个例子就是Lua， 主要的实现方式是将其编译成字节码，然后在一个寄存器为基础的虚拟机上执行，在12年后，LuaJIT作为另外一种实现实现方式出现了。LuaJIT的实现者 Mike Pall的目标是创建一个尽可能快的Lua编译器。事实同样如此，通过JIT的方式将繁琐的字节码格式转换成不同基础架构的机器码，从各个性能测试来看，LuaJIT比原生的Lua解释器要快，而且不仅仅是一点点快，有时至少快50倍。

所以所有编译器都是从很小的改进空间开始，这也是我们开始要做的原因。有很多方法来构建更快的解释器，但是将不会很容易理解，在这里我们将理解和开始着手构建。
<h2 id="ch04-A-Tree-Walking-Interpreter">4.3 树遍历计算</h2>
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
<h2 id="ch04-Representing-Objects">4.4 表达对象</h2>
等一下，什么？你从来没有说过Monkey是面向对象的编程语言啊！是的，我从来没有说过并且它也不是面向对象的编程语言。那为什么我们需要一个“对象系统”呢，亦或者是“值系统”或者“对象描述”？答案是我们需要定义我们计算返回的值。我们需要一个系统来描述我们抽象语法树种的值或者我们在内存中计算的生成出来的内容。

让我们来看看在接下来的Monkey代码中如何计算值
```
let a = 5;
// [...]
a + a;
```
正如你看到的，我们将一个字面整数绑到 `a` 的变量中。接下来无论发生什么，比如我们遇到了一个 `+` 的表达式，我们需要获取 `a` 绑定的值。 为了计算 `a+a`，我们需要去获取5。在抽象语法树中，它使用 `*ast.IntegerLiteral`对象来表达。但是我们如何去记录这些表达形式在计算剩余的抽象语法树的时候呢？

当在构建一门解释器语言的内部值表达的时候有很多种选择， 这方面的主题的智慧在全世界的关于解释器和编译的代码库中到处都有。每一个解释器都有其自己的实现方式，为了适应解释语言的要求，它们往往与先前的有点点不同。

在解释型语言中我们可以使用宿主语言来描述一些原生的类型（比如整型、布尔型等等），但是不是封装所有类型。在其他语言中值和对象都只能是用指针来表示，然而在其他的却是原生类型和指针混合使用。

为什么会有这么多种类呢？其中之一是宿主语言不同，如何在你的解释语言中描述一个字符串取决于你的实现解释器的语言的如何描述。一个用 Ruby 编写的解释器是不能用同样的方式用C语言来描述。

除此之外，被解释语言不同也是同样的原因。一些解释语言只需要仅仅描述原始的数据类型，比如整型、字符类型或者字节。但是其他的你需要有列表、字典、函数和复合数据类型。这些不同点导致了在值描述提出了不同的需求。

除了宿主语言和被解释的语言， 在设计和实现值描述影响最大的是在计算执行的过程中的速度和内存消耗。如果你想构建一个高效的解释器就不能使用膨胀的值系统；如果你想编写自己的垃圾收集器，你需要考虑如何在系统中跟踪每一个值。但是换一句话说，如果你不关心性能，保持解释器简单明了是非常重要，除非有更高的要求提出。

但是问题在于此， 在宿主语言中的解释语言有许许多多不同的值描述的方法。最好的方法，或许是唯一的方法去了解不同描述方法是去阅读这些流行的解释器的源代码。我发自内心推荐[Wren_source_code](https://github.com/munificent/wren)，它包含了两种不同的值描述方法，可以通过编译条件打开和关闭它们。

除了考虑宿主语言在值描述的问题，还需要考虑如何公开这些值描述给解释语言的使用者，也就是说我们公开的API接口长得什么样？

以JAVA语言举例，提供了基础数据类型（整型，字节，短整型，长整型，单精度浮点型，双精度浮点型，布尔型和字符）和引用数据类型给使用者。基础数据类型没有很大的描述在Java实现中，它们与原生的部分一一映射，但是引用数据类型是宿主语言中复合数据类型的引用。

Ruby语言的使用者不需要接触基础数据类型，类似原生值类型是不存在的，因为在Ruby中所有的都是对象，因此他们被封装到内部实现中。在Ruby内部不区分一个字节类型和一个 `pizza`类的实例对象，他们都是值类型，只不过封装了不同的值。

有许多方法将编程语言的数据类型暴露给使用者，这个去取决于如何设计者语言，以及我们在依次强调的性能要求。如果你不考虑性能问题，所有事情都好办了；但是你考虑了，你需要采取一些更聪明的选择来达到这个目标。

## 对象系统的基础

因为我们对Monkey解释器的性能暂时没有考虑，所以我们选择一种简单的方式：我们将要为我们每一个遇到的类型进行描述为一个 `Object`，它是我么设计的接口，每一个值都被封装到一个结构中，该结构将实现 `Object` 接口。

在新的 `object` 包中，我们定义 `Object` 接口和 `ObjectType` 类型：
```go
//object/object.go
package object
type ObjectType string

type Object interface {
    Type() ObjectType
    Inspect() string
}
```
它看上去很简单，和我们先前在 `token`包以及 `Token`和 `TokenType`类型非常像，不同点就是 `token`是一个结构，而 `Object`是一个接口。原因是每一个值都需要不同的内部表达而这样做很容易定义不同的结构类型而不是将布尔型和整型都放到同样的结构体字段中。

目前在Monkey解释器中我们只有三种数据类型: null，布尔型和整型。让我们开始实现整型并构建我们的值系统。

**Integers**

`object.Integer`类型就跟你想象中的一样小巧：
```go
import (
    "fmt"
)
type Integer struct {
    Value int64
}
func (i *Integer) Inpsect() string { return fmt.Sprintf("%d", i.Value) }
```
无论什么时候我们在源代码中遇到整型字面值，我们首先将其转换成一个 `ast.IntegerLiteral`，然后在计算那个抽象语法树的节点的时候，我们将其转换为一个 `object.Integer`，将它的值存入我们结构中然后这个结构的引用传出去。

为了让 `object.Integer`结构去实现 `object.Object`接口，仍然需要一个 `Type()`方法用来返回它的 `ObjectType`。就像我们在 `token.TokenType`，我们为每一个 `ObjectType`定义一个常量：
```go
// object/object.go
import "fmt"
type ObjectType string 
const (
    INTEGER_OBJ = "INTEGER"
)
```
正如我说的，就像我们在 `token` 包中所做的一样，这样我们就可以将 `Type()` 方法添加到`*object.Integer`结构中：
```go
//object/object.go
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
```
到此为止， 我们完成了 `Integer`类型，接下来进入下一个数据类型：布尔型。

**Booleans**
如果你在本小节中期待一个大的事情，我很抱歉你会失望的，`object.Boolean` 也是一个很简单的东西。
```go
//object/object.go
const (
//[...]
    BOOLEAN_OBJ = "BOOLEAN"
)

type Boolean struct {
    Value bool
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inpsect() string { return fmt.Sprintf("%t", b.Value)}
```
仅仅是封装了一个`bool`型变量的结构。

我们即将结束对象系统的基础内容，在来时我们 `Eval` 函数之前我们需要做一个其他的事情。

**Null**
1965年Tony Hoare 在 ALGOL W 语言中引入了空引用，也叫做他的[百万美元的错误](https://www.infoq.com/presentations/Null-References-The-Billion-Dollar-Mistake-Tony-Hoare)。由于他的引用，空引用导致了无初次的系统奔溃，当一个描述中没有值得时候。Null(在其他语言中也称为nil)没有好的名声。

我自己也思考过，Monkey语言中是否应该有空引用，一方面来讲，不要引入，因为编程语言将会变得安全如果它不允许有空或者空引用，但另一方面来讲，我们并不是重新发明一个轮子，而是去学习一些东西。我发现我自己在处理null的时候导致我再次思考什么时候有机会试用它。就像当我车里有个爆炸性的东西会让我开车更慢更小心一点。这一点让我非常满意的选择能够设计一门编程语言。于是我实现了 `Null` 数据类型同时让我仔细小心当我在后面使用它的时候。
```go
//object/object.go
const (
// [...]
    NULL_OBJ = "NULL"
)
type Null struct {}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (N *Null) Inspect() string { return "null" }
```
`object.Null`与`object.boolean` 和 `object.Integer` 一样，除了它没有封装任何值。它代表了值得缺失。

有了 `object.Null`, 我们的对象系统现在就可以代表布尔型、整型和空类型，对于我们开始`Eval` 函数足够了。
<h2 id="ch04-Evaluaiton-Expression">4.5 表达式计算</h2>
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
<h2 id="ch04-Conditionals">4.6 条件语句</h2>
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
<h2 id="ch04-Return-Statement">4.7 返回语句</h2>
接下来的返回语句，这个在任何标准的计算器都不会出现的，但是`Monkey`有。 它不仅仅在函数体而且作为`Monkey`语言的顶层的语句。 它在任何地方使用都不会有任何影响，因为它不改变任何东西。返回语句将停止后面的所有计算，并且带着它计算的值离开。

这儿有个一个最上层的返回语句：
```go
5 * 5 * 5;
return 10;
9 * 9 * 9;
```
执行这个程序将会返回10. 如果这些语句在一个函数体内，那么调用者将会得到10。最重要的是最后一行代码：`9 * 9 * 9`表达式将永远不会被执行。

有一些不同的方式去实现返回语句，在一些宿主语言中我们可以使用`goto`或者异常等方式。但是在`go`语言中，实现异常捕获非常困难，而且我们也不想采用`goto`这种不简洁的方式。为了支持返回语句，我们会传递一个`返回值`给执行器， 当我们遇到一个 `return`, 我们将会其封装起来，并返回里面的对象，因此我们能够跟踪记录它，以便是否决定是否继续计算。

接下来就是刚刚说的对象`object.ReturnValue`的实现：
```go
// object/object.go
const (
//[...]
    RETURN_VALUE_OBJ = "RETURN_VALUE"
)
type ReturnValue struct {
    Value Object
}
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *RetrunValue) Inspet() string {
    return rv.Value.Inspect()
}
```
仅仅是封装了其他的`object`对象，真正有趣的是当它在使用的时候。

接下来的测试示范了我们在Monkey编程上下文中使用返回语句:
```go
// evaluator/evaluator_test.go
func TestReturnStatement(t *testing.T){
    tests := []struct {
        input string
        expected int64
    }{
        {"return 10;", 10},
        {"return 10; 9;", 10},
        {"return 2 * 5; 9;", 10},
        {"9; return 2 * 5; 9;", 10},
    }
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}
```
为了让测试通过，我们需要在`evalStatement`做一些修改，我们已经在`Eval`函数中增加了一个`*ast.ReturnStatement`分支：
```go
// evalutor/evaluator.go
func Eval(node *ast.Node) object.Object {
// [...]
    case *ast.ReturnStatement:
        val := Eval(node.ReturnValue)
        return *object.ReturnValue{Value: val}
// [...]
}
func evalStatements(stmts []ast.Statement) object.Object {
    var result object.Object
    for _, statement := range stmts {
        result = Eval(statment)
        if returnValue, ok := result.(*object.ReturnValue);ok {
            return returnValue.Value
        }
    }
    return result
}
```
首先需要修改的部分是执行器的`*ast.ReturnValue`，在这我们执行带有return语句的表达式。然后将`Eval`计算得到的结果封装成`object.ReturnValue`对象并记录跟踪它。

在`evalStatement`语句中，我们执行`evalProgramStatement`和`evalBlockStatement`方法来分别计算一系列的语句。我们检查每一步执行的结果是不是`object.ReturnValue`，如果是则停止计算，然后返回那个被封装的值。有一点非常重要，我们不返回`object.ReturnValue`，而是返回它封装的值，这也是用户所需要被返回的值。

但是问题是我们有时候想持续跟跟踪`object.ReturnValues`而不是一遇到就把获取未封装的值。这个在语句块中经常遇到。
```go
if (10 > 1) {
    if (10 > 1){
        return 10;
    }
    return 1;
}
```
这个程序应该返回10，但是在目前我们版本中，它只是返回1，一个小的测试可以确认：
```go
// evalutor/evalutor_test.go
func TestReturnStatements(t *testing.T){
    tests :=[]struct {
        input string
        expected int64
    }{
        {`
        if (10 > 1){
            if (10 > 1){
                return 10;
            }
            return 1;
        }
        `, 10,
        },
    }
}
```
正如我们所期待的，这个测试是失败的
```
$ go test ./evalutor
--- FAIL: TestReturnStatements(0.00s)
    evaluator_test.go:159: object has wrong value. got=1, want=10
FAIL
FAIL monkey/evalutor 0.007s
```
我敢打赌你已经发现我们当前的版本出现的问题，但是还是让我说出来：如果我们有嵌套的语句块（这是在Money语言中完全合法的）我们不能一遇到`object.ReturnValue`就将封装的值取出来，因为我们还要继续跟踪该值，直到我们到达最外面一层语句块，停止执行。

在当前版本中非嵌套的语句块可以很好地执行，但是遇到嵌套的语句块，首先要做的事承认我们不能再继续使用`evalStatement`函数来执行语句块。这也是我们为什么要重新命名`evalProgram`函数让其不是那么泛化。
```go
// evaluator/evaluator.go
func Eval(node ast.Node) object.Objcet {
// [...]
    case *ast.Program:
        return evalProgram(node)
// [...]
}
func evalProgram(program *ast.Program) object.Object {
    var result object.Object
    for _, statement := range program.Statements {
        result = Eval(statement)
        if returnValue, ok := result.(*object.ReturnValue); ok {
            return returnValue.Value
        }
    }
    return result
}
```
为了执行`*ast.BlockStatement`，我们引入了新的函数`evalBlockStatement`:
```go
// evalutor/evalutor.go
func Eval(node ast.Node) object.Object {
//[...]
    case *ast.BlockStatement:
        return evalBlockStatement(node)
//[...]
}
func evalBlockStatement(block *ast.BlockStatement) object.Object {
    var result object.Object
    for _, statement := range block.Statements {
        result = Eval(statement)
        if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
            return reuslt
        }
    }
    return result
}
```
在这里对每一个执行结果，我们显式说明不拆封返回值只是检查其`Type()`方法，如果它是`object.RETURN_VALUE_OBJ`， 我们就简单的返回`*object.ReturnValue`而不是取出其中的`.Value`字段，所以它停止执行可能的外部语句块，像气泡一样一直到达`evalProgram`方法，然后取出被封装的值（在后面设置到函数调用的时候，这部分将会做出一些改变）

测试通过：
```
go test ./evalutor
ok monkey/evalutor 0.007s
```
返回语句完成，我们终于不再是构建一个计算器。由于`evalProgram`和`evalBlockStatement`对我们来讲还是很陌生，我们将继续研究它们。

<h2 id="ch04-Error-Handling">4.8 错误处理</h2>
还记得我们先前返回`NULL`对象， 我说过你现在不用担心这个，过会我们会回来处理的。现在正是处理真正的错误时候，以免后面太迟了。总体上来讲，我们只需要回退一小部分先前的代码。老实地说，我先前并没有一开始就实现错误处理是因为我认为实现表达式比处理处理有趣多了。但是现在我们必须把它加上，否则将来调试起来我们的解释器将会非常笨重。

首先，让我们先定义一下什么是真正的错误处理机制，这个并不是什么用户自定义异常，而是内部的错误处理。是那些错误操作符，不支持的操作运算亦或者是在执行过程中出现的用户或者内部的异常。

至于这些错误处理方式的实现由多种方法，但是大多数都是处理返回语句，原因也很简单，因为错误和返回语句是相似的，都是停止执行一系列语句。

首先我们需要处理一个错误对象
```go
// object/object.go
const (
// [...]
    ERROR_OBJ = "ERROR"
)
type Error struct {
    Message string
}
func (e *Error) Type() OjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message}
```
你可以看到 `object.Error` 真的非常简单， 它仅仅封装了一个`string`对象来处理错误消息。在一个生产的解释器中，我们需要将栈调用信息给这个错误对象，通过增加一些错误的所在的具体行和列信息添加到消息中。这些并不难实现，只需要在词法解析的过程中将行和列号添加到其中去。至于为什么我们的解释器并没有这么做，是因为我们想把事情做得简单一点，我们只需要一个错误消息即可， 它能给我的很多具体的反馈信息并且能够停止执行所有代码。

我们会在几处合适的地方增加错误机制，以便提高我们解释器的容错能力。现在测试函数能够展示我们的错误处理机制如何工作的：
```go
// evalutor/evalutor_test.go
func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5+true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5+true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true+false", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5;true+false;5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10>1) { true+false;}", "unknown operator: BOOLEAN + BOOLEAN"},
		{`if (10 > 1) { 
      if (10>1) {
			return true+false;	
			}
			return 1;
}`, "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}
```

但我们运算测试的时候，又遇到了我们的老朋友`NULL`。
```
$ go test ./evaluator
--- FAIL TestErrorHanding (0.00s)
---
FAIL
FAIL monkey/evaluator 0.007s
```
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

