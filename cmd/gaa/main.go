package main

import (
	"fmt"
	"os"

	"github.com/rog-golang-buddies/go-automatic-apps/internal/parser"
	cmd "github.com/rog-golang-buddies/go-automatic-apps/pkg/commands"
)

func main() {
	commands, flags, err := parser.ParseCliArguments(os.Args[1:])

	if err != nil {
		cmd.PrintUsageInstructions()
		return
	}

	baseCommand := commands[0]
	subCommands := commands[1:]

	switch baseCommand {
	case "help":
		err = cmd.HelpCommand.Run(&subCommands, &flags)
	case "version":
		err = cmd.VersionCommand.Run(&subCommands, &flags)
	default:
		fmt.Println("unrecognised action")
		cmd.PrintUsageInstructions()
	}

	if err != nil {
		panic(fmt.Errorf("error executing command %v", err))
	}
}
