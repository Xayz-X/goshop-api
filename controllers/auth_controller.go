package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Xayz-X/goshop-api/models"
	"github.com/Xayz-X/goshop-api/services"
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

func (u *UserCollection) GetUserByEmail(email string) (*mongo.SingleResult, error) {
	result := u.UserCol.FindOne(context.Background(), bson.D{{Key: "email", Value: email}})
	if err := result.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserCollection) CreateNewUser(payload models.UserRegisterPayload) (*mongo.InsertOneResult, error) {
	result, err := u.UserCol.InsertOne(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userPayload models.UserRegisterPayload
	// decode the payload -> if error raise then write the error
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
		return
	}

	client, database := services.ConnectToDb()
	defer client.Disconnect(context.Background())

	// create user collection
	userCol := database.Collection("user")
	uc := NewUserCollection(userCol)

	// connect to db and check for same email else register new user.
	log.Println("New user created ", userPayload.Name)
	_, err = uc.CreateNewUser(userPayload)
	if err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
	}
	utils.WriterJSON(w, http.StatusCreated, &userPayload)
}
