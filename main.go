package main

import (
	"fmt"
	"os"

	"github.com/FalconXAsmit/gocraft/internal/commands"
)

func runCommand(command string, args []string) error {
	switch command {
	case "version":
		commands.Version()

	case "info":
		commands.Info()

	case "help", "--help", "-h":
		commands.Help()

	case "init":
		if len(args) != 1 {
			fmt.Println("Usage: gocraft init <project-name>")
			return nil
		}
		return commands.Init(args[0])

	case "env":
		commands.Env()

	case "port":
		if len(args) == 0 {
			commands.ListPorts()
			return nil
		}
		if args[0] == "--watch" {
			if len(args) != 1 {
				fmt.Println("Usage: gocraft port --watch")
				return nil
			}
			commands.WatchPorts()
			return nil
		}
		if len(args) != 1 {
			fmt.Println("Usage: gocraft port <port>")
			return nil
		}
		commands.Port(args[0])

	default:
		fmt.Println("Unknown Command:", command)
		fmt.Println()
		fmt.Println("Run 'gocraft help' for available commands.")
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		commands.Help()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	err := runCommand(command, args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
