package cmd

import (
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Manage clients and their data.",
}

func init() {
	RootCmd.AddCommand(clientCmd)
}
