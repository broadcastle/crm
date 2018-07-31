package cmd

import (
	"broadcastle.co/code/crm/code/manage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact.",
	Run:   manage.ContactCreate,
}

func init() {
	contactCmd.AddCommand(addCmd)

	contactFlags(addCmd)
	addCmd.Flags().BoolP("fast", "f", false, "create a client entry only using a name and email")

}
