package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	todo string
	id   string
}

var Tasks []Task

func createTask(task Task) {
	Tasks = append(Tasks, task)
}
func viewTasks() {
	for _, task := range Tasks {
		fmt.Printf("Task: %vID: %v\n", task.todo, task.id)
	}
}
func deleteTask(Id string) {
	for index, task := range Tasks {
		if task.id == Id {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
		}
	}
}

func main() {
	fmt.Print("**************************************\n Weclome to your daily task manager! \n**************************************\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Choose one of the following options:-\n1:- Create\n2:- View\n3:- Update\n4:- Delete\n5:- Quit\n")
		userChoice, err := reader.ReadString('\n')
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))
		if err != nil {
			fmt.Println("An Unexpected Error has occured")
			break
		}
		if userChoice == "quit" {
			fmt.Println("Exiting the Program!")
			break
		} else if userChoice == "create" {
			fmt.Print("Ente the Task Description: ")
			var task Task
			taskreader := bufio.NewReader(os.Stdin)
			ToDo, _ := taskreader.ReadString('\n')
			task.todo = ToDo
			task.id = strconv.Itoa(len(Tasks))
			createTask(task)
			fmt.Println("The Task has been added succesfully!")

		} else if userChoice == "view" {
			viewTasks()

		} else if userChoice == "update" {

		} else if userChoice == "delete" {
			idreader := bufio.NewReader(os.Stdin)
			Id, _ := idreader.ReadString('\n')
			deleteTask(Id)
			fmt.Println("The Task has been Deleted succesfully!")
		} else {

		}
	}
}
