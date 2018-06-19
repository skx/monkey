# Monkey

This repository contains an interpreter for the "Monkey" programming
language, as described in [Write an Interpreter in Go](https://interpreterbook.com).

The code was originally written by [gaufung](https://github.com/gaufung/Monkey), but has been updated by [skx](https://github.com/skx/Monkey):

* Added single & multi-line comments.
* Added postfix operators (`i++`, `i--`).
* Allow accessing bytes of a string via the index-operator
* Added a driver to read from STDIN, or a named file.
* Added a collection of standard-library functions.
    * Including file input, type-discovery, string functions, etc.
* Added a new way to define functions, via `function`.
* Added `<=` + `>=` comparison functions.
* Scripts can access their command-line arguments.
* String interpolation is supported.

# 0 Installation

If you have a working [golang](https://golang.org/) setup you can
install the intepreter via:

   $ go get -u  github.com/skx/Monkey
   $ go install github.com/skx/Monkey


# 1 Introduction
 source code.

[中文翻译](book/README.md)

**Monkey** interpreter language which is implemented Go language.
```
$ ./monkey ./examples/stdin.mon
Enter your name:Steve
Hello, Steve
```

You can also directly execute a script, if you've installed `monkey`:

    $ chmod 755 examples/arguments.mon
    $ ./examples/arguments.mon test me
    We received 3 arguments to our script.
       0 ./examples/arguments.mon
       1 test
       2 me

# 2 Syntax

**NOTE**: Example-programs can be found beneath [examples/](examples/).


## 2.1 Definition
using `let` as keyword, each line ends with `;`.
```
>>>let a = 3;
>>>let b = 1.2;
>>>a+b
4.2
```

## 2.2 Arithmetic operations
`monkey` supports all basic arithmetic operation of `int` and `float` types. `int` type is represented by `int64` and `float` type is represented by `float64`.

```
>>> let a = 3;
>>> let b = 1.2;
>>> a + b
4.2
>>> a - b
1.8
>>> a * b
3.6
>>> a / b
2.5
```

The complete list of operators includes:

* `+`, `-`, `*`, `/`

When operating with integers the modulus operator is available too `%`.


## 2.3 Builtin containers
`monkey` contains two builtin containers: `array` and `map`.
- array

array is a list which organizes items by linear sequence. But types of items can be different from each other.

```
>>> let a = [1, 2.3, "array"];
>>> a
[1, 2.3, array]
>>> let b = push(a, "another");
>>> b
[1, 2.3, array, another]
```

- map

map is treated as `key-value` container. please attention to that only `boolean`, `int` and `string` types can be used as key.

```
>>> let a = {"name":"monkey", true:1, 7:"seven"};
>>> a
{name: monkey, true: 1, 7: seven}
>>> a["name"]
monkey
>>> a[true]
1
>> a[7]
seven
>>>let b = set(a, 8, "eight");
>>> b
{name: moneky, true: 1, 7: sevent, 8: eight}
```

## 2.4 Builtin functions

- `len`

yield the length of builtin containers.

- `first`

yield the first element of array.

- `last`

yield the last element of array.

- `rest`

yield an array which excludes the first element.

- `push`

push an elements into the array.

- `set`

insert key value pair into the map.

- `puts`

print literal value of objects.

- `string`

convert the given item to a string.

- `int`

convert the given float/string to an integer.  (Useful when using `math.random()`, etc.)

- `type`

returns the type of a variable.


## 2.4.1 The Standard Library

In addition to the built-in functions which are documented above we also
have a minimal-standard library.  The library includes some string/file
primitives as well as maths-helpers.

You can see the implementations beneath [evaluator/stdlib*](evaluator/),
and several of these things are documented in [examples/](examples/).


## 2.5 Function

`monkey` use `fn` as the definition of function. Apart from regular function using, `monkey` also includes high order function.

```
>>>let add = fn(a, b) { return a + b;};
>>> add(1,2)
3
>>>let addTwo = fn(a,b, f) { return 2 + f(a, b);};
>>>addTwo(1,2, add)
5
```

It is also possible to define a function without using `let` via the `function` keyword:

>>>function hello() { puts "Hello, world" ; };
>>> hello()
Hello, world


## 2.6 If-else statements

`monkey` supports if-else statements.
```
>>> let max = fn(a, b) { if (a > b) { return a;} else { return b; } };
>>> max(1, 2)
2
```

## 2.7 For-loop statements

`monkey` support for-loop statement.

```
>>> let sum = fn(x) { let i = 1; let sum = 0; for (i < x) { let sum = sum + i; let i = i+1; } return sum; };
>>> sum(100)
4950
```

## 2.8 Comments

`monkey` support two kinds of comments:

* Comments beginning with `//` last until the following newline.
* Comments between `/*` and `*/` may span multiple lines.


## 2.9 Postfix Operators

The `++` and `--` modifiers are permitted for integer-variables, for examle:

    let i = 0;
    for ( i <= 5 ) {
       i++;
    }
