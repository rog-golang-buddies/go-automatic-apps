package commands

import (
	"fmt"
	"runtime"
)

var VersionCommand = BaseCommand{
	Name:        "version",
	Description: "Provides useful information about go environment",
	Run:         versionCommand,
}

func versionCommand(commands *[]string, flags *map[string]interface{}) error {
	fmt.Printf("Go Version: %v\n", runtime.Version())
	fmt.Printf("GAA: %v\n", runtime.GOARCH)
	fmt.Printf("GOOS: %v\n", runtime.GOOS)
	return nil
}
