package main

import (
	"fmt"
	"jokes-app/cmd"
	"os"
)

func main() {
	command := cmd.NewCmdRoot()
	if err := command.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%v\n", err)
		if err != nil {
			fmt.Println("Error while printing to stderr: ", err.Error())
		}
		os.Exit(1)
	}
}
