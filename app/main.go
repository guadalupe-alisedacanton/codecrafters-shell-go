package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"os/exec"
)

var builtins = []string{"exit", "echo", "type"}

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
				fmt.Println(argument + " is a shell builtin")
			} else {
				path, err := exec.LookPath(argument)
				if err == nil {
					fmt.Println(argument + " is " + path)
				} else {
				 fmt.Println(argument + ": not found")
				}
			}
			continue

			/*
			When type receives a command input, your shell must follow these steps:

			[DONE] Check if the command is a builtin command (like exit or echo). If it is, report it as a builtin (<command> is a shell builtin) and stop.
			
			If the command is not a builtin, your shell must go through every directory in PATH. For each directory:
				Check if a file with the command name exists.
				Check if the file has execute permissions.
				If the file exists and has execute permissions, print <command> is <full_path> and stop.
				If the file exists but lacks execute permissions, skip it and continue to the next directory.
				[DONE] If no executable is found in any directory, print <command>: not found.
			*/
		}

		fmt.Println(command + ": command not found")
	}

}
