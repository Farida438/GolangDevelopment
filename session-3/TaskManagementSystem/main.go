package main

import "fmt"
import "errors"

func simulate() (string, error) {
	return "Project can't proceed ", errors.New("No tasks completed")
}

func main() {
	//TASK1
	projectName := "Task Management System"
	totalTasks := 120
	completedTasks := 115

	tasksRemained := totalTasks - completedTasks
	fmt.Println("Tasks remained:", tasksRemained)

	if completedTasks > 90 {
		fmt.Println("Project is near to completion")
	} else if completedTasks == tasksRemained {
		fmt.Println("Half of project has been done")
	} else if completedTasks == totalTasks {
		fmt.Println("Project is COMPLETE!")
	} else if completedTasks == 0 {
		fmt.Println("Project can't run")
	} else {
		fmt.Println("Project is running")
	}

	//TASK2

	projectStatus := "starting phase"

	if completedTasks > 30 && completedTasks <= 60 {
		projectStatus = "half of project is completed"
	} else if completedTasks > 60 && completedTasks <= 90 {
		projectStatus = "still running "

	} else if completedTasks > 90 && completedTasks < totalTasks {
		projectStatus = "approximately completed "
	} else {
		projectStatus = "completed"
	}
	fmt.Println(projectStatus)

	//- Add a `switch` statement that categorizes the project progress:
	//     - If less than 30 tasks are completed, mark the progress as "Starting phase".
	//     - If 30 to 60 tasks are completed, mark it as "Midway".
	//     - If more than 60 tasks are completed, mark it as "Final phase".

	switch true {
	case completedTasks < 30:
		progress := "Starting Phase"
		fmt.Println("Progress:", progress)
	case completedTasks >= 30 && completedTasks < 60:
		progress := "Midway"
		fmt.Println("Progress:", progress)
	case completedTasks > 60:
		progress := "Final Phase"
		fmt.Println("Progress:", progress)
	}

	//TASK3
	//	 Use a `for` loop to iterate tasks by number.
	//	 For simplicity, create a for loop and iterate the remaining number of tasks, and output with task names (e.g., `"Task 1"`, `"Task 2"`, etc.).
	//     totalTasks := 120
	//	completedTasks := 115

	for i := 0; i < (totalTasks - completedTasks); i++ {
		fmt.Println("Task", i+1)
	}

	//TASK4
	//	*Error Handling**:
	//   - Implement a simple error-checking mechanism: if the tasks completed exceeds the total number of tasks, print an error message and reset the tasks completed to the maximum number.
	//   - Simulate an error scenario where the project cannot proceed if no tasks are completed. Use a custom error and handle it gracefully.

	if completedTasks > totalTasks {
		err := errors.New("Tasks completed exceeds total tasks")
		if err != nil {
			fmt.Println(err)
			completedTasks = totalTasks
		}
	}

	if completedTasks == 0 {
		str, err := simulate()

		if err != nil {
			fmt.Println("Error:", err, str)
		}
	}

	fmt.Println(" Welcome to the Task Management System!\n"+
		" Project start date is: 2024-09-18 00:00:00 +0000 UTC\n"+
		"Project: ", projectName, "\n Tasks remaining", tasksRemained, "out of", totalTasks,
		"\n Project status is: ", projectStatus,
		"\nTask Priorities: 1-low, 2-Medium, 3-High",
		"\n Task List:",
		"\n Is the project completed?", completedTasks == totalTasks)

}
