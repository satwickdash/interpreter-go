package main

import (
	"fmt"
	"interpreter-go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is Zine. Please type in your commands.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
