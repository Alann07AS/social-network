package user

import (
	"database/sql"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

var database *sql.DB

func CRUD_Init(db *sql.DB) {
	database = db
}

func Create(Email, Password, First_name, Last_name, Date_of_birth, Is_public_profile string) error {
	// check mail validity
	if err := CheckMailValidity(Email); err != nil {
		return err
	}

	// check password Strength
	if err := CheckPassWordStrength(Email); err != nil {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = database.Exec(`
		INSERT into user (email, password, first_name, last_name, date_of_birth, is_public_profile)
		VALUES (?, ?, ?, ?, ?, ?)
	`, Email, string(pass), First_name, Last_name, Date_of_birth, Is_public_profile)
	if err != nil {
		return err
	}

	return nil
}

func Read(user_id uint64, initiator_token, field_name string, value *interface{}) error {
	// check for initiator right
	if false { // if not the right
		return ERR_FORBIDEN_OPPERATION
	}
	return nil
}

func Update(user_id uint64, initiator_token, field_name string, value interface{}) error {
	// check for initiator right
	/*
		Rajout des compte admin plus tard
	*/
	if false { // if not the right
		return ERR_FORBIDEN_OPPERATION
	}

	if !reflect.DeepEqual(reflect.TypeOf(value), reflect.String) {
		return ERR_BAD_VALUE_TYPE
	}

	switch field_name {
	case "id", "created_at", "updated_at":
		return ERR_FORBIDEN_OPPERATION
	case "email":
		if err := CheckMailValidity(value.(string)); err != nil {
			return err
		}

	case "password":
		if err := CheckPassWordStrength(value.(string)); err != nil {
			return err
		}

	case "first_name", "last_name":

	case "date_of_birth":

	case "avatar_image_path": // peutetre cree une methode particulieur

	case "nickname":

	case "about_me":

	case "is_public_profile":

	default:
		return NoFieldName(field_name)
	}
	return nil
}

func Delete() error {
	// check for initiator right
	/*
		Rajout des compte admin plus tard
	*/
	if false { // if not the right
		return ERR_FORBIDEN_OPPERATION
	}
	return nil
}
