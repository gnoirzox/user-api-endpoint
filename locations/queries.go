package locations

import (
	"database/sql"
	"log"

	"../utils"
)

func SetCoordinates(location *Location) error {
	_, err := GetCoordinates(location.UserId)

	if err == sql.ErrNoRows {
		err = InsertCoordinates(location)
	} else {
		err = UpdateCoordinates(location)
	}

	if err != nil {
		log.Println(err.Error())

		return err
	}

	return nil
}

func InsertCoordinates(location *Location) error {
	db := utils.OpenDBConnection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Locations (longitude, latitude, user_id) VALUES (?,?,?)")
	defer insert.Close()

	if err != nil {
		log.Println(err.Error())
	}

	_, err = insert.Exec(&location.Longitude, &location.Latitude, &location.UserId)

	if err != nil {
		log.Println(err.Error())

		return err
	}

	return nil
}

func UpdateCoordinates(location *Location) error {
	db := utils.OpenDBConnection()

	transaction, err := db.Begin()
	defer transaction.Rollback()

	if err != nil {
		log.Println(err.Error())

		return err
	}

	update, err := transaction.Prepare("UPDATE Locations SET `longitude`=?, `latitude`=? WHERE `user_id`=?")
	defer update.Close()

	if err != nil {
		log.Println(err.Error())

		return err
	}

	_, err = update.Exec(&location.Longitude, &location.Latitude, &location.UserId)

	if err != nil {
		log.Println(err.Error())

		return err
	}

	err = transaction.Commit()

	if err != nil {
		log.Println(err.Error())

		return err
	}

	return nil
}

func GetCoordinates(UserId int64) (*Location, error) {
	db := utils.OpenDBConnection()
	defer db.Close()

	row := db.QueryRow("SELECT longitude, latitude, user_id FROM Locations WHERE user_id=?", UserId)

	var longitude float64
	var latitude float64
	var userId int64

	err := row.Scan(&longitude, &latitude, &userId)

	if err != nil {
		log.Println(err.Error())

		return nil, err
	}

	location := &Location{Longitude: longitude, Latitude: latitude, UserId: userId}

	return location, nil
}
