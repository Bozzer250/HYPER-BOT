package models

import (
	"context"
	"hyperbot/configs"

	"go.mongodb.org/mongo-driver/bson"
)

type Transactions struct {
	ID              string `json:"id" bson:"_id"`
	Status          string `json:"status" bson:"status"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
	UserID          string `json:"user_id" bson:"user_id"`
	AccountID       string `json:"account_id" bson:"account_id"`
	Amount          int    `json:"amount" bson:"amount"`
	TransactionType string `json:"transaction_type" bson:"transaction_type"`
	Provider        string `json:"provider" bson:"provider"`
	ProviderId      string `json:"provider_id" bson:"provider_id"`
	UpdatedAt       string `json:"updated_at" bson:"updated_at"`
}

func CreateTransaction(transaction Transactions) error {
	_, err := configs.MI.DB.Collection("transactions").InsertOne(context.TODO(), transaction)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransactionStatus(transactionID, status string) error {
	_, err := configs.MI.DB.Collection("transactions").UpdateOne(context.TODO(), map[string]string{"_id": transactionID}, map[string]string{"status": status})
	if err != nil {
		return err
	}
	return nil
}

func GetTransactionsByAccountId(accountID string) ([]Transactions, error) {
	var transactions []Transactions
	cursor, err := configs.MI.DB.Collection("transactions").Find(context.TODO(), map[string]string{"account_id": accountID})
	if err != nil {
		return transactions, err
	}
	if err = cursor.All(context.TODO(), &transactions); err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionsByUserId(userID string) ([]Transactions, error) {
	var transactions []Transactions
	cursor, err := configs.MI.DB.Collection("transactions").Find(context.TODO(), map[string]string{"user_id": userID})
	if err != nil {
		return transactions, err
	}
	if err = cursor.All(context.TODO(), &transactions); err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionById(transactionID string) (Transactions, error) {
	var transaction Transactions
	err := configs.MI.DB.Collection("transactions").FindOne(context.TODO(), map[string]string{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetTransactionByProviderId(providerId string) (Transactions, error) {
	var transaction Transactions
	err := configs.MI.DB.Collection("transactions").FindOne(context.TODO(), map[string]string{"provider_id": providerId}).Decode(&transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func AddProviderIdToTransaction(transactionID, providerId string) error {
	update := bson.M{"$set": bson.M{"provider_id": providerId}}
	_, err := configs.MI.DB.Collection("transactions").UpdateOne(context.TODO(), bson.M{"_id": transactionID}, update)
	if err != nil {
		return err
	}
	return nil
}
