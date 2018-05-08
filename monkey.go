package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command\n")
	fmt.Printf(`Enter "exit()" or CTRL+C to quit command interface`)
	fmt.Printf("\n")
	repl.Start(os.Stdin, os.Stdout)

}
