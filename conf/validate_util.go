package conf

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// IsEmpty checks if a string is empty or only contains whitespace.
func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsValidEmail checks if an email is in a valid format.
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidPhoneNumber checks if a string is a valid phone number (basic validation).
func IsValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

// IsPositiveNumber checks if a string can be converted to a positive number.
func IsPositiveNumber(str string) bool {
	if len(str) == 0 {
		return false
	}
	var num float64
	_, err := fmt.Sscanf(str, "%f", &num)
	return err == nil && num > 0
}

// IsValidDate checks if a string is in the format of YYYY-MM-DD (basic validation).
func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// IsValidURL checks if a string is a valid URL.
func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^https?://(?:www\.)?[a-zA-Z0-9-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(url)
}

// IsAlphanumeric checks if a string contains only alphanumeric characters.
func IsAlphanumeric(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(str)
}

// IsValidUsername checks if a username is between 3 to 15 characters and contains only letters, numbers, and underscores.
func IsValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,15}$`)
	return re.MatchString(username)
}

// IsValidPostalCode checks if a postal code is valid (basic validation).
func IsValidPostalCode(postalCode string) bool {
	re := regexp.MustCompile(`^\d{5}(-\d{4})?$`) // US format
	return re.MatchString(postalCode)
}
