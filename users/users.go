package users

import (
	"log"
	"net/mail"
	"strconv"
	"strings"
)

type User struct {
	Username    string `json:"username,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	LocationUrl string `json:"location"`
}

func (u *User) IsValidUsername() bool {
	if len(u.Username) < 3 || len(u.Username) > 12 {
		log.Println("Wrong lenght for the User.Username. It should be between 3 and 12 characters.")

		return false
	}

	if strings.Count(strings.ToLower(u.Username), "dog") != strings.Count(strings.ToLower(u.Username), "bulldog") {
		log.Println("'dog' is not allowed in the User.Username")

		return false
	}

	if strings.Count(strings.ToLower(u.Username), "cat") !=
		(strings.Count(strings.ToLower(u.Username), "catfish") + strings.Count(strings.ToLower(u.Username), "scatter")) {
		log.Println("'cat' is not allowed in the User.Username")

		return false
	}

	if strings.Count(strings.ToLower(u.Username), "horse") != strings.Count(strings.ToLower(u.Username), "seahorse") {
		log.Println("'horse' is not allowed in the User.Username")

		return false
	}

	return true
}

func (u *User) IsValidPhoneNumber() bool {
	if len(u.PhoneNumber) < 11 || len(u.PhoneNumber) > 11 {
		return false
	}

	for _, value := range []byte(u.PhoneNumber) {
		_, err := strconv.Atoi(string(value))

		if err != nil {
			log.Println("The provided phone number is not valid for User.PhoneNumber")
			log.Println(err.Error())

			return false
		}
	}

	return true
}

func (u *User) IsValidEmail() bool {
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		log.Println("The provided email address is not respecting the RFC 5322 format for User.Email")
		log.Println(err.Error())

		return false
	}

	return true
}
