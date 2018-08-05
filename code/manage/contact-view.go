package manage

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactView returns the selected contact given their ID's.
func ContactView(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	contacts := []db.Contact{}
	var result bytes.Buffer

	if len(args) < 1 {
		logrus.Info("nothing returned")
		return
	}

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Warn(err)
			break
		}

		contact := db.Contact{}
		contact.ID = id

		if err := contact.Query(); err != nil {
			logrus.Warn(err)
			break
		}

		// Get the relationships of a status.
		var e db.Relationship

		if err := db.DB.Model(&contact).Related(&e).Error; err != nil {
			logrus.Warn(err)
			break
		}

		contact.Relationship = e

		// Append the contact to the contacts array.
		contacts = append(contacts, contact)

	}

	output, err := json.Marshal(&contacts)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Indent(&result, output, "", " "); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(string(result.Bytes()))

	dir, err := cmd.Flags().GetString("output")
	if err != nil {
		logrus.Warn(err)
		return
	}

	force, _ := cmd.Flags().GetBool("force")

	if dir != "" {

		// Create a file with the dirs.
		if _, err := os.Stat(dir); err == nil && !force {
			logrus.Warn("file exists, force flag required")
		}

		if err := ioutil.WriteFile(dir, result.Bytes(), 0644); err != nil {
			logrus.Fatal(err)
		}

	}

}
