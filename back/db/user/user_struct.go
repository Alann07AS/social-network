package user

import (
	"time"
)

type User struct {
	ID                uint64
	Email             string
	Password          string
	First_name        string
	Last_name         string
	Date_of_birth     time.Time
	Avatar_image_path string
	Nickname          string
	About_me          string
	Is_public_profile bool
	Created_at        time.Time
	Updated_at        time.Time
}
