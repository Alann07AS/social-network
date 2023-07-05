package user

import "time"

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

func Change(user_id uint64, initiator_token, field_name string, value interface{}) error {
	// check for initiator right
	if false { // if not the right
		return ERR_FORBIDEN_OPPERATION
	}

	switch field_name {
	case "id", "created_at", "updated_at":
		return ERR_FORBIDEN_OPPERATION
	case "email":

	case "password":

	case "first_name":

	case "last_name":

	case "date_of_birth":

	case "avatar_image_path":

	case "nickname":

	case "about_me":

	case "is_public_profile":

	default:
		return ERR_
	}
	return nil
}

func Get(user_id uint64, initiator_token, field_name string, value *interface{}) error {
	// check for initiator right
	if false { // if not the right
		return ERR_FORBIDEN_OPPERATION
	}
}
