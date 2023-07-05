package user

import "net/mail"

func CheckMailValidity(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
