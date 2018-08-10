package utils

import (
	"strconv"

	"broadcastle.co/code/crm/code/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CobraInput returns a string from a flag.
func CobraInput(cmd *cobra.Command, flag string, hint string, value string, skip bool, hidden bool) (string, error) {

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
		return InputHidden(hint, result)
	}

	return Input(hint, result)

}

// CobraInputBool returns a bool from a flag.
func CobraInputBool(cmd *cobra.Command, flag string, hint string, value bool, skip bool) (bool, error) {

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

	return InputBool(hint, result)

}

// Contacts returns contacts from a 'id' flag or from arguments.
func Contacts(cmd *cobra.Command, args []string) ([]db.Contact, error) {

	contacts := []db.Contact{}

	for _, x := range args {

		single := db.Contact{}

		id64, err := strconv.ParseUint(x, 10, 64)

		if err == nil {
			single.ID = uint(id64)
		} else {
			single.Slug = x
		}

		if err := single.Search(); err != nil {
			logrus.Warn(err)
			break
		}

		contacts = append(contacts, single)

	}

	return contacts, nil

}
