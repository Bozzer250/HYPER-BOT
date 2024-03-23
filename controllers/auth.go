package controllers

import (
	"fmt"
	"hyperbot/integrations"
	"hyperbot/models"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

func VerifyAndUpdateOtp(phone, code string) (string, error) {
	fmt.Printf("Phone: %s, Code: %s\n", phone, code)
	otp, err := models.GetOtpByPhoneAndCode(phone, code)
	if err != nil {
		return "", fmt.Errorf("could not find otp")
	}
	if otp.Code != code || otp.Status != "active" {
		return "", fmt.Errorf("invalid otp")
	}
	_ = models.UpdateOtpStatus(otp.ID, "expired")
	return otp.UserId, nil
}

func createAndSendOtp(UserId, phone string) error {
	uid := uuid.New().String()
	code := 1000 + rand.Intn(9000)
	otp := models.Otp{
		Phone:  phone,
		Code:   strconv.Itoa(code),
		Status: "active",
		ID:     uid,
		UserId: UserId,
	}
	err := models.CreateOtp(otp)
	if err != nil {
		return err
	}
	errr := integrations.SendOtpViaSMS(phone, otp.Code)
	if errr != nil {
		return errr
	}
	return nil
}

func createNewUser(phone string) error {
	uid := uuid.New().String()
	user := models.User{
		Phone:  phone,
		Status: "active",
		ID:     uid,
	}
	err := models.CreateUser(user)
	if err != nil {
		return err
	}
	return createAndSendOtp(uid, phone)
}

func HandleUserAuth(phone string) error {
	user, err := models.GetUserByPhone(phone)
	if err != nil {
		return createNewUser(phone)
	}
	if user.Status != "active" {
		return fmt.Errorf("user is not active")
	}
	return nil
}
