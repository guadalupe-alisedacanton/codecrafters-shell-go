package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	// Captures the user's command in the "command" variable
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	fmt.Println(command[:len(command)-1] + ": command not found")

}
