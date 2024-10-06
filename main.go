package main

import (
	"fmt"
	"moodeng/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Moodeng!\n", user.Username)
	fmt.Printf("You can tart typing commands...\n")
	repl.Start(os.Stdin, os.Stdout)
}
