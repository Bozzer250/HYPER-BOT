package models

import (
	"context"
	"hyperbot/configs"
)

type Account struct {
	ID          string  `json:"id" bson:"_id"`
	Status      string  `json:"status" bson:"status"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
	UserID      string  `json:"user_id" bson:"user_id"`
	AccountType string  `json:"account_type" bson:"account_type"`
	DailyRate   float64 `json:"daily_rate" bson:"daily_rate"`
}

// create a new account
func CreateAccount(account Account) error {
	_, err := configs.MI.DB.Collection("accounts").InsertOne(context.TODO(), account)
	if err != nil {
		return err
	}
	return nil
}

func GetAccountsByUserID(userID string) ([]Account, error) {
	var accounts []Account
	cursor, err := configs.MI.DB.Collection("accounts").Find(context.TODO(), map[string]string{"user_id": userID})
	if err != nil {
		return accounts, err
	}
	if err = cursor.All(context.TODO(), &accounts); err != nil {
		return accounts, err
	}
	return accounts, nil
}

func UpdateAccountStatus(accountID, status string) error {
	_, err := configs.MI.DB.Collection("accounts").UpdateOne(context.TODO(), map[string]string{"_id": accountID}, map[string]string{"status": status})
	if err != nil {
		return err
	}
	return nil
}

func GetAccountByID(accountID string) (Account, error) {
	var account Account
	err := configs.MI.DB.Collection("accounts").FindOne(context.TODO(), map[string]string{"_id": accountID}).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}
