package main

import "fmt"

func HelpCallBack() error {
	fmt.Println("Welcome to pokemon CLI")
	fmt.Println("Available Commands")
	
	availableCommands := getCommands()

	for _, cmd := range availableCommands{
		fmt.Printf("- %s : %v \n", cmd.name, cmd.description)
	}
	return nil
}