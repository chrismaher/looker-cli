package cmd

import (
	"github.com/chrismaher/looker-cli/client"
	"github.com/chrismaher/looker-cli/tabwriter"
	"github.com/spf13/cobra"
	"log"
)

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "A list of Looker users",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.New()
		if err != nil {
			log.Panic(err)
		}

		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			email = "%"
		}
		fields := []string{"email", "is_disabled"}
		users, err := client.GetUsers(email, map[string][]string{"fields": fields})
		if err != nil {
			log.Panic(err)
		}
		tabwriter.TabPrinter(users, fields)
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)

	listUsersCmd.Flags().StringP("email", "e", "", "The user's email")
}
