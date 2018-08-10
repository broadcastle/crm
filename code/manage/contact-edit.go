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

	contacts, err := utils.Contacts(cmd, args)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, contact := range contacts {

		update, rel := db.Contact{}, db.Relationship{}

		if err := db.DB.Model(&contact).Related(&rel).Error; err != nil {
			logrus.Fatal(err)
		}

		update.Name, err = utils.CobraInput(cmd, "name", "Client Name", contact.Name, false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		update.Email, err = utils.CobraInput(cmd, "email", "Client Email", contact.Email, false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		update.Number, err = utils.CobraInput(cmd, "phone", "Client Phone Number", contact.Number, false, false)
		if err != nil {
			logrus.Fatal(err)
		}

		update.Contacted, err = utils.CobraInputBool(cmd, "contacted", "Contact was made", contact.Contacted, false)
		if err != nil {
			logrus.Warn(err)
		}

		rel.Lead, err = utils.CobraInputBool(cmd, "lead", "Contact is a lead", rel.Lead, false)
		if err != nil {
			logrus.Warn(err)
		}

		rel.Subscriber, err = utils.CobraInputBool(cmd, "subscriber", "Contact is a subscriber", rel.Subscriber, false)
		if err != nil {
			logrus.Warn(err)
		}

		if !rel.Subscriber {
			rel.Customer, err = utils.CobraInputBool(cmd, "customer", "Contact is a customer", rel.Customer, false)
			if err != nil {
				logrus.Warn(err)
			}
		}

		rel.Advocate, err = utils.CobraInputBool(cmd, "advocate", "Contact is a advocate", rel.Advocate, false)
		if err != nil {
			logrus.Warn(err)
		}

		rel.Other, err = utils.CobraInput(cmd, "other", "Other customer relationship", rel.Other, false, false)
		if err != nil {
			logrus.Warn(err)
		}

		update.Relationship = rel

		if err := contact.Update(update); err != nil {
			logrus.Fatal(err)
		}

	}

}
