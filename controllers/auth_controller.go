package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Xayz-X/goshop-api/models"
	"github.com/Xayz-X/goshop-api/utils"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userPayload models.UserRegisterPayload

	// decode the payload -> if error raise then write the error
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
		return
	}

	// connect to db and check for same email else register new user.
	log.Println("New user created ", userPayload.Name)
	// send back the response of newly created user.
	utils.WriterJSON(w, http.StatusCreated, &userPayload)
}
