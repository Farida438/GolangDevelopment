package sql_intro

import "fmt"

func RunTask1() {
	fmt.Println("Docker command to create PostgreSQL container:")
	fmt.Println("docker run --name school-db -e POSTGRES_USER=student -e POSTGRES_PASSWORD=securepass -p 5432:5432 -d postgres")
}

//these commands were run in terminal and container was created
