package controllers

import "fmt"

func PurchaseNewAccount(phone, referralCode string) error {
	fmt.Printf("Phone: %s, Referral Code: %s\n", phone, referralCode)
	return nil
}
