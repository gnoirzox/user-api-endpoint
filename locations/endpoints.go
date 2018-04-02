package locations

import (
	"encoding/json"
	"log"
	"net/http"

	"../utils"
)

func PostLocation(w http.ResponseWriter, r *http.Request) {
	var location Location

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&location)

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid resquest payload"})

		return
	}

	defer r.Body.Close()

	if !location.IsValidLongitude() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted Longitude is not valid. It should be between -180 and 180."})

		return
	}

	if !location.IsValidLatitude() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted Latitude is not valid. It should be between -90 and 90."})

		return
	}

	if !location.IsValidUserId() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted user id is not valid. The user does not exist."})

		return
	}

	err = SetCoordinates(&location)

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	utils.ReturnJsonResponse(w, http.StatusOK, map[string]string{"result": "success"})
}
