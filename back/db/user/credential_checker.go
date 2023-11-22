package user

import "net/mail"

func CheckMailValidity(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func CheckPassWordStrength(pass string) error { // to complette
	return nil
}
