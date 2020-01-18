package cmd

import (
	"fmt"

	"github.com/chrismaher/looker-cli/client"
	"github.com/spf13/cobra"
)

// listQueriesCmd represents the listQueries command
var listQueriesCmd = &cobra.Command{
	Use:   "list-queries",
	Short: "A brief description of your command",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New()
		queries := client.GetQueries()

		for _, q := range queries {
			fmt.Println(q)
		}
	},
}

func init() {
	rootCmd.AddCommand(listQueriesCmd)
}
