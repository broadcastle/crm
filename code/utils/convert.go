package utils

import "strconv"

// UfS takes a string input and returns a uint.
func UfS(input string) (uint, error) {

	id64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id64), nil

}
