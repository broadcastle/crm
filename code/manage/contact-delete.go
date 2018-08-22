package manage

import (
	"broadcastle.co/code/crm/code/tui"
	"broadcastle.co/code/crm/code/utils"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactRemove removes a contact with a given ID from the database.
func ContactRemove(cmd *cobra.Command, args []string) {

	logrus.Debug("beginning the contact removal process")

	Init()
	defer Close()

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	for x := range contacts {

		app := &tui.App{
			tview.NewApplication(),
		}

		modal := app.Remove(contacts[x])

		if err := app.SetRoot(modal, true).SetFocus(modal).Run(); err != nil {
			logrus.Fatal(err)
		}

	}

}
