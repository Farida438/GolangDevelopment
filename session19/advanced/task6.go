package advanced

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func RunTask6() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("UPDATE students SET age=$1 WHERE name=$2", 21, "Ali")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction successful")
}
