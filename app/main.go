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
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		} else if strings.HasPrefix(command, "type ") {
			argument := command[5:]
			if slices.Contains(builtins, argument) {
				fmt.Println(argument + " is a shell builtin")
			} else {
				path := findExecutablePath(argument)
				if path != "" {
					fmt.Println(argument + " is " + path)
				} else {
					fmt.Println(argument + ": not found")
				}
			}
			continue
		} else if command == "pwd" {

		}
		arguments := strings.Fields(command)
		customCommand := arguments[0]
		path := findExecutablePath(customCommand)
		if path != "" {
			cmd := exec.Command(customCommand, arguments[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			continue
		} else {
			fmt.Println(command + ": command not found")
		}
	}

}
