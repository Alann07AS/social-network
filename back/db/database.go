package db

import (
	"database/sql"
	"fmt"
	"log"

	"social/db/user"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func DbInit() {
	// database connection string
	dbConnStr := "db/database.db"

	// Create a new SQLite database connection
	var err error
	database, err = sql.Open("sqlite3", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// defer database.Close()

	// Create a new instance of the SQLite database driver
	driver, err := sqlite3.WithInstance(database, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create SQLite database driver instance: %v", err)
	}

	// Create a new migrate instance with the SQLite driver
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	// Apply all available migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully!")
	// Perform other database operations...

	// init all sub package who need to acces database
	user.CRUD_Init(database) // user
}
