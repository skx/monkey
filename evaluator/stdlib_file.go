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
var fileHandles = make(map[uintptr]*os.File)
var fileReaders = make(map[uintptr]*bufio.Reader)
var fileWriters = make(map[uintptr]*bufio.Writer)

//
// Horrid hack - setup STDIN/STDOUT/STDERR
//
func setupHandles() {
	if fileHandles[0] != nil {
		return
	}
	fileHandles[0] = os.Stdin
	fileHandles[1] = os.Stdout
	fileHandles[2] = os.Stderr

	fileReaders[0] = bufio.NewReader(os.Stdin)
	fileReaders[1] = bufio.NewReader(os.Stdout)
	fileReaders[2] = bufio.NewReader(os.Stderr)

	fileWriters[0] = bufio.NewWriter(os.Stdin)
	fileWriters[1] = bufio.NewWriter(os.Stdout)
	fileWriters[2] = bufio.NewWriter(os.Stderr)
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
	fileHandles[file.Fd()] = file

	// but also store a reader / writer as appropriate
	if md == os.O_RDONLY {
		fileReaders[file.Fd()] = bufio.NewReader(file)
	} else {
		fileWriters[file.Fd()] = bufio.NewWriter(file)
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
	if fileWriters[uintptr(handle)] != nil {
		fileWriters[uintptr(handle)].Flush()
	}
	fileHandles[uintptr(handle)].Close()
	delete(fileHandles, uintptr(handle))
	delete(fileReaders, uintptr(handle))
	delete(fileWriters, uintptr(handle))
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
	reader := fileReaders[uintptr(id)]
	if reader == nil {
		return newError("Reading from an unopened file-handle.")
	}

	line, err := reader.ReadString('\n')
	if err == nil {
		return &object.String{Value: line}
	}
	return &object.String{Value: ""}
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

	writer := fileWriters[uintptr(id)]
	if writer == nil {
		return newError("Writing to an unopened file-handle.")
	}

	_, err := writer.Write([]byte(txt))
	if err == nil {
		writer.Flush()
		return &object.Boolean{Value: true}
	}

	return &object.Boolean{Value: false}
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
