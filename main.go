package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", currentUser.Username)
	fmt.Printf("Feel free t type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
