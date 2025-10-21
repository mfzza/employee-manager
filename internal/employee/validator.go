package employee

import (
	"fmt"
	"regexp"
)

var (
	emailRegex = regexp.MustCompile("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@(?:\\w+\\.)+\\w+$")
	phoneRegex = regexp.MustCompile("^[0-9]{2,3}-[0-9]{3,4}-[0-9]{4}$")
)

func ValidateEmail(email string) error {
	if emailRegex.MatchString(email) {
		return nil
	}
	return fmt.Errorf("The email you entered is is in a valid format. Please check and try again.\n")
}

func ValidatePhone(phone string) error {
	if phoneRegex.MatchString(phone) {
		return nil
	}
	return fmt.Errorf("Phone number you entered is not in valid format. Example of valid format\nXX-XXX-XXXX, XX-XXXX-XXXX, XXX-XXX-XXXX or XXX-XXXX-XXXX\n")
}
