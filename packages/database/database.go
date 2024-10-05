package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func GetDatabase() *sql.DB {
	d, err := sql.Open("sqlite", "./database/main.sqlite")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db")
		os.Exit(1)
	}
	return d
}
