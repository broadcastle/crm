package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/tui"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func contactFromFlags(cmd *cobra.Command) (contact db.Contact, err error) {

	contact.Name, err = cmd.Flags().GetString("name")
	if err != nil {
		return
	}

	contact.Email, err = cmd.Flags().GetString("email")
	if err != nil {
		return
	}

	contact.Number, err = cmd.Flags().GetString("phone")
	if err != nil {
		return
	}

	contact.Contacted, err = cmd.Flags().GetBool("contacted")
	if err != nil {
		return
	}

	return

}

// ContactCreate starts the client creation process.
func ContactCreate(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	fast, err := cmd.Flags().GetBool("fast")
	if err != nil {
		logrus.Warn(err)
	}

	contact, err := contactFromFlags(cmd)
	if err != nil {
		logrus.Fatal(err)
	}

	if fast && contact.Name != "" && contact.Email != "" {

		if err := contact.Create(); err != nil {
			logrus.Fatal(err)
		}

		return

	}

	app := &tui.App{
		tview.NewApplication(),
	}

	form := app.Form(contact)

	form.SetBorder(true).
		SetTitle("Create a New Contact").
		SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		logrus.Fatal(err)
	}

}
