package evaluator

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/skx/monkey/object"
)

//
// Mapping of file-IDs to file-handles
//
var file_handles = make(map[uintptr]*os.File)
var file_readers = make(map[uintptr]*bufio.Reader)
var file_writers = make(map[uintptr]*bufio.Writer)

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

	file_writers[0] = bufio.NewWriter(os.Stdin)
	file_writers[1] = bufio.NewWriter(os.Stdout)
	file_writers[2] = bufio.NewWriter(os.Stderr)
}

// array = directory.glob( "/etc/*.conf" )
func dirGlob(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	pattern := args[0].(*object.String).Value

	entries, err := filepath.Glob(pattern)
	if err != nil {
		return NULL
	}

	// Create an array to hold the results and populate it
	l := len(entries)
	result := make([]object.Object, l, l)
	for i, txt := range entries {
		result[i] = &object.String{Value: txt}
	}
	return &object.Array{Elements: result}
}

// handle = file.open(path)
func fileOpen(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 1 && len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=1|2",
			len(args))
	}

	path := args[0].(*object.String).Value

	md := os.O_RDONLY
	if len(args) == 2 {
		mode := args[1].(*object.String).Value
		if mode == "w" {
			md = os.O_WRONLY
		}

	}
	file, err := os.OpenFile(path, os.O_CREATE|md, 0644)
	if err != nil {
		return &object.Integer{Value: -1}
	}

	// convert handle to integer to return it
	file_handles[file.Fd()] = file

	// but also store a reader / writer as appropriate
	if md == os.O_RDONLY {
		file_readers[file.Fd()] = bufio.NewReader(file)
	} else {
		file_writers[file.Fd()] = bufio.NewWriter(file)
	}

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

	// If the file was opened for writing then we must flush
	// it before we close the handle - otherwise our written
	// data might be lost.
	if file_writers[uintptr(handle)] != nil {
		file_writers[uintptr(handle)].Flush()
	}
	file_handles[uintptr(handle)].Close()
	delete(file_handles, uintptr(handle))
	delete(file_readers, uintptr(handle))
	delete(file_writers, uintptr(handle))
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
	if reader == nil {
		return newError("Reading from an unopened file-handle.")
	}

	line, err := reader.ReadString('\n')
	if err == nil {
		return &object.String{Value: line}
	} else {
		return &object.String{Value: ""}
	}
}

// write(handle, text)
func writeOutput(args ...object.Object) object.Object {
	setupHandles()
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2",
			len(args))
	}

	id := args[0].(*object.Integer).Value
	txt := args[1].(*object.String).Value

	writer := file_writers[uintptr(id)]
	if writer == nil {
		return newError("Writing to an unopened file-handle.")
	}

	_, err := writer.Write([]byte(txt))
	if err == nil {
		writer.Flush()
		return &object.Boolean{Value: true}
	} else {
		return &object.Boolean{Value: false}
	}
}

func init() {
	RegisterBuiltin("directory.glob",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (dirGlob(args...))
		})
	RegisterBuiltin("read",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (readInput(args...))
		})
	RegisterBuiltin("write",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (writeOutput(args...))
		})
	RegisterBuiltin("file.open",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (fileOpen(args...))
		})
	RegisterBuiltin("file.close",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (fileClose(args...))
		})
	RegisterBuiltin("file.lines",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (fileLines(args...))
		})
}
