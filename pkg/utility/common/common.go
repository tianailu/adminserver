package common

import "strconv"

func IsNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func ToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return result
}
