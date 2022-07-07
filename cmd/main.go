package main

import (
	"fmt"
	"os"

	"github.com/rog-golang-buddies/internal/utils"
)

func main() {
	commands, flags, err := utils.ParseCliArguments(os.Args[1:])

	if err != nil {
		panic(err)
	}

	fmt.Printf("CLI Commands: %v\n", commands)
	fmt.Printf("CLI Flags: %v\n", flags)
}
