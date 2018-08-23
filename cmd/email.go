package cmd

import (
	"github.com/spf13/cobra"
)

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Work with emails.",
}

var emailCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a email and send it.",
	// Run:   email.Create,
}

var emailRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a email.",
	// Run:   email.Remove,
}

var emailViewCmd = &cobra.Command{
	Use:   "view",
	Short: "view a email.",
	// Run:   email.View,
}

func init() {

	RootCmd.AddCommand(emailCmd)

	emailCmd.AddCommand(emailCreateCmd)
	emailCmd.AddCommand(emailRemoveCmd)
	emailCmd.AddCommand(emailViewCmd)

}
