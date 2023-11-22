package sn

import (
	"net/http"
)

/*
	Serveur API which return Error value or Data value

	{
		Error: "The error value"
	}


	{
		datas: [{All data value}]
	}
*/

func StartAppServeur() {
	http.HandleFunc("/users/register", createUser)
	http.ListenAndServe(":8080", nil)
}
