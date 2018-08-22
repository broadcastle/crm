package cmd

import (
	"broadcastle.co/code/crm/code/manage"
	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Manage your contacts.",
}

var contactAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact.",
	Args:  cobra.NoArgs,
	Run:   manage.ContactCreate,
}

var contactEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing contact.",
	Args:  cobra.ExactArgs(1),
	Run:   manage.ContactEdit,
}

var contactRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a contact.",
	Args:  cobra.MinimumNArgs(1),
	Run:   manage.ContactRemove,
}

var contactViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Look at the contact(s) listed.",
	Run:   manage.ContactView,
}

var contactEmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Email a contact.",
	Args:  cobra.ExactArgs(1),
	Run:   manage.Email,
}

func init() {

	contactCmd.Aliases = append(contactCmd.Aliases, "contacts")

	RootCmd.AddCommand(contactCmd)
	contactCmd.AddCommand(contactAddCmd)
	contactCmd.AddCommand(contactEditCmd)
	contactCmd.AddCommand(contactRemoveCmd)
	contactCmd.AddCommand(contactViewCmd)
	contactCmd.AddCommand(contactEmailCmd)

	// // View
	// contactViewCmd.Flags().StringP("output", "o", "", "save the results to a file")
	// contactViewCmd.Flags().Bool("force", false, "force the results to be saved to an existing file")
}
