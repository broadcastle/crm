package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactRemove removes a contact with a given ID from the database.
func ContactRemove(cmd *cobra.Command, args []string) {

	logrus.Debug("beginning the contact removal process")

	db.Init()
	defer db.Close()

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	for x := range contacts {

		if err := contacts[x].Remove(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("%s was removed as a contact", contacts[x].Name)

	}

}
