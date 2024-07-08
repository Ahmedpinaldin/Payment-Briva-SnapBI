package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// StringInSet validates that a string is in a given set of allowed values.
func StringInSet(value string, allowedValues map[string]bool, fieldName string) error {
	if value == "" || !allowedValues[value] {
		return fmt.Errorf("%s must not be null and should be one of the allowed values", fieldName)
	}
	return nil
}

// IntegerInRange validates that a string represents an integer within a specified range.

func JustInteger(value string, fieldName string) error {
	if value == "" {
		return fmt.Errorf("%s must not be null", fieldName)
	}

	_, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s must be an integer", fieldName)
	}

	return nil
}

// IntegerInRange validates that a string represents an integer within a specified range.
func IntegerInRange(value string, min, max int, fieldName string) error {
	if value == "" {
		return fmt.Errorf("%s must not be null", fieldName)
	}

	numValue, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s must be an integer", fieldName)
	}

	if numValue < min || numValue > max {
		return fmt.Errorf("%s must be between %d and %d", fieldName, min, max)
	}

	return nil
}

// StringMaxLength validates that a string has a maximum length.
func StringMaxLength(value string, maxLength int, fieldName string) error {
	if utf8.RuneCountInString(value) > maxLength {
		return fmt.Errorf("%s must not exceed %d characters", fieldName, maxLength)
	}
	return nil
}

// StringIsAlpha validates that a string contains only alphabetic characters.
func StringIsAlpha(value, fieldName string) error {
	if value != "0" {
		for _, r := range value {
			if !unicode.IsLetter(r) {
				return fmt.Errorf("%s must contain only alphabetic characters", fieldName)
			}
		}
		return nil
	}
	return nil
}

func ValidateNoRekening(noRekening string) error {
	// Regular expression pattern for numbers, dots, and alphabets
	pattern := "^[0-9A-Za-z.]+$"
	re := regexp.MustCompile(pattern)

	// Check if the input matches the pattern
	if !re.MatchString(noRekening) {
		return fmt.Errorf("NO_REKENING must contain only numbers, dots, and alphabets")
	}

	return nil
}
func ValidateNumberAlpha(userApprove string) error {
	// Regular expression pattern for numbers and alphabets
	pattern := "^[0-9A-Za-z]+$"
	re := regexp.MustCompile(pattern)

	// Check if the input matches the pattern
	if !re.MatchString(userApprove) {
		return fmt.Errorf("user_approve must contain only numbers and alphabets")
	}

	return nil
}
