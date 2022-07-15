package commands

import "fmt"

var HelpCommand = BaseCommand{
	Name:        "help",
	Description: "Provides list of all commands",
	Run:         helpCommand,
}

func helpCommand(commands *[]string, flags *map[string]interface{}) error {
	// We have to "remake" the help command otherwise we get a runtime cycle error
	commandsList := []BaseCommand{{Name: "help", Description: "Provides list of all commands"}, VersionCommand}

	for _, command := range commandsList {
		fmt.Printf("[%v]: %v\n", command.Name, command.Description)
	}

	return nil
}
