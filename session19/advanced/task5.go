package advanced

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=student password=securepass dbname=school sslmode=disable"

func RunTask5() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO students(name, age) VALUES($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Leyla", 23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert successful")
}
