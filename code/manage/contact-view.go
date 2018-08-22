package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/tui"
	"broadcastle.co/code/crm/code/utils"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactView returns the selected contact given their ID's.
func ContactView(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	if len(contacts) < 1 {
		contacts, err = db.QueryContacts()
		if err != nil {
			logrus.Fatal(err)
		}
	}

	app := &tui.App{
		tview.NewApplication(),
	}

	table := app.Table(contacts)

	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		logrus.Fatal(err)
	}

	// dir, err := cmd.Flags().GetString("output")
	// if err != nil {
	// 	logrus.Warn(err)
	// 	return
	// }

	// force, _ := cmd.Flags().GetBool("force")

	// if dir != "" {

	// 	// Create a file with the dirs.
	// 	if _, err := os.Stat(dir); err == nil && !force {
	// 		logrus.Warn("file exists, force flag required")
	// 	}

	// 	if err := ioutil.WriteFile(dir, result.Bytes(), 0644); err != nil {
	// 		logrus.Fatal(err)
	// 	}

	// }

}
