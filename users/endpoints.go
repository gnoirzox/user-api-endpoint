package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "The provided user id is not valid."})

		return
	}

	user, err := RetrieveUser(int64(userId))

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusNotFound, map[string]string{"error": "The user does not exist."})

		return
	}

	utils.ReturnJsonResponse(w, http.StatusOK, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var u User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	defer r.Body.Close()

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid resquest payload"})

		return
	}

	if !u.IsValidUsername() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted username is not valid. It should not contain cat, dog or horse with a length between 3 and 12 characters."})

		return
	}

	if !u.IsValidPhoneNumber() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted phoneNumber is not valid. It should have a length of 11."})

		return
	}

	if !u.IsValidEmail() {
		utils.ReturnJsonResponse(w, http.StatusBadRequest,
			map[string]string{"error": "The submitted Email is not valid."})

		return
	}

	err = InsertUser(&u)

	if err != nil {
		log.Println(err.Error())

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	utils.ReturnJsonResponse(w, http.StatusOK, map[string]string{"result": "success"})
}
