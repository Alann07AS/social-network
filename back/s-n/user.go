package sn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"social/db/user"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Email             string
		Password          string
		First_name        string
		Last_name         string
		Date_of_birth     string
		Is_public_profile string
	}{}

	raw_data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("read error: " + err.Error()))
		return
	}

	err = json.Unmarshal(raw_data, &data)
	if err != nil {
		w.Write([]byte("unmarshal error: " + err.Error()))
		return
	}

	err = user.Create(data.Email, data.Password, data.First_name, data.Last_name, data.Date_of_birth, data.Is_public_profile)
	if err != nil {
		w.Write([]byte("create user error: " + err.Error()))
		return
	}
}
