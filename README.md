# Monkey
A interpreter language implementation in Go 

# 1 Introduction
[Write an Interpreter in Go](https://interpreterbook.com) source code. 

[中文翻译](book/README.md)

**Monkey** interpreter language which is implemented Go language. After typing `monkey` in terminal, you get into the `monkey` programming language.

```
$ ./monkey
Hello $Username! This is Monkey programming language!
Feel free to type in commnd
Enter "exit()" or CTRL+C to quit command interface
>>>
``` 

# 2 Syntax

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
>>> let a = {"name":"moneky", true:1, 7:"sevent"};
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

insert key value pair into the map

- `puts`

print literal value of objects.


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

## 2.5 If-else statements

`monkey` supports if-else statements.
```
>>> let max = fn(a, b) { if (a > b) { return a;} else { return b; } };
>>> max(1, 2)
2
```

## 2.6 For-loop statements

`monkey` support for-loop statement.

```
>>> let sum = fn(x) { let i = 1; let sum = 0; for (i < x) { let sum = sum + i; let i = i+1; } return sum; };
>>> sum(100)
4950
```


# 3 Extensions
To make `monkey` interpreter language to be more powerfull, it is deserved improving it.

- [x] float type
- [x] unicode literal
- [x] loop branch
- [ ] translate into Chinese
- [ ] ... 