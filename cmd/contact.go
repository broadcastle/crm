package cmd

import (
	"github.com/spf13/cobra"
)

// contactCmd represents the client command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Manage your contacts.",
}

func init() {
	RootCmd.AddCommand(contactCmd)
}

func contactFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "n", "", "name for this contact")
	cmd.Flags().StringP("email", "e", "", "email for this contact")
	cmd.Flags().String("phone", "", "phone number for this contact")

	cmd.Flags().BoolP("contacted", "c", false, "contact was made")

	cmd.Flags().Bool("lead", false, "contact is a lead")
	cmd.Flags().Bool("customer", false, "contact is a customer")
	cmd.Flags().Bool("subscriber", false, "contact is a subscriber")
	cmd.Flags().Bool("advocate", false, "contact is a advocate")
	cmd.Flags().String("other", "", "some other contact relationship")
}
