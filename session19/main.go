package main

import (
	"session19/advanced"
	"session19/task/db_sql_pkg"
	"session19/task/sql_intro"
)

func main() {

	sql_intro.RunTask1()
	sql_intro.RunTask2()

	db_sql_pkg.RunTask3()
	db_sql_pkg.RunTask4()

	advanced.RunTask5()
	advanced.RunTask6()
}
