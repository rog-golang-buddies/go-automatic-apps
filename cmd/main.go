package main

import (
	"fmt"
	"os"

	"github.com/rog-golang-buddies/go-automatic-apps/internal/parser"
)

func main() {
	commands, flags, err := parser.ParseCliArguments(os.Args[1:])

	if err != nil {
		panic(err)
	}

	fmt.Printf("CLI Commands: %v\n", commands)
	fmt.Printf("CLI Flags: %v\n", flags)
}
