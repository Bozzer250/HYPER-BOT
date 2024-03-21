package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	Status    int    `json:"status"`
}
