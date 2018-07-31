package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Init starts the database.
func Init() {
	logrus.Debug("starting the usage of the database")
	db.Init()
}

// Close ends the manage package.
func Close() {
	logrus.Debug("ending the usage of the database")
	db.Close()
}

// ContactCreate starts the client creation process.
func ContactCreate(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	logrus.Info("creating a contact through the CLI")

	defer logrus.Info("finished creating a contact through the CLI")

	var err error

	fast, _ := cmd.Flags().GetBool("fast")

	contact := db.Contact{}

	contact.Name, err = input(cmd, "name", "Contact Name", "", false, false)
	if err != nil {
		logrus.Fatal(err)
	}

	contact.Email, err = input(cmd, "email", "Contact Email", "", false, false)
	if err != nil {
		logrus.Fatal(err)
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

	contact.Relationship.Customer, err = inputBool(cmd, "customer", "Contact is a customer", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Subscriber, err = inputBool(cmd, "subscriber", "Contact is a subscriber", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Advocate, err = inputBool(cmd, "advocate", "Contact is a advocate", false, fast)
	if err != nil {
		logrus.Warn(err)
	}

	contact.Relationship.Other, err = input(cmd, "other", "Other customer relationship", "", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	logrus.Debugf("basic client info\nName: %s\nEmail: %s\nPhone Number: %s", contact.Name, contact.Email, contact.Number)

	if err := contact.Create(); err != nil {
		logrus.Fatal(err)
	}

}

// ContactEdit changes a contact based on the given id.
func ContactEdit(cmd *cobra.Command, args []string) {

	db.Init()
	defer db.Close()

	id, err := utils.UfS(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	contact := db.Contact{}
	contact.ID = id

	if err := contact.Query(); err != nil {
		logrus.Fatal(err)
	}

	update := db.Contact{}

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

	rel, ord := db.Relationship{}, db.Relationship{}

	if err := db.DB.Model(&contact).Related(&rel).Error; err != nil {
		logrus.Fatal(err)
	}

	ord.Lead, err = inputBool(cmd, "lead", "Contact is a lead", rel.Lead, false)
	if err != nil {
		logrus.Warn(err)
	}

	ord.Customer, err = inputBool(cmd, "customer", "Contact is a customer", rel.Customer, false)
	if err != nil {
		logrus.Warn(err)
	}

	ord.Subscriber, err = inputBool(cmd, "subscriber", "Contact is a subscriber", rel.Subscriber, false)
	if err != nil {
		logrus.Warn(err)
	}

	ord.Advocate, err = inputBool(cmd, "advocate", "Contact is a advocate", rel.Advocate, false)
	if err != nil {
		logrus.Warn(err)
	}

	ord.Other, err = input(cmd, "other", "Other customer relationship", rel.Other, false, false)
	if err != nil {
		logrus.Warn(err)
	}

	logrus.Info(update)

	if err := contact.Update(update); err != nil {
		logrus.Fatal(err)
	}

	if err := db.DB.Model(&contact).Association("Relationship").Replace(ord).Error; err != nil {
		logrus.Fatal(err)
	}

}

// ContactRemove removes a contact with a given ID from the database.
func ContactRemove(cmd *cobra.Command, args []string) {

	logrus.Debug("beginning the contact removal process")

	db.Init()
	defer db.Close()

	for _, x := range args {

		id, err := utils.UfS(x)
		if err != nil {
			logrus.Fatal(err)
		}

		contact := db.Contact{}
		contact.ID = id

		if err := contact.Remove(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("contact with id #%s was removed", x)

	}

}

func input(cmd *cobra.Command, flag string, hint string, value string, skip bool, hidden bool) (string, error) {

	result, err := cmd.Flags().GetString(flag)
	if err != nil {
		return "", err
	}

	if result == "" && value != "" {
		result = value
	}

	if skip {
		return result, nil
	}

	if hidden {
		return utils.InputHidden(hint, result)
	}

	return utils.Input(hint, result)

}

func inputBool(cmd *cobra.Command, flag string, hint string, value bool, skip bool) (bool, error) {
	result, err := cmd.Flags().GetBool(flag)
	if err != nil {
		return false, err
	}

	if value {
		result = value
	}

	if skip {
		return result, nil
	}

	return utils.InputBool(hint, result)

}
