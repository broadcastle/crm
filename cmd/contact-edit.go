package cmd

import (
	"broadcastle.co/code/crm/code/manage"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing contact.",
	Run:   manage.ContactEdit,
}

func init() {
	contactCmd.AddCommand(editCmd)

	contactFlags(editCmd)
}
