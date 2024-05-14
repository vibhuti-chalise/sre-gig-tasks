package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "jokes-app",
		Short: "command line jokes-app ",
		Long:  `This is a jokes-app that copies jokes from an API and allows to get jokes from a remote Jokes API; add jokes to a local file; and list jokes from the save local file`,
	}
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(addCmd)

	return rootCmd
}
