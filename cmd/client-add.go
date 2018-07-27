package cmd

import (
	"broadcastle.co/code/crm/manage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new client.",
	Run:   manage.ContactCreate,
}

func init() {
	clientCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Name for this client.")
	addCmd.Flags().StringP("email", "e", "", "Email for this client.")
	addCmd.Flags().String("phone", "", "Phone number for this client")
	addCmd.Flags().BoolP("fast", "f", false, "Create a client entry only using a name and email.")
}
