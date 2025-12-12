package main

import (
	"fmt"
	"os"
	"strings"
)

type Todo map[string]bool

func main() {
	args := os.Args[1:]
	
	cmd := args[0]

	var todoTitle string

	if len(args) > 1 {
		todoTitle = strings.Join(args[1:], " ")
	}

	todos := make([]Todo, 2)


	switch cmd {
	case "add":
		todos = append(todos, Todo{todoTitle: false})
		fmt.Printf("Added todo: %s\n", todoTitle)
	case "list":
		for i, todo := range todos {
			for task, done := range todo {
				status := "not done"
				if done {
					status = "done"
				}
				fmt.Printf("%d: %s - %s\n", i, task, status)
			}
		}
	default:
		fmt.Println("Unknown command")
	}
}
