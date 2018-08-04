package manage

import (
	"broadcastle.co/code/crm/code/db"
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

		contact.Name, err = input(cmd, "name", "Contact Name", "", false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		if len(contact.Name) > 0 {
			break
		}

	}

	for {

		contact.Email, err = input(cmd, "email", "Contact Email", "", false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		if len(contact.Email) > 0 {
			break
		}

	}

	contact.Number, err = input(cmd, "phone", "Contact Phone Number", "", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Contacted, err = inputBool(cmd, "contacted", "Contact was made", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Lead, err = inputBool(cmd, "lead", "Contact is a lead", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Subscriber, err = inputBool(cmd, "subscriber", "Contact is a subscriber", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	if !contact.Relationship.Subscriber {

		contact.Relationship.Customer, err = inputBool(cmd, "customer", "Contact is a customer", false, fast)
		if err != nil {
			logrus.Warn(err)
		}

	}

	contact.Relationship.Advocate, err = inputBool(cmd, "advocate", "Contact is a advocate", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Other, err = input(cmd, "other", "Other customer relationship", "N/A", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	if err := contact.Create(); err != nil {
		logrus.Fatal(err)
	}

}
