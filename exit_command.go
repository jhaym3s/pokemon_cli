package main

import (
	"fmt"
	"os"
)

func ExitCommand() error{
	fmt.Println("Ending program")
	
	os.Exit(0)
	return nil
}