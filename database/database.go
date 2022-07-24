package database

import (
	"database/sql"
	"fmt"
)

// SQL Database
var DBClient *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:devchua1995@tcp(localhost:3306)/godevdb?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr.Error())
	}
	fmt.Println("MySQL Info")
	fmt.Println(db)

	// Expose Database to Golang App
	DBClient = db
}
