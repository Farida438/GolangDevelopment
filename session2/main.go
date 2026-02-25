package main

import "fmt"
import "strconv"

const (
	B = iota * 1024
	KB
	MB
	GB
)

//After searching on internet I came to using strconv package for conversion of bool into string

func main() {
	var projectStatus2 string = "in progress"

	var nameofPR1 string = "Program Debugging Center"
	var prcompleted bool = false
	var change string = strconv.FormatBool(prcompleted)

	var tasksProj1 int32 = 20
	var tasksProj2 int32 = 20
	var tasksProj3 int32 = 89

	const availabletasks int64 = 165

	const low int64 = 0
	const medium int64 = 1
	const high int64 = 2

	fmt.Printf("Welcome to %s\n Here we have totally %d tasks\n"+
		" %d  is completed out of %d tasks\n"+
		" Status of this project: %s\n"+
		"Task priorities : low - %d, medium - %d, high - %d\n"+
		"Is project completed? %s %d",
		nameofPR1, availabletasks, tasksProj1+tasksProj2+tasksProj3, availabletasks,
		projectStatus2, low, medium, high, change, MB,
	)
}
