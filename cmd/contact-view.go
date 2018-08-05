package cmd

import (
	"broadcastle.co/code/crm/code/manage"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Look at the contact(s) listed.",
	Run:   manage.ContactView,
}

func init() {
	contactCmd.AddCommand(viewCmd)

	viewCmd.Flags().StringP("output", "o", "", "save the results to a file")
	viewCmd.Flags().Bool("force", false, "force the results to be saved to an existing file")

}
