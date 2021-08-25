package api

import (
	"errors"
	"log"
	"net/mail"
	"regexp"
)

var regexpPhone *regexp.Regexp = regexp.MustCompile(`^\d{8,15}$`)

func validateId(id int64) error {
	if id <= 0 {
		log.Printf("Invalid id is provided: %d", id)
		return errors.New("Invalid id is provided")
	}
	return nil
}

func validateName(name string) error {
	if len(name) == 0 {
		return nil
	}
	if len(name) == 1 {
		log.Printf("Invalid name is provided: %s", name)
		return errors.New("Invalid name is provided")
	}
	return nil
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return nil
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("Invalid email is provided: %s", email)
		return errors.New("Invalid email is provided")
	}
	return nil
}

func validatePhone(phone string) error {
	if len(phone) == 0 {
		return nil
	}
	if !regexpPhone.MatchString(phone) {
		log.Printf("Invalid phone number is provided: %s", phone)
		return errors.New("Invalid phone number is provided")
	}
	return nil
}
