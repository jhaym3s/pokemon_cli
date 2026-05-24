package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command: ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := CleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		availableCommands := getCommands()
		selectedCommand, ok := availableCommands[command]
		if !ok {
			fmt.Println("Enter a proper command")
			continue
		}
		selectedCommand.callBack()

	}

}

func CleanInput(input string) []string {
	lowerCased := strings.ToLower(input)
	words := strings.Fields(lowerCased)
	return words
}

type cliCommands struct {
	name        string
	description string
	callBack    func() error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Display the commands",
			callBack:    HelpCallBack,
		},
		"exit": {
			name:        "Exit",
			description: "Exit loop",
			callBack:    ExitCommand,
		},

		"map": {
			name:        "map",
			description: "Get location from pokeapi",
			callBack:    LocationCallBack,
		},
	}
}
