package utils

import "github.com/spf13/cobra"

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
