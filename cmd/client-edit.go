package cmd

import (
	"broadcastle.co/code/crm/manage"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Run:   manage.ContactEdit,
}

func init() {
	clientCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("name", "n", "", "change the contact name")
	editCmd.Flags().StringP("email", "e", "", "change the contact email")
	editCmd.Flags().String("phone", "", "add / change the contact phone")
}
