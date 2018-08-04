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
