package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strattonw/aug/config"
	"github.com/strattonw/aug/magazine"
	"os"
)

var fetchConfiguration config.Configuration

var fetchCommand = &cobra.Command{
	Use:   "fetch <log group>",
	Short: "Fetch log events",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("fetching events requires log group argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fetchConfiguration.Group = args[0]

		a, err := magazine.NewDefault()

		if err != nil {
			fmt.Println("Error", err)
			os.Exit(2)
		}

		if err = a.GetEvents(fetchConfiguration.FilterLogEventsInput()); err != nil {
			fmt.Println("Error", err)
			os.Exit(2)
		}
	},
}

func init() {
	fetchCommand.Flags().StringVar(
		&fetchConfiguration.Start,
		"start",
		"",
		`start getting the logs from this point
Takes an absolute timestamp in RFC3339 format, or parsable via time.ParseDuration.`,
	)
	fetchCommand.Flags().StringVar(
		&fetchConfiguration.End,
		"stop",
		"now",
		`stop getting the logs at this point
Takes an absolute timestamp in RFC3339 format, or parsable via time.ParseDuration.`,
	)
	fetchCommand.Flags().StringVar(&fetchConfiguration.Filter, "filter", "", "event filter pattern")
}
