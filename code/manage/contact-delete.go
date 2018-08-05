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

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			// logrus.Fatal(err)
			logrus.Warn(err)
			break
		}

		contact := db.Contact{}
		contact.ID = id

		if err := contact.Remove(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("contact with id #%s was removed", x)

	}

}
