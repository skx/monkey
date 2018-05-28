# 对象描述
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

有了 `object.Null`, 我们的对象系统现在就可以代表布尔型、整型和空类型，对于我们开始`Eval` 函数足够了