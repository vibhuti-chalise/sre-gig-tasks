package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list jokes",
	Long:  "list all available jokes",
	Run: func(cmd *cobra.Command, args []string) {
		filecontent, err := os.ReadFile("jokes.txt")

		if err != nil {
			fmt.Println("Error reading file\n")
		}
		content := string(filecontent)
		fmt.Println(fmt.Sprintln(content))
	},
}
