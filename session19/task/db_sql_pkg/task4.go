package db_sql_pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func RunTask4() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
