package controllers

import (
	"fmt"
	"hyperbot/models"

	"github.com/google/uuid"
)

func CalculateTodayProfit(userId string) error {
	accounts, err := models.GetAccountsByUserID(userId)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		if account.Status == "active" {
			tuid := uuid.New().String()
			profit := account.DailyRate * account.Price
			transaction := models.Transactions{
				ID:              tuid,
				UserID:          userId,
				AccountID:       account.ID,
				Amount:          int(profit),
				TransactionType: "profit",
				Provider:        "hyperbot",
				Status:          "pending",
			}
			err := models.CreateTransaction(transaction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CronjobToCalculatodayProfit() {
	users, err := models.GetAllActiveUsers()
	if err != nil {
		fmt.Printf("\nError getting all active users: %v", err)
	}
	for _, user := range users {
		errr := CalculateTodayProfit(user.ID)
		fmt.Printf("\nError calculating profit for user id %s:\n %v", user.ID, errr)
	}
}
