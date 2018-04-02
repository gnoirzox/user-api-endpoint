package locations

import (
	"database/sql"
	"log"

	"../utils"
)

type Location struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	UserId    int64   `json:"user_id,omitempty"`
}

func (location *Location) IsValidLongitude() bool {
	if location.Longitude < -180 || location.Longitude > 180 {
		log.Println("Wrong value set for the Location.Longitude")

		return false
	}

	return true
}

func (location *Location) IsValidLatitude() bool {
	if location.Latitude < -90 || location.Latitude > 90 {
		log.Println("Wrong value set for the Location.Latitude")

		return false
	}

	return true
}

func (location *Location) IsValidUserId() bool {
	db := utils.OpenDBConnection()

	defer db.Close()

	row := db.QueryRow("SELECT id FROM Users WHERE id = ?")
	err := row.Scan(&location.UserId)

	if err != nil {
		log.Println(err.Error())

		return false
	}

	if err == sql.ErrNoRows {
		return false
	}

	return true
}
