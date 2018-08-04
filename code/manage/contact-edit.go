package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ContactEdit changes a contact based on the given ID.
func ContactEdit(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	id, err := utils.UfS(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	contact, update, rel := db.Contact{}, db.Contact{}, db.Relationship{}
	contact.ID = id

	if err := contact.Query(); err != nil {
		logrus.Fatal(err)
	}

	if err := db.DB.Model(&contact).Related(&rel).Error; err != nil {
		logrus.Fatal(err)
	}

	update.Name, err = input(cmd, "name", "Client Name", contact.Name, false, false)
	if err != nil {
		logrus.Fatal(err)
	}

	update.Email, err = input(cmd, "email", "Client Email", contact.Email, false, false)
	if err != nil {
		logrus.Fatal(err)
	}

	update.Number, err = input(cmd, "phone", "Client Phone Number", contact.Number, false, false)
	if err != nil {
		logrus.Fatal(err)
	}

	update.Contacted, err = inputBool(cmd, "contacted", "Contact was made", contact.Contacted, false)
	if err != nil {
		logrus.Warn(err)
	}

	rel.Lead, err = inputBool(cmd, "lead", "Contact is a lead", rel.Lead, false)
	if err != nil {
		logrus.Warn(err)
	}

	rel.Subscriber, err = inputBool(cmd, "subscriber", "Contact is a subscriber", rel.Subscriber, false)
	if err != nil {
		logrus.Warn(err)
	}

	if !rel.Subscriber {
		rel.Customer, err = inputBool(cmd, "customer", "Contact is a customer", rel.Customer, false)
		if err != nil {
			logrus.Warn(err)
		}
	}

	rel.Advocate, err = inputBool(cmd, "advocate", "Contact is a advocate", rel.Advocate, false)
	if err != nil {
		logrus.Warn(err)
	}

	rel.Other, err = input(cmd, "other", "Other customer relationship", rel.Other, false, false)
	if err != nil {
		logrus.Warn(err)
	}

	update.Relationship = rel

	if err := contact.Update(update); err != nil {
		logrus.Fatal(err)
	}

}
