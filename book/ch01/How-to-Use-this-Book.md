# 如何使用这本书
这本书既不是一本参考手册，也不是关于描述实现解释器相关概念的论文的集合。这本书用来从头到尾，按照我推荐的顺序阅读，同时输入和修改提供的代码。

每一个章节是建立在先前章节上的，主要包括代码和内容。在每一章节中我们一点一点地构建我们的解释器。为了使它更容易理解，本书提供了一个叫`code`的文件夹，如果你购买的本书没有改文件夹，你可以从下面的地址下载到

[https://interpreterbook.com/waiig_cod_1.1.zip](https://interpreterbook.com/waiig_cod_1.1.zip)

`code` 文件夹被分成几个子文件夹，每一章节分为一个文件夹，其中包含了相应章节的内容。

有时我仅仅偶尔想起一些代码，但是并没有书中显示这些代码（因为他不仅仅占用了太多的空间，因为它们是一些测试文件的中的测试用例，或者仅仅是一些细节）。你能够在相应章节中找到这些代码。
接下来你需要哪些功能呢？不多，一个文本编辑器和Go编程语言，任何Go语言版本大于1.0即可工作。但是为了将来版本，我进行一些免责申明：我在编写的时候使用的是Go1.7。

我同样推荐使用[direnv](https://direnv.net)，它能根据你的 `.envrc` 文件改变你的shell环境。本书的`code`文件夹中的每一个子文件中都有一个 `.envrc`文件，它用来将`GOPATH` 添加到其子文件下中，它将允许我们不同章节下的代码都能够工作。
让我们开始行动吧！