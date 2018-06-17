package evaluator

import (
	"bufio"
	"monkey/object"
	"os"
)

//
// Mapping of file-IDs to file-handles
//
var file_handles = make(map[uintptr]*os.File)
var file_readers = make(map[uintptr]*bufio.Reader)

//
// Horrid hack - setup STDIN/STDOUT/STDERR
//
func setupHandles() {
	if file_handles[0] != nil {
		return
	}
	file_handles[0] = os.Stdin
	file_handles[1] = os.Stdout
	file_handles[2] = os.Stderr

	file_readers[0] = bufio.NewReader(os.Stdin)
	file_readers[1] = bufio.NewReader(os.Stdout)
	file_readers[2] = bufio.NewReader(os.Stderr)
}

// handle = file.open(path)
func fileOpen(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	path := args[0].(*object.String).Value
	file, err := os.Open(path)
	if err != nil {
		return &object.Integer{Value: -1}
	}

	// convert handle to integer to return it
	file_handles[file.Fd()] = file
	// but also store a reader
	file_readers[file.Fd()] = bufio.NewReader(file)

	return &object.Integer{Value: int64(file.Fd())}
}

// file.close(handle)
func fileClose(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	handle := args[0].(*object.Integer).Value

	file_handles[uintptr(handle)].Close()
	delete(file_handles, uintptr(handle))
	delete(file_readers, uintptr(handle))
	return NULL
}

// [] = file.lines("path")
func fileLines(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	path := args[0].(*object.String).Value
	file, err := os.Open(path)
	if err != nil {
		return NULL
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// make results
	l := len(lines)
	result := make([]object.Object, l, l)
	for i, txt := range lines {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

// str = read(handle)
func readInput(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	id := args[0].(*object.Integer).Value
	reader := file_readers[uintptr(id)]

	line, err := reader.ReadString('\n')
	if err == nil {
		return &object.String{Value: line}
	} else {
		return &object.String{Value: ""}
	}
}

func init() {
	registerBuiltin("read",
		func(args ...object.Object) object.Object {
			return (readInput(args...))
		})
	registerBuiltin("file.open",
		func(args ...object.Object) object.Object {
			return (fileOpen(args...))
		})
	registerBuiltin("file.close",
		func(args ...object.Object) object.Object {
			return (fileClose(args...))
		})
	registerBuiltin("file.lines",
		func(args ...object.Object) object.Object {
			return (fileLines(args...))
		})
}
