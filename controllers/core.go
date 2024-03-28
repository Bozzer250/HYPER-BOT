package controllers

import (
	"fmt"
	"hyperbot/configs"
	"hyperbot/integrations"
	"hyperbot/models"
	"time"

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
	accountType, price, dailyRate := GetPackageDetails(packageName)
	// create new account
	uid := uuid.New().String()
	account := models.Account{
		ID:           uid,
		Status:       "pending",
		AccountType:  accountType,
		DailyRate:    dailyRate,
		UserID:       User.ID,
		Price:        float64(price),
		ReferralCode: referralCode,
		CreatedAt:    time.Now(),
	}
	_ = models.CreateAccount(account)
	tuid := uuid.New().String()
	amount := price * configs.GetCashinUsdRate()
	transaction := models.Transactions{
		ID:              tuid,
		Status:          "pending",
		UserID:          User.ID,
		AccountID:       account.ID,
		Amount:          float64(amount),
		TransactionType: "cashin",
		Provider:        "paypack",
		CreatedAt:       time.Now(),
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

func LookupAllPendingPaypackTransactions() error {
	transactions, err := models.GetPendingTransactionsByProvider("paypack")
	if err != nil {
		return err
	}
	for _, transaction := range transactions {
		cashin, err := integrations.PollTransactionStatus(transaction.ProviderId)
		if err != nil {
			fmt.Printf("Error getting paypack transaction: %v", err)
			continue
		}
		if cashin.Status == "success" {
			_ = models.UpdateTransactionStatus(transaction.ID, "success")
			_ = activateAccount(transaction.AccountID)
		}
		if cashin.Status == "failed" || cashin.Status == "expired" || cashin.Status == "cancelled" {
			_ = models.UpdateTransactionStatus(transaction.ID, "failed")
			_ = models.UpdateAccountStatus(transaction.AccountID, "failed")
		}
	}
	return nil
}

func activateAccount(accountId string) error {
	account, err := models.GetAccountByID(accountId)
	if err != nil {
		return err
	}
	if account.Status == "active" {
		return fmt.Errorf("account is already active")
	}
	errr := models.UpdateAccountStatus(accountId, "active")
	if errr != nil {
		return errr
	}
	if len(account.ReferralCode) > 3 {
		userId, _ := models.GetUserIdFromReferralCode(account.ReferralCode)
		transaction := models.Transactions{
			ID:              uuid.New().String(),
			Status:          "completed",
			UserID:          userId,
			Amount:          account.Price * 0.1,
			TransactionType: "referral",
			CreatedAt:       time.Now(),
		}
		_ = models.CreateTransaction(transaction)
	}

	return nil
}
