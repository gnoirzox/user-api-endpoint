package utils

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "Snatch:AndHide@tcp(127.0.0.1:3306)/SnatchHQ")

	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func FormatMapsUrl(longitude float64, latitude float64) string {
	longitudeString := strconv.FormatFloat(longitude, 'f', -1, 64)
	latitudeString := strconv.FormatFloat(latitude, 'f', -1, 64)

	return "https://www.google.com/maps/@" + latitudeString + "," + longitudeString + ",14z"
}

func ReturnJsonResponse(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Println(err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}
