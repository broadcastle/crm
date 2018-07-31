package cmd

import (
	"broadcastle.co/code/crm/code/manage"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a contact.",
	Run:   manage.ContactRemove,
}

func init() {
	contactCmd.AddCommand(removeCmd)
}
