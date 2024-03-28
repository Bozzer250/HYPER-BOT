package models

import (
	"context"
	"hyperbot/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID           string    `json:"id" bson:"_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`
	ReferralCode string    `json:"referral_code"`
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

func GetUserById(id string) User {
	var user User
	_ = configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user
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

func GetUserIdFromReferralCode(referralCode string) (string, error) {
	var user User
	err := configs.MI.DB.Collection("users").FindOne(context.TODO(), bson.M{"referral_code": referralCode}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
