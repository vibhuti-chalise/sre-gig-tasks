package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jokes-app",
	Short: "command line jokes-app ",
	Long:  `This is a jokes-app that copies jokes from an API and allows to get jokes from a remote Jokes API; add jokes to a local file; and list jokes from the save local file`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
