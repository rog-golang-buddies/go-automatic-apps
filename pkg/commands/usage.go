package commands

import "fmt"

func PrintUsageInstructions() {
	fmt.Println("usage: gaa [action] [-flags]")
	fmt.Println("For more options, run: gaa help")
}
