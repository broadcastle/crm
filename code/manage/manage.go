package manage

import (
	"broadcastle.co/code/crm/code/db"
	"broadcastle.co/code/crm/code/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Init starts the database.
func Init() {
	db.InitSQLite()

	db.DB.AutoMigrate(&db.Contact{})

}

// Close ends the manage package.
func Close() {
	db.DB.Close()
}

// ContactCreate starts the client creation process.
func ContactCreate(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

	logrus.Info("creating a client through cli")

	defer logrus.Info("finished cli client creation process")

	var err error

	fast, _ := cmd.Flags().GetBool("fast")

	name, err := input(cmd, "name", "Client Name", "", false, false)
	if err != nil {
		logrus.Fatal(err)
	}

	email, err := input(cmd, "email", "Client Email", "", false, true)
	if err != nil {
		logrus.Fatal(err)
	}

	phone, err := input(cmd, "phone", "Client Phone Number", "", fast, false)
	if err != nil {
		logrus.Warn(err)
	}

	logrus.Debugf("basic client info\nName: %s\nEmail: %s\nPhone Number: %s", name, email, phone)

	c := db.Contact{
		Name:   name,
		Email:  email,
		Number: phone,
	}

	if err := c.Create(); err != nil {
		logrus.Fatal(err)
	}

}

// ContactEdit changes a contact based on the given id.
func ContactEdit(cmd *cobra.Command, args []string) {

	Init()
	defer Close()

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

	if err := contact.Update(update); err != nil {
		logrus.Fatal(err)
	}

}

// ContactRemove removes a contact with a given ID from the database.
func ContactRemove(cmd *cobra.Command, args []string) {

	logrus.Debug("beginning removal process.")

	Init()
	defer Close()

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

		logrus.Debug("contact was removed.")

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
