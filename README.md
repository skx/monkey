[![Go Report Card](https://goreportcard.com/badge/github.com/skx/monkey)](https://goreportcard.com/report/github.com/skx/monkey)
[![license](https://img.shields.io/github/license/skx/monkey.svg)](https://github.com/skx/monkey/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/monkey.svg)](https://github.com/skx/monkey/releases/latest)



* [Monkey](#monkey)
  * [My changes](#my-changes)
  * [See also](#see-also)
* [1. Installation](#1-installation)
  * [Binary Releases](#binary-releases)
* [1.1 Usage](#11-usage)
* [2 Syntax](#2-syntax)
  * [2.1 Definitions](#21-definitions)
  * [2.2 Arithmetic operations](#22-arithmetic-operations)
  * [2.3 Builtin containers](#23-builtin-containers)
    * [2.3.1 Arrays](#231-arrays)
    * [2.3.2 Hashes](#232-hashes)
  * [2.4 Builtin functions](#24-builtin-functions)
    * [2.4.1 The Standard Library](#241-the-standard-library)
  * [2.5 Functions](#25-functions)
  * [2.6 If-else statements](#26-if-else-statements)
    * [2.6.1 Ternary expressions](#261-ternary-expressions)
  * [2.7 Switch statements](#27-switch-statements)
  * [2.8 For-loop statements](#28-for-loop-statements)
    * [2.8.1 Foreach statements](#281-foreach-statements)
  * [2.9 Comments](#29-comments)
  * [2.10 Postfix Operators](#29-postfix-operators)
  * [2.11 Command Execution](#211-command-execution)
  * [2.12 Regular Expressions](#212-regular-expressions)
  * [2.13 File I/O](#213-file-io)
  * [2.14 File Operations](#214-file-operations)
  * [3. Object Methods](#3-object-methods)
      * [3.1 Defininig New Object Methods](#31-defininig-new-object-methods)
* [Github Setup](#github-setup)
* [Fuzz Testing](#fuzz-testing)




# Monkey

This repository contains an interpreter for the "Monkey" programming language, as described in [Write an Interpreter in Go](https://interpreterbook.com).

#### My changes

The interpreter in _this_ repository has been significantly extended from the starting point:

* Added single-line & multi-line comments.
* Added postfix operators (`i++`, `i--`).
* Allow accessing individual characters of a string via the index-operator.
* Added a driver to read from STDIN, or a named file, rather than a REPL.
    * This allows executing the examples easily (for example "`./monkey examples/hello.mon`".)
* Added a collection of standard-library functions.
    * Including file input/output, type-discovery, string, and math functions.
* Added a new way to define functions, via `function`.
* Added the general-purpose comparision functions `<=` & `>=`.
* Allow string comparisons via `==`, `!=`, `<=`, & `>=`.
* Allow comparisions to be complex:
  * `if ( a >= 'a' && a <= 'z' ) ..`
  * `if ( a || b ) ..`
* Allow assignments without `let`.
    * This also allows operators such as "`+=`", "`-=`", "`*=`", & "`/=`" to work.
* Added command-line handling, so that scripts can read their own arguments.
* Added global-constants available by default
    * For example `PI`, `E`, `STDIN`, `STDOUT`, & `STDERR`.
* Most scripts will continue running in the face of errors.
    * To correct/detect "obvious" errors add `pragma("strict");` to your script, which will cause the interpreter to show a suitable error-message and terminate.
* Function arguments may have defaults.  For example:
  * `function greet( name = "World" ) { puts("Hello, " + name + "\n"); }`
* Moved parts of the standard-library to 100% pure monkey, rather than implementing it in go.
  * See [data/stdlib.mon](data/stdlib.mon) for the implementation.
  * See also the notes on [object-based methods](#31-defininig-new-object-methods).
* Added the `eval` function.
  * Which allows executing monkey-code from a string.
* Improved error-reporting from the parser.
  * It will now show the line-number of failures (where possible).
* Added support for regular expressions, both literally and via `match`
  * `if ( name ~= /steve/i ) { puts( "Hello Steve\n"); } `
* Added support for [ternary expressions](#261-ternary-expressions).
* Added support for creating arrays of consecutive integers via the range operator (`1..10`).
* Added the ability to iterate over the contents of arrays, hashes, and strings via the `foreach` statement.
* Added `printf` and `sprintf` primitives, which work as you would expect.
  * `printf( "%d %s", 3, "Steve" );`
* Added support for `switch` statements, with block-based `case` expressions.
  * No bugs due to C-style "fall-through".
* Add support for explicit `null` usage:
  * `a = null;  if ( a == null ) { .. }`

#### See Also

If you enjoyed this repository you might find the related ones interesting:

* A tutorial-lead approach to implementing a FORTH interpreter:
  * https://github.com/skx/foth
* A simple TCL-like interpreter:
  * https://github.com/skx/critical
* A BASIC interpreter:
  * https://github.com/skx/gobasic
* An embedded scripting language, based upon the same Monkey core
  * This follows the second book, but large parts of the code were replaced with different implementations, and things were extended a lot.
  * https://github.com/skx/evalfilter

Finally I put together a couple of "complex" compilers, which convert input into AMD64 assembly language:

* A mathematical compiler
  * https://github.com/skx/math-compiler
* A brainfuck compiler:
  * https://github.com/skx/bfcc



## 1. Installation

Due to the embedded [standard-library implementation](data/stdlib.mon), which is implemented in monkey, you'll need to compile this project with go version 1.16beta1 or higher.

You can install from source like so:

    git clone https://github.com/skx/monkey
    cd monkey
    go install


### Binary Releases

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
execute that instead, allowing simple tests to be made.




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

Variables may be updated without the need for `let`, for example this works
as you would expect:

    let world = "Earth";
    world = "world";
    puts( "Hello, " + world + "!\n");

If you're __not__ running with `pragma("strict");` you can also declare and
use variables without the need for `let`, but that should be avoided as
typos will cause much confusion!

     name = "Steve";
     puts( "Hello, " + name + "\n");



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
        puts( "Array index ", i, " contains ", a[i], "\n");
        i++
     }

With the definition we included that produces this output:

     Array index 0 contains 1
     Array index 1 contains 2.3
     Array index 2 contains array
     Array index 3 contains another

As a helper you may define an array of consecutive integers via the range operator (`..`):

     let a = 1..10;


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
    puts(b);  // Outputs: {name: monkey, true: 1, 7: seven, 8: eight}

You can iterate over the keys in a hash via the `keys` function, or delete
keys via `delete` (again these functions returns an updated value rather than
changing it in-place).

Hash functions are demonstrated in the [examples/hash.mon](examples/hash.mon) sample.



## 2.4 Builtin functions

The core primitives are:

* `delete`
  * Deletes a hash-key.
* `int`
  * convert the given float/string to an integer.
* `keys`
  * Return the keys of the specified array.
* `len`
  * Yield the length of builtin containers.
* `match`
  * Regular-expression matching.
* `pragma`
  * Allow the run-time environment to be controlled.
  * We currently support only `pragma("strict");`.
* `push`
  * push an elements into the array.
* `puts`
  * Write literal value of objects to STDOUT.
* `printf`
  * Write values to STDOUT, via a format-string.
* `set`
  * insert key value pair into the map.
* `sprintf`
  * Create strings, via a format-string.
* `string`
  * convert the given item to a string.
* `type`
  * returns the type of a variable.

The following functions are also part of our standard library, but are
implemented in 100% pure monkey:

* `first`
  * yield the first element of array.
* `last`
  * yield the last element of array.
* `rest`
  * yield an array which excludes the first element.



## 2.4.1 The Standard Library

In addition to the core built-in functions we also have a minimal-standard library.  The library includes some string/file primitives, a regular-expression matcher, and some maths-helpers.

You can see the implementation of the go-based standard-library beneath [evaluator/stdlib*](evaluator/), and several of these functions are documented in the various [examples/](examples/).

**NOTE**: Parts of our standard-library are implemented in 100% pure monkey,
and these are embedded in our compiled interpreter.  The source of the functions
can be viewed in [data/stdlib.mon](data/stdlib.mon).

If you wish to make changes to the monkey-based standard-library you'll
need to rebuild the interpreter after making your changes, to ensure they are bundled into the executable.

Nothing special is required, the following will suffice as you'd expect:

    go build .



## 2.5 Functions

`monkey` uses `fn` to define a function which will be assigned to a variable for
naming/invocation purposes:


    let add = fn(a, b) { return a + b;};
    puts(add(1,2));  // Outputs: 3

    // functions can be used via their variables
    let addTwo = fn(a,b, f) { return 2 + f(a, b);};
    puts( addTwo(1,2, add) ); // outputs: 5.

It is also possible to define a function without the use of `let`, via the `function` keyword.  This was added to make the language feel more natural to C-developers:

    function hello() { puts "Hello, world\n" ; };
    hello();   // Outputs: Hello, world" to the console.

You may specify a default value for arguments which are not provided, for example:

    let foo = fn( name = "World!") {
      puts( "Hello, " + name + "\n" );
    };

    foo();
    foo( "Steve" );

This will output what you expect:

    Hello, World!
    Hello, Steve

The same thing works for literal functions:

    // Function with a default (string) argument
    function meh( arg = "Steve" ) {
      puts( "Argument:", arg, " has type:", type(arg), "\n");
    };

    // Call it with no argument and the default will be used.
    meh();

    // But of course all the rest work just fine.
    meh( 1 );
    meh( 1/3.0 );
    meh( "Steve" );
    meh( [1,2,3,4] );
    meh( {"Steve":"Kemp", true:1, false:0, 7:"seven"} );



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


### 2.6.1 Ternary Expressions

`monkey` supports the use of ternary expressions, which work as you
would expect with a C-background:

    function max(a,b) {
      return( a > b ? a : b );
    };

    puts( "max(1,2) -> ", max(1, 2), "\n" );
    puts( "max(-1,-2) -> ", max(-1, -2), "\n" );

Note that in the interests of clarity nested ternary-expressions are illegal!



## 2.7 Switch Statements

Monkey supports the `switch` and `case` expressions, as the following example demonstrates:

```
  name = "Steve";

  switch( name ) {
    case /^steve$/i {
       printf("Hello Steve - we matched you via a regexp\n");
    }
    case "St" + "even" {
       printf("Hello SteveN, you were matched via an expression\n" );
    }
    case 3 {
       printf("Hello number three, we matched you literally.\n");
    }
    default {
       printf("Default case: %s\n", string(name) );
    }
  }
```

See also [examples/switch.mon](examples/switch.mon).



## 2.8 For-loop statements

`monkey` supports a golang-style for-loop statement.

     let sum = fn(x) {
        let i = 1;
        let sum = 0;

        for (i < x) {
           sum += i;
           i++;
        }
        return sum;
     };

     puts(sum(100));  // Outputs: 4950



## 2.8.1 Foreach statements

In addition to iterating over items with the `for` statement, as shown above, it is also possible to iterate over various items via the `foreach` statement.

For example to iterate over an array:

     a = [ "My", "name", "is", "Steve" ]
     foreach item in a {
          puts( "\t",  item , "\n");
     }

Here you see that we've iterated over the items of the array, we can also see their offsets like so:

     foreach offset, item in a {
          puts( offset, "\t",  item , "\n");
     }

The same style of iteration works for Arrays, Hashes, and the characters which make up a string.  You can see examples of this support in [examples/iteration.mon](examples/iteration.mon).

When iterating over hashes you can receive either the keys, or the keys and value at each step in the iteration, otherwise you receive the value and an optional index.



## 2.9 Comments

`monkey` support two kinds of comments:

* Single-line comments begin with `//` and last until the end of the line.
* Multiline comments between `/*` and `*/`.



## 2.10 Postfix Operators

The `++` and `--` modifiers are permitted for integer-variables, for example the following works as you would expect showing the numbers from `0` to `5`:

    let i = 0;
    for ( i <= 5 ) {
       puts( i, "\n" );
       i++;
    }

Another feature borrowed from C allows variables to be updated in-place via the operators `+=`, `-=`, `*=`, & `/=`.

Using `+=` our previous example could be rewritten as:

    let i = 0;
    for ( i <= 5 ) {
       puts( i, "\n" );
       i += 1;
    }

The update-operators work with integers and doubles by default, when it comes to strings the only operator supported is `+=`, allowing for a string-append:

    let str = "Forename";
    str += " Surname";
    str += "\n";
    puts( str );           // -> "Forename Surname\n"



## 2.11 Command Execution

As with many scripting languages commands may be executed via the backtick
operator (``).

      let uptime = `/usr/bin/uptime`;

      if ( uptime["exitCode"] == 0 ) {
          puts( "STDOUT: ", uptime["stdout"].trim() , "\n");
      } else {
          puts( "An error occurred while running the command: ", uptime["stderr"].trim(), "\n");
      }

The output will be a hash containing the keys `stdout`, `stderr`, and `exitCode`, as demonstrated in [examples/exec.mon](examples/exec.mon).



## 2.12 Regular Expressions

The `match` function allows matching a string against a regular-expression.

If a match fails NULL will be returned, otherwise a hash containing any
capture groups in the match.

This is demonstrated in the [examples/regexp.mon](examples/regexp.mon) example.

You can also perform matching (complete with captures), with a literal regular expression object:

    if ( Name ~= /steve/i ) { puts( "Hello Steve\n" ); }
    if ( Name !~ /[aeiou]/i ) { puts( "You have no vowels.\n" ); }

    // captures become $1, $2, $N, etc.
    ip = "192.168.1.1";
    if ( ip ~= /([0-9]+)\.([0-9]+)\.([0-9]+)\.([0-9]+)/  ) {
        printf("Matched! %s.%s.%s.%s\n", $1, $2, $3, $4 );
    }



## 2.13 File I/O

The `open` primitive is used to open files, and can be used to open files for either reading, or writing:

    // Open a file for reading
    fh = open( "/etc/passwd" );
    fh = open( "/etc/passwd", "r" );

    // Open a file for writing
    fh = open( "/tmp/blah", "w" );

    // Open a file for appending
    fh = open( "/tmp/blah", "wa" );

Once you have a file-object you can invoke methods upon it:

* `read()`
  * Read a line of input, returning that input as a string.
* `readlines()`
  * Read the lines of the given file, and return them as an array.
* `write(data)`
  * Write the data to the given file.

These are demonstrated in the following examples:

* [examples/file.mon](examples/file.mon)
  * Simple example.
* [examples/file-writing.mon](examples/file-writing.mon)
  * Simple example.
* [examples/wc.mon](examples/wc.mon)
* [examples/wc2.mon](examples/wc2.mon)
  * Counting lines.

By default three filehandles will be made available, as constants:

* `STDIN`
  * Use for reading STDIN.
* `STDOUT`
* `STDERR`
  * Used for writing messages.



## 2.14 File Operations

The primitive `stat` will return a hash of details about the given file, or
directory entry.

You can change the permissions of a file via the `chmod` function, but note that the second argument is an __octal__ string:

    chmod( "/tmp/evil.sh", "755")
    chmod( "/tmp/normal", "644")

To remove a file, use `unlink`:

    unlink( "/tmp/trash.txt" )

And finally to make a directory:

    mkdir( "/tmp/blah" );




# 3. Object Methods

There is now support for "object-methods".  Object methods are methods
which are defined against a _type_.  For example all of our primitive
types allow a `methods()` method, which returns the methods which are
available against them.

Similarly each of them implement a `type()` function which returns the
type involved:

    let i = 1;
    puts( i.type() );

    let s = "Steve";
    puts( s.type() );

Or even:

    puts( "Steve".type() );

Seeing methods available works as you would expect:

    a = [ "Array", "Is", "Here" ];

    let i = 0;
    for ( i < len(a.methods() ) ) {
       puts( "Method " + a.methods()[i] + "\n" );
       i++;
    }

This shows:

    Method find
    Method len
    Method methods
    Method string

The `string` object has the most methods at the time of writing, but
no doubt things will change over time.



## 3.1 Defininig New Object Methods

The object-methods mentioned above are implemented in Go, however it is also
possible to define such methods in 100% monkey!

You can define a method via something like:

    function string.steve() {
       puts( "Hello, I received '", self, "' as an argument\n" );
    }

Note that the function has access to the object it was invoked upon via the
implicit `self` name.  Invocation would look as you expect:

    let s = "Hello, world";
    s.steve();   -> Hello, I received 'Hello, world' as an argument

You can see [data/stdlib.mon](data/stdlib.mon) implements some primitives
in this fashion, for example the functional-programming methods `array.map`,
`array.filter`, `string.toupper`, etc, etc.



## Github Setup

This repository is configured to run tests upon every commit, and when
pull-requests are created/updated.  The testing is carried out via
[.github/run-tests.sh](.github/run-tests.sh) which is used by the
[github-action-tester](https://github.com/skx/github-action-tester) action.

Releases are automated in a similar fashion via [.github/build](.github/build),
and the [github-action-publish-binaries](https://github.com/skx/github-action-publish-binaries) action.



## Fuzz Testing

Fuzz-testing involves creating random input, and running the program to test with that, to see what happens.

The intention is that most of the random inputs will be invalid, so you'll be able to test your error-handling and see where you failed to consider specific input things.

The 1.18 release of the golang compiler/toolset has integrated support for fuzz-testing, and you can launch it like so:

```sh
go test -fuzztime=300s -parallel=1 -fuzz=FuzzMonkey -v
```

Sample output looks like this:

```
$ go test -fuzztime=300s -parallel=1 -fuzz=FuzzMonkey -v
=== RUN   FuzzMonkey
fuzz: elapsed: 0s, gathering baseline coverage: 0/240 completed
fuzz: elapsed: 0s, gathering baseline coverage: 240/240 completed, now fuzzing with 1 workers
fuzz: elapsed: 3s, execs: 4321 (1440/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 6s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
cfuzz: elapsed: 9s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 12s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 15s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 18s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 21s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 24s, execs: 4321 (0/sec), new interesting: 6 (total: 246)
fuzz: elapsed: 27s, execs: 73463 (23060/sec), new interesting: 17 (total: 257)
fuzz: elapsed: 30s, execs: 75639 (725/sec), new interesting: 18 (total: 258)
fuzz: elapsed: 33s, execs: 125712 (16701/sec), new interesting: 25 (total: 265)
fuzz: elapsed: 36s, execs: 139338 (4543/sec), new interesting: 34 (total: 274)
fuzz: elapsed: 39s, execs: 173881 (11511/sec), new interesting: 49 (total: 289)
fuzz: elapsed: 42s, execs: 198046 (8055/sec), new interesting: 55 (total: 295)
fuzz: elapsed: 45s, execs: 210203 (4054/sec), new interesting: 75 (total: 315)
fuzz: elapsed: 48s, execs: 262945 (17580/sec), new interesting: 85 (total: 325)
fuzz: elapsed: 51s, execs: 297505 (11517/sec), new interesting: 108 (total: 348)
fuzz: elapsed: 54s, execs: 308672 (3722/sec), new interesting: 116 (total: 356)
fuzz: elapsed: 57s, execs: 341614 (10984/sec), new interesting: 123 (total: 363)
fuzz: elapsed: 1m0s, execs: 366053 (8146/sec), new interesting: 131 (total: 371)
fuzz: elapsed: 1m3s, execs: 396575 (10172/sec), new interesting: 137 (total: 377
...
```

Steve
--
