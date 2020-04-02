package cmd

import (
	"fmt"
	"log"

	"github.com/chrismaher/looker-cli/client"
	"github.com/spf13/cobra"
)

// listRolesCmd represents the listRoles command
var listRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: "A list of Looker roles",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.New()
		if err != nil {
			log.Panic(err)
		}

		roles, err := client.GetRoles()
		if err != nil {
			log.Panic(err)
		}

		for _, r := range roles {
			fmt.Println(r)
		}
	},
}

func init() {
	rootCmd.AddCommand(listRolesCmd)
}
