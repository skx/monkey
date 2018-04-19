package main

import (
	"os/user"
	"fmt"
	"monkey/repl"
	"os"
)

func main(){
	user, err := user.Current()
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.Start(os.Stdin, os.Stdout)

}
