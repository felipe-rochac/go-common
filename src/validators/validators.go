package validators

import (
	"regexp"
)

// ValidateEmail validates an email address.
func ValidateEmail(email string) bool {
	// Email regex pattern
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// ValidatePassword validates a password.
func ValidatePassword(password string) bool {
	// Password regex pattern
	pattern := `^.{8,}$`
	match, _ := regexp.MatchString(pattern, password)
	return match
}

func ValidatePhone(phone string) bool {
	// Phone regex pattern
	pattern := `^\+?[\d\s]+$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

func ValidateCountryCode(countryCode string) bool {
	// Country code regex pattern
	pattern := `^[A-Z]{2}$`
	match, _ := regexp.MatchString(pattern, countryCode)
	return match
// CountryCode represents a country code as an enum.
