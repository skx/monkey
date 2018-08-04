[![Travis CI](https://img.shields.io/travis/skx/monkey/master.svg?style=flat-square)](https://travis-ci.org/skx/monkey)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/monkey)](https://goreportcard.com/report/github.com/skx/monkey)
[![license](https://img.shields.io/github/license/skx/monkey.svg)](https://github.com/skx/monkey/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/monkey.svg)](https://github.com/skx/monkey/releases/latest)

# Monkey

This repository contains an interpreter for the "Monkey" programming language, as described in [Write an Interpreter in Go](https://interpreterbook.com).

This repository started life as the implementation written by [gaufung](https://github.com/gaufung/Monkey) which extended the code in the book to add a `for` statement.

#### My changes

The interpreter in _this_ repository has been further extended:

* Added single-line & multi-line comments.
* Added postfix operators (`i++`, `i--`).
* Allow accessing individual characters of a string via the index-operator.
* Added a driver to read from STDIN, or a named file, rather than a REPL.
    * This allows executing the examples easily (for example "`./monkey examples/hello.mon`".)
* Added a collection of standard-library functions.
    * Including file input/output, type-discovery, string, and math functions.
* Added a new way to define functions, via `function`.
* Added the `<=` + `>=` comparison functions.
* Allow assignments without `let` (after initial declaration).
    * This allows operators such as "`+=`", "`-=`", "`*=`", & "`/=`" to work.
* Added command-line handling, so that scripts can read their own arguments.
* Added global-constants available by default
    * For example `PI`, `E`, `STDIN`, `STDOUT`, & `STDERR`.
* Most scripts will continue running in the face of errors.
    * To correct/detect "obvious" errors add `pragma("strict");` to your script, which will cause the interpreter to show a suitable error-message and terminate.


## 1. Installation

If you have a working [golang](https://golang.org/) setup you can install the interpreter via:

    $ go get -u  github.com/skx/monkey
    $ go install github.com/skx/monkey

Alternatively you could install a binary-release, from the [release page](https://github.com/skx/monkey/releases).

If you're an [emacs](https://www.gnu.org/software/emacs/) user might also wish to install the [monkey.el](emacs/monkey.el) file, which provides syntax highlighting for monkey-scripts.


### 1.1 Usage

To execute a monkey-script simply pass the name to the interpreter:

     $ monkey ./example/hello.mon

Scripts can be made executable by adding a suitable shebang line:

     $ cat hello.mon
     #!/usr/bin/env monkey
     puts( "Hello, world!\n" );

Execution then works as you would expect:

     $ chmod 755 hello.mon
     $ ./hello.mon
     Hello, world!

If no script-name is passed to the interpreter it will read from STDIN and
execute that instead.  This could be used like so:

     $ echo 'puts("Read from STDIN!\n");' | monkey


# 2 Syntax

**NOTE**: Example-programs can be found beneath [examples/](examples/) which
demonstrate these things, as well as parts of the standard-library.


## 2.1 Definitions

Variables are defined using the `let` keyword, with each line ending with `;`.

      let a = 3;
      let b = 1.2;

Variables may be integers, floats, strings, or arrays/hashes (which are discussed later).

Some variables are defined by default, for example:

    puts( PI ); // Outputs: 3.14159..
    puts( E );  // Outputs: 2.71828..

Once defined variables may be updated without the need for `let`, for example:

    let world = "Earth";
    world = "world";
    puts( "Hello, " + world + "!\n");



## 2.2 Arithmetic operations

`monkey` supports all the basic arithmetic operation of `int` and `float` types.

The `int` type is represented by `int64` and `float` type is represented by `float64`.


       let a = 3;
       let b = 1.2;

       puts( a + b  );  // Outputs: 4.2
       puts( a - b  );  // Outputs: 1.8
       puts( a * b  );  // Outputs: 3.6
       puts( a / b  );  // Outputs: 2.5
       puts( 2 ** 3 ) ; // Outputs: 8

Here `**` is used to raise the first number to the power of the second.
When operating with integers the modulus operator is available too, via `%`.


## 2.3 Builtin containers

`monkey` contains two builtin containers: `array` and `hash`.


### 2.3.1 Arrays

An array is a list which organizes items by linear sequence.  Arrays can hold multiple types.

     let a = [1, 2.3, "array"];
     let b = [false, true, "Hello World", 3, 3.13];


Adding to an array is done via the `push` function:

     let a = push(a, "another");

You can iterate over the contents of an array like so:

     let i = 0;
     for( i < len(a) ) {
        puts( "Array index ", i, " contains ", a[i], "n");
        i++
     }

With the definition we included that produces this output:

     Array index 0 contains 1
     Array index 1 contains 2.3
     Array index 2 contains array
     Array index 3 contains another


### 2.3.2 Hashes

A hash is a key/value container, but note that keys may only be of type `boolean`, `int` and `string`.


    let a = {"name":"monkey",
             true:1,
             7:"seven"};

    puts(a); // Outputs: {name: monkey, true: 1, 7: seven}

    puts(a["name"]); // Outputs: monkey

Updating a hash is done via the `set` function, but note that this returns
an updated hash - rather than changing in-place:

    let b = set(a, 8, "eight");
    puts( b);  // Outputs: {name: monkey, true: 1, 7: seven, 8: eight}

You can iterate over the keys in a hash via the `keys` function, or delete
keys via `delete` (again these functions returns an updated value rather than
changing it in-place).

Hash functions are demonstrated in the [examples/hash.mon](examples/hash.mon) sample.


## 2.4 Builtin functions

The core primitives are:

* `delete`
  * Deletes a hash-key.
* `first`
  * yield the first element of array.
* `int`
  * convert the given float/string to an integer.
* `keys`
  * Return the keys of the specified array.
* `len`
  * Yield the length of builtin containers.
* `last`
  * yield the last element of array.
* `match`
  * Regular-expression matching.
* `pragma`
  * Allow the run-time environment to be controlled.
  * We currently support only `pragma("strict");`.
* `push`
  * push an elements into the array.
* `puts`
  * print literal value of objects.
* `rest`
  * yield an array which excludes the first element.
* `set`
  * insert key value pair into the map.
* `string`
  * convert the given item to a string.
* `type`
  * returns the type of a variable.


## 2.4.1 The Standard Library

In addition to the core built-in functions we also have a minimal-standard library.  The library includes some string/file primitives as well as maths-helpers.

You can see the implementations beneath [evaluator/stdlib*](evaluator/),
and several of these things are documented in [examples/](examples/).


## 2.5 Functions

`monkey` use `fn` to define a function which will be assigned to a variable for
naming/invocation purposes:


    let add = fn(a, b) { return a + b;};
    puts(add(1,2));  // Outputs: 3

    // functions can be used via their variables
    let addTwo = fn(a,b, f) { return 2 + f(a, b);};
    puts( addTwo(1,2, add) ); // outputs: 5.

It is also possible to define a function without the use of `let`, via the `function` keyword.  This was added to make the language feel more natural to C-developers:

    function hello() { puts "Hello, world\n" ; };
    hello();   // Outputs: Hello, world" to the console.


## 2.6 If-else statements

`monkey` supports if-else statements.

    let max = fn(a, b) {
      if (a > b) {
        return a;
      } else {
        return b;
        }
    };

    puts( max(1, 2) );  // Outputs: 2


## 2.7 For-loop statements

`monkey` supports a golang-style for-loop statement.

     let sum = fn(x) {
        let i = 1;
        let sum = 0;

        for (i < x) {
           sum = sum + i;
           i++;
        }
        return sum;
     };

     puts(sum(100));  // Outputs: 4950


## 2.8 Comments

`monkey` support two kinds of comments:

* Single-line comments begin with `//` and last until the end of the line.
* Multiline comments between `/*` and `*/`.


## 2.9 Postfix Operators

The `++` and `--` modifiers are permitted for integer-variables, for example:

    let i = 0;
    for ( i <= 5 ) {
       puts( i, "\n" );
       i++;
    }

These postfix-operators update the contents of the named variable to increase/decrease it by one.  We also allow variables to be updated in-place via `+=`, `-=`, `*=`, & `/=`.

Using `+=` our previous example could be rewritten as:

    let i = 0;
    for ( i <= 5 ) {
       puts( i, "\n" );
       i += 1;
    }


## 2.10 Command Execution

As with many scripting languages commands may be executed via the backtick
operator (`\``).

      let uptime = `/usr/bin/uptime`;

      if ( uptime ) {
          puts( "STDOUT: ", string.trim(uptime["stdout"] ) , "\n");
          puts( "STDERR: ", string.trim(uptime["stderr"] ) , "\n");
      } else {
          puts( "Failed to run command\n");
      }

The output will be a hash with two keys `stdout` and `stderr`.  NULL is
returned if the execution fails.  This can be seen in [examples/exec.mon](examples/exec.mon).


## 2.11 Regular Expressions

The `match` function allows matching a string against a regular-expression.

If a match fails NULL will be returned, otherwise a hash containing any
capture groups in the match.

This is demonstrated in the [examples/regexp.mon](examples/regexp.mon) example.


Steve
--
