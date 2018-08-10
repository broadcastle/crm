package cmd

import (
	"broadcastle.co/code/crm/code/note"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
}

var noteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note.",
	Args:  cobra.NoArgs,
	Run:   note.Create,
}

var noteEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing note.",
	Args:  cobra.ExactArgs(1),
	Run:   note.Edit,
}

var noteRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a note.",
	Args:  cobra.MinimumNArgs(1),
	Run:   note.Remove,
}

var noteViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Look at the note(s) listed.",
	Run:   note.View,
}

func init() {

	noteCmd.Aliases = append(contactCmd.Aliases, "notes")

	RootCmd.AddCommand(noteCmd)

	noteCmd.AddCommand(noteAddCmd)
	noteCmd.AddCommand(noteEditCmd)
	noteCmd.AddCommand(noteRemoveCmd)
	noteCmd.AddCommand(noteViewCmd)

	noteFlags(noteAddCmd)
	noteFlags(noteEditCmd)

	noteAddCmd.Flags().Int("id", 0, "id of contact")
	noteAddCmd.MarkFlagRequired("id")

	noteViewCmd.Flags().BoolP("raw", "r", false, "show the raw output")

}

func noteFlags(cmd *cobra.Command) {

	cmd.Flags().BoolP("task", "t", false, "this is/was a task")
	cmd.Flags().BoolP("call", "c", false, "this is/was a call")
	cmd.Flags().BoolP("email", "e", false, "this is/was a email")
	cmd.Flags().String("header", "", "header for this note")
	cmd.Flags().BoolP("fast", "f", false, "create a quick note")

}
