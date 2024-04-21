package utils

import (
	"net/mail"
)

func isEmail (text string) bool {
	_, err mail.ParseAddress(email)
	return err == nil
}