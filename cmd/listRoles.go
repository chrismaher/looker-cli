package cmd

import (
	"fmt"

	"github.com/chrismaher/looker-cli/client"
	"github.com/spf13/cobra"
)

// listRolesCmd represents the listRoles command
var listRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: "A list of Looker roles",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.New()
		roles := client.GetRoles()

		for _, r := range roles {
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(listRolesCmd)
}
