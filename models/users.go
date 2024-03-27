package models

import (
	"context"
	"hyperbot/configs"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

// get user by phone
func GetUserByPhone(phone string) (User, error) {
	var user User
	err := configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// create new user
func CreateUser(user User) error {
	_, err := configs.MI.DB.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllActiveUsers() ([]User, error) {
	var users []User
	cursor, err := configs.MI.DB.Collection("users").Find(context.TODO(), bson.M{"status": "active"})
	if err != nil {
		return users, err
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return users, err
	}
	return users, nil
}
