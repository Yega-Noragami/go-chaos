package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupTable() {
	// Open conn to DB
	db, err := sql.Open("mysql", getDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Prepare `CREATE TABLE` statement
	stmt, err := db.Prepare("CREATE TABLE `lists`(temp_key varchar(255),temp_value varchar(255),created_at timestamp,deleted_at timestamp,updated_at timestamp);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// Execute prepared statement
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
	}
}
