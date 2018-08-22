package note

import (
	"bytes"
	"encoding/json"

	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/tui"
	"broadcastle.co/code/crm/code/utils"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Create a new note.
func Create(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	app := &tui.App{
		tview.NewApplication(),
	}

	form := app.NoteForm(0)

	form.SetBorder(true).
		SetTitle("Create A New Note").
		SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		logrus.Fatal(err)
	}

}

// Edit a note
func Edit(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	id, err := utils.UfS(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	app := &tui.App{
		tview.NewApplication(),
	}

	form := app.NoteForm(id)

	form.SetBorder(true).
		SetTitle("Edit A Note").
		SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		logrus.Fatal(err)
	}

}

// Remove a note
func Remove(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Warn(err)
			break
		}

		note := db.Note{}
		note.ID = id

		if err := note.Remove(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("note #%s was removed", x)

	}

}

// View a note
func View(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	notes := []db.Note{}
	var result bytes.Buffer
	var err error

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Warn(err)
			break
		}

		note := db.Note{}
		note.ID = id

		if err := note.Query(); err != nil {
			logrus.Warn(err)
			break
		}

		notes = append(notes, note)

	}

	if len(args) < 1 {

		notes, err = db.QueryNotes()
		if err != nil {
			logrus.Info("no notes")
			return
		}

	}

	var output []byte

	output, err = json.Marshal(&notes)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Indent(&result, output, "", " "); err != nil {
		logrus.Fatal(err)
	}

}
