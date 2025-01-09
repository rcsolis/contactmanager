package validation

import (
	"regexp"
)

func isValidPhoneNumber(phoneNumber string) bool {
	validPhoneNumber := regexp.MustCompile(`^\+[0-9]{2}-[0-9]{2}-[0-9]{8}$`)
	return validPhoneNumber.MatchString(phoneNumber)
}

func isValidEmail(email string) bool {
	validEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validEmail.MatchString(email)
}

func ValidateData(id string, name string, email string, phoneNumber string) bool {
	return id != "" && name != "" && isValidEmail(email) && isValidPhoneNumber(phoneNumber)
}

func ValidateInput(query string) bool {
	return query == ""
}

func ValidateUpdateData(id string, name string, email string, phoneNumber string) bool {
	if id == "" {
		return false
	}
	if email != "" {
		if !isValidEmail(email) {
			return false
		}
	}
	if phoneNumber != "" {
		if !isValidPhoneNumber(phoneNumber) {
			return false
		}
	}
	return true
}
