package database // import "github.com/eriol/wp24-deities/database"

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// Open an SQLite3 database at the specified path.
func Open(path string) (err error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	database = db

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := createDatabase(); err != nil {
				return err
			}
		}

	}

	return nil
}

// Close SQLite3 databaser.
func Close() error {
	return database.Close()
}

func createDatabase() error {

	tables := `
    PRAGMA foreign_keys = ON;

    CREATE TABLE IF NOT EXISTS deities (
        deity_id TEXT NOT NULL PRIMARY KEY,
        name TEXT COLLATE NOCASE,
        description TEXT COLLATE NOCASE,
        gender TEXT
    );

    CREATE TABLE IF NOT EXISTS sports (
        sport_id TEXT NOT NULL PRIMARY KEY,
        name TEXT COLLATE NOCASE
    );

    CREATE TABLE IF NOT EXISTS olympian_influence (
        deity_id TEXT NOT NULL,
        sport_id TEXT NOT NULL,
        influence REAL,
        FOREIGN KEY(deity_id) REFERENCES deities(deity_id)
        FOREIGN KEY(sport_id) REFERENCES sports(sport_id)
    );
    `

	_, err := database.Exec(tables)

	if err != nil {
		return err
	}

	return nil
}
