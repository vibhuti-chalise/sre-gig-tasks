package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add jokes",
	Long:  "add jokes & store in a file",
	Run: func(cmd *cobra.Command, args []string) {

		joketext := strings.Join(args, " ")

		// checks if args is empty and reports a null joke value
		if len(args) >= 1 {
			fmt.Printf("Adding joke \"%s\" \n ", joketext)

			file, err := os.OpenFile("jokes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
			if err != nil {
				log.Printf("unable to open the file")
			}

			defer file.Close()

			if _, err := file.WriteString(joketext + "\n"); err != nil {
				log.Println("Unable to add joke to the file")
			}
		} else {
			fmt.Println("\nCannot add a null joke. Provide the joke as arguments to jokes-app add subcommand\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
