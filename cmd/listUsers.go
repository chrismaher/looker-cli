package cmd

import (
	"github.com/chrismaher/looker-cli/client"
	"github.com/chrismaher/looker-cli/tabwriter"
	"github.com/spf13/cobra"
)

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "A list of Looker users",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New()
		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			email = "%"
		}
		fields := []string{"email", "is_disabled"}
		users := client.GetUsers(email, map[string][]string{"fields": fields})
		tabwriter.TabPrinter(users, fields)
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)

	listUsersCmd.Flags().StringP("email", "e", "", "The user's email")
}
