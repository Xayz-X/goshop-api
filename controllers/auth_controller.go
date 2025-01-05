package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Xayz-X/goshop-api/models"
	"github.com/Xayz-X/goshop-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	UserCol *mongo.Collection
}

func NewUserCollection(userCol *mongo.Collection) *UserCollection {
	return &UserCollection{
		UserCol: userCol,
	}
}

func (u *UserCollection) GetUserByEmail(email string) (*models.UserRegisterPayload, error) {
	var user models.UserRegisterPayload
	result := u.UserCol.FindOne(context.Background(), bson.D{{Key: "email", Value: email}})
	if err := result.Err(); err != nil {
		return nil, err
	}
	result.Decode(user)
	return &user, nil
}

func (u *UserCollection) GetAllUsers() ([]*models.UserRegisterPayload, error) {
	var userList []*models.UserRegisterPayload

	// Perform the query
	cursor, err := u.UserCol.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background()) // Ensure the cursor is closed

	// Iterate through the cursor
	for cursor.Next(context.Background()) {
		var user models.UserRegisterPayload
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		userList = append(userList, &user)
	}

	// Check for errors during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}

func (u *UserCollection) CreateNewUser(payload models.UserRegisterPayload) (*mongo.InsertOneResult, error) {
	result, err := u.UserCol.InsertOne(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserCollection) ListAllUser(w http.ResponseWriter, r *http.Request) {
	userList, err := u.GetAllUsers()
	if err != nil {
		utils.WriterError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriterJSON(w, http.StatusCreated, &userList)
}

func (u *UserCollection) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userPayload models.UserRegisterPayload
	// decode the payload -> if error raise then write the error
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
		return
	}

	// check if any user exist with same email

	user, _ := u.GetUserByEmail(userPayload.Email)
	if user != nil {
		log.Println("User already exists with the email:", user.Email)
		utils.WriterError(w, http.StatusConflict, errors.New("user already exists with this email"))
		return
	}

	// create new user as same user does nto exist
	log.Println("New user created ", userPayload.Name)
	_, err = u.CreateNewUser(userPayload)
	if err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriterJSON(w, http.StatusCreated, &userPayload)
}
