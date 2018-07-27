package cmd

import (
	"broadcastle.co/code/crm/manage"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Run:   manage.ContactRemove,
}

func init() {
	clientCmd.AddCommand(removeCmd)
}
