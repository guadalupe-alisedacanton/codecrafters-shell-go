package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var builtins = [3]string{"exit", "echo", "type"}

func main() {
	for {
		fmt.Print("$ ")

		// Captures the user's command in the "command" variable
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = command[:len(command)-1]
		if command == "exit" {
			break
		}
		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}
		if strings.HasPrefix(command, "type ") {
			argument := command[5:]
			if slices.Contains(builtins, argument) {
				fmt.Println(command + " is a shell builtin")
				continue
			}
		}

		fmt.Println(command + ": command not found")
	}

}
