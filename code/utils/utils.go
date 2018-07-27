package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// Input is shorthand for InputVisible.
func Input(hint, value string) (string, error) {
	return InputVisible(hint, value)
}

// InputHidden provides a hint for a input, but does not show the user input.
func InputHidden(hint, value string) (string, error) {

	// Generate a maximum of 10 stars for hidden values.
	var stars string

	for x := range value {

		stars += "*"

		if x > 10 {
			break
		}
	}

	// Display the prompt.
	switch value != "" {
	case true:
		fmt.Printf("%s [hidden][%s]: ", hint, stars)
	default:
		fmt.Printf("%s [hidden]: ", hint)

	}

	// Receive the prompt.
	hidden, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}

	fmt.Println()

	// Return the input or default value.
	result := string(hidden)

	if result == "" {
		return value, nil
	}

	return result, nil
}

// InputVisible provides a hint for a input, and the user input is visible.
func InputVisible(hint, value string) (string, error) {

	// Display prompt
	switch value != "" {
	case true:
		fmt.Printf("%s [%s]: ", hint, value)
	default:
		fmt.Printf("%s: ", hint)

	}

	// Receive text
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Return input or default value.
	text = strings.TrimSpace(text)

	if text == "" {
		return value, nil
	}

	return text, nil

}
