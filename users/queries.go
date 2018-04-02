package users

import (
	"database/sql"
	"log"

	"../utils"
)

func InsertUser(user *User) error {
	db := utils.OpenDBConnection()

	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Users (username, phone_number, email) VALUES (?, ?, ?)")

	if err != nil {
		log.Println(err.Error())
	}

	_, err = insert.Exec(&user.Username, &user.PhoneNumber, &user.Email)

	defer insert.Close()

	if err != nil {
		log.Println(err.Error())

		return err
	}

	return nil
}

func RetrieveUser(UserId int64) (*User, error) {
	db := utils.OpenDBConnection()

	defer db.Close()

	row := db.QueryRow("SELECT Users.username, Users.phone_number, Users.email, Locations.longitude, Locations.latitude FROM Users LEFT JOIN Locations ON Users.id = Locations.user_id WHERE Users.id = ?", UserId)

	var username string
	var phoneNumber string
	var email string
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64

	err := row.Scan(&username, &phoneNumber, &email, &longitude, &latitude)

	if err != nil {
		log.Println(err.Error())

		return nil, err
	}

	var locationUrl string

	if longitude.Valid && latitude.Valid {
		locationUrl = utils.FormatMapsUrl(longitude.Float64, latitude.Float64)
	}

	user := &User{
		Username:    username,
		PhoneNumber: phoneNumber,
		Email:       email,
		LocationUrl: locationUrl,
	}

	return user, nil
}
