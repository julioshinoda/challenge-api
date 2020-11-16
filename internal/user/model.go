package user

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
}
