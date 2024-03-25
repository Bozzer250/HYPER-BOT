package controllers

import (
	"fmt"
	"hyperbot/configs"
	"hyperbot/integrations"
	"hyperbot/models"

	"github.com/google/uuid"
)

func GetPackageDetails(packageName string) (string, int, float64) {
	switch packageName {
	case "EURUSD":
		return "eurusd", 20, 1
	case "NASDAQ":
		return "nasdaq", 50, 3
	case "GOLD":
		return "gold", 300, 8
	default:
		return "eurusd", 20, 1
	}
}

func PurchaseNewAccount(phone, referralCode, packageName string) error {
	User, err := models.GetUserByPhone(phone)
	if err != nil {
		// create new user
		return err
	}
	fmt.Printf("\nUser: %v and id is %s", User, User.ID)
	accountType, price, dailyRate := GetPackageDetails(packageName)
	// create new account
	uid := uuid.New().String()
	account := models.Account{
		ID:          uid,
		Status:      "pending",
		AccountType: accountType,
		DailyRate:   dailyRate,
		UserID:      User.ID,
	}
	_ = models.CreateAccount(account)
	tuid := uuid.New().String()
	amount := price * configs.GetCashinUsdRate()
	transaction := models.Transactions{
		ID:              tuid,
		Status:          "pending",
		UserID:          User.ID,
		AccountID:       account.ID,
		Amount:          amount,
		TransactionType: "cashin",
		Provider:        "paypack",
	}
	_ = models.CreateTransaction(transaction)
	// send transacion to paypack
	cashin, err := integrations.PaypackCashIn(amount, phone)
	if err != nil {
		fmt.Printf("Error cashing in: %v", err)
		return err
	}
	errr := models.AddProviderIdToTransaction(tuid, cashin.Ref)
	if errr != nil {
		fmt.Printf("Error adding provider id to transaction: %v", errr)
	}
	return nil
}
