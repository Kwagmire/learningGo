package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type MyToDo map[string]*Task

type Task struct {
	description string
	status      string
	createdAt   time.Time
}

func getInput(s string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(s)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading input: ", err)
		return input, err
	}

	input = strings.TrimSpace(input)
	return input, err
}

func (list *MyToDo) addTask() {
	input, err := getInput("Enter a description of this task: ")
	if err != nil {
		//return _, err
		return
	}

	task := Task{}
	task.description = input
	task.status = "pending"
	task.createdAt = time.Now()

	id := len(*list)
	idString := strconv.Itoa(id)
	(*list)[idString] = &task
	//return &Task, nil
}

func (list *MyToDo) delTask() bool {
	input, err := getInput("Task ID?: ")
	if err != nil {
		return false
	}

	if _, ok := (*list)[input]; !ok {
		fmt.Println("Task doesn't exist")
		return false
	}
	delete(*list, input)
	return true
}

func (list *MyToDo) completeTask() bool {
	input, err := getInput("Task ID?: ")
	if err != nil {
		return false
	}

	if value, ok := (*list)[input]; !ok {
		fmt.Println("Task doesn't exist")
		return false
	} else if value.status == "complete" {
		fmt.Println("Task already completed")
		return false
	}

	(*list)[input].status = "complete"
	return true
}

func (list *MyToDo) viewList() {
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "ID", "Description", "Status", "Created")
	for k, v := range *list {
		fmt.Printf("%-20s%-20s%-20s%-20s\n", k, v.description, v.status, v.createdAt.Format(time.UnixDate))
	}
}

func main() {
	newList := make(MyToDo)

	for {
		input, err := getInput("Enter option[1 - add task, 2 - delete task, 3 - complete task, 4 - view list, 5 - end program]: ")
		if err != nil {
			continue
		}

		switch input {
		case "1":
			newList.addTask()
		case "2":
			newList.delTask()
		case "3":
			newList.completeTask()
		case "4":
			newList.viewList()
		case "5":
			return
		default:
			fmt.Println("Invalid option! Valid options are integers are 1 - 5")
		}
	}
}
