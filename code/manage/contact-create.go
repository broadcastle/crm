package manage

import (
	"broadcastle.co/code/crm/code/tui"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactCreate starts the client creation process.
func ContactCreate(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	app := &tui.App{
		tview.NewApplication(),
	}

	form := app.ContactForm(0)

	form.SetBorder(true).
		SetTitle("Create a New Contact").
		SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		logrus.Fatal(err)
	}

}
