package manage

import (
	"broadcastle.co/code/crm/code/tui"
	"broadcastle.co/code/crm/code/utils"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactEdit changes a contact based on the given ID.
func ContactEdit(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, contact := range contacts {

		app := &tui.App{
			tview.NewApplication(),
		}

		form := app.Form(contact)

		form.SetBorder(true).
			SetTitle("Edit " + contact.Name).
			SetTitleAlign(tview.AlignCenter)

		if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
			logrus.Fatal(err)
		}

	}

}
