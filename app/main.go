package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{"exit", "echo", "type"}

func findExecutablePath(argument string) string {
	path, err := exec.LookPath(argument)
	if err == nil {
		return path
	} else {
		return ""
	}
}
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
				// path, err := exec.LookPath(argument)
				// if err == nil {
				// 	fmt.Println(argument + " is " + path)
				// } else {
				//  fmt.Println(argument + ": not found")
				// }
				path := findExecutablePath(argument)
				if path != "" {
					fmt.Println(argument + " is " + path)
				} else {
					fmt.Println(argument + ": not found")
				}
			}
			continue
		}
		arguments := strings.Fields(command)
		customCommand := arguments[0]
		path := findExecutablePath(customCommand)
		if path != "" {
			cmd := exec.Command(customCommand, arguments[1:]...)
			cmd.Run()
			continue
		}

		fmt.Println(command + ": command not found")
	}

}
