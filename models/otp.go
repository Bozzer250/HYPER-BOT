package models

import (
	"context"
	"hyperbot/configs"

	"go.mongodb.org/mongo-driver/bson"
)

type Otp struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
	Status    string `json:"status"`
	UserId    string `json:"user_id"`
}

// create otp
func CreateOtp(otp Otp) error {
	_, err := configs.MI.DB.Collection("otps").InsertOne(context.TODO(), otp)
	if err != nil {
		return err
	}
	return nil
}

// update otp status
func UpdateOtpStatus(id string, status string) error {
	_, err := configs.MI.DB.Collection("otps").UpdateOne(context.TODO(), bson.M{"ID": id}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return err
	}
	return nil
}

// get otps by phone
func GetOtpsByPhone(phone string) ([]Otp, error) {
	var otps []Otp
	cursor, err := configs.MI.DB.Collection("otps").Find(context.TODO(), bson.M{"phone": phone})
	if err != nil {
		return otps, err
	}
	err = cursor.All(context.Background(), &otps)
	if err != nil {
		return otps, err
	}
	return otps, nil
}

// get otp by phone and code
func GetOtpByPhoneAndCode(phone string, code string) (Otp, error) {
	var otp Otp
	err := configs.MI.DB.Collection("otps").FindOne(context.TODO(), bson.M{"phone": phone, "code": code}).Decode(&otp)
	if err != nil {
		return otp, err
	}
	return otp, nil
}
