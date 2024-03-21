package models

type Otp struct {
	ID        int    `json:"id"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
	Status    string `json:"status"`
}
