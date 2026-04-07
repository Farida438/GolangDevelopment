package db_sql_pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=student password=securepass dbname=school sslmode=disable"

func RunTask3() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	fmt.Println("Connection successful")
}
