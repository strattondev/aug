package cmd

import "github.com/spf13/cobra"

var AugCommand = &cobra.Command{
	Use:     "magazine <command>",
	Short:   "",
	Long:    "",
	Example: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	AugCommand.AddCommand(fetchCommand)
}
