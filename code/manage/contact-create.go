package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactCreate starts the client creation process.
func ContactCreate(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	logrus.Info("creating a contact through the CLI")

	defer logrus.Info("finished creating a contact through the CLI")

	var err error

	fast, _ := cmd.Flags().GetBool("fast")

	contact := db.Contact{}

	for {

		contact.Name, err = utils.CobraInput(cmd, "name", "Contact Name", "", false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		if len(contact.Name) > 0 {
			break
		}

	}

	for {

		contact.Email, err = utils.CobraInput(cmd, "email", "Contact Email", "", false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		if len(contact.Email) > 0 {
			break
		}

	}

	contact.Number, err = utils.CobraInput(cmd, "phone", "Contact Phone Number", "", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Contacted, err = utils.CobraInputBool(cmd, "contacted", "Contact was made", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Lead, err = utils.CobraInputBool(cmd, "lead", "Contact is a lead", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Subscriber, err = utils.CobraInputBool(cmd, "subscriber", "Contact is a subscriber", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	if !contact.Relationship.Subscriber {

		contact.Relationship.Customer, err = utils.CobraInputBool(cmd, "customer", "Contact is a customer", false, fast)
		if err != nil {
			logrus.Warn(err)
		}

	}

	contact.Relationship.Advocate, err = utils.CobraInputBool(cmd, "advocate", "Contact is a advocate", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Other, err = utils.CobraInput(cmd, "other", "Other customer relationship", "N/A", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	if err := contact.Create(); err != nil {
		logrus.Fatal(err)
	}

}
