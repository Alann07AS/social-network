package main

import (
	"social/db"
	sn "social/s-n"
)

func main() {
	db.DbInit()
	sn.StartAppServeur()
	// log.Fatal(user.Create("hello@hotmail.dbile", "1234", "alann", "asa", "10.12.2001", "public"))
}
