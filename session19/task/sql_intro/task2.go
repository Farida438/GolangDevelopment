package sql_intro

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=student password=securepass dbname=school sslmode=disable"

func RunTask2() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Insert rows
	_, err = db.Exec(`
	INSERT INTO students (name, age)
	VALUES 
	('Ali', 20),
	('Geray', 22),
	('Medina', 21)
	`)
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	// Query rows
	rows, err := db.Query("SELECT id, name, age FROM students")
	if err != nil {
		log.Fatal("Query failed:", err)
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
