package stringutil

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Json(obj interface{}) string {
	// Marshal the object into a JSON string
	jsonData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error marshalling object to JSON string: %v\n", err)
		return ""
	}
	return string(jsonData)
}

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func IsNotEmpty(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

func IsPositiveNumber(str string) bool {
	if len(str) == 0 {
		return false
	}
	var num float64
	_, err := fmt.Sscanf(str, "%f", &num)
	return err == nil && num > 0
}

func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^https?://(?:www\.)?[a-zA-Z0-9-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(url)
}

func IsAlphanumeric(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(str)
}

func IsValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,15}$`)
	return re.MatchString(username)
}

func IsValidPostalCode(postalCode string) bool {
	re := regexp.MustCompile(`^\d{5}(-\d{4})?$`)
	return re.MatchString(postalCode)
}

func IsValidThaiPostalCode(postalCode string) bool {
	re := regexp.MustCompile(`^\d{5}$`)
	return re.MatchString(postalCode)
}
