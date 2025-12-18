package main

import (
	"fmt"
	"os"

	"github.com/alexcloudstar/go_build_cli_notes/file"
	"github.com/alexcloudstar/go_build_cli_notes/notes"
)

const fileName = "notes.json"

func main() {
	var choice int64

	var userNotes []notes.Note
	nextId := 1

	err := file.CreateFile(fileName)

	if err != nil {
		fmt.Println("Error creating or opening file:", err)
		return
	}

	for {
		printOptions()
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
		}

		runChoice(choice, &userNotes, &nextId)

	}
}

func printOptions() {
	fmt.Println("Welcome! Let's choose an option:")
	fmt.Println("1. Add Note")
	fmt.Println("2. View Notes")
	fmt.Println("3. Update Note")
	fmt.Println("4. Delete Note")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice (1-5): ")
}

func runChoice(choice int64, userNotes *[]notes.Note, nextId *int) {
	switch choice {
	case 1:
		note := notes.AddNote(int64(*nextId))

		*userNotes = append(*userNotes, note)
		*nextId++

		err := file.WriteFile(fileName, userNotes)

		if err != nil {
			fmt.Println("Error saving notes to file:", err)
		}

	case 2:
		fileNotes, err := file.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error reading notes from file:", err)
			return
		}

		notes.ListNotes(fileNotes)

		
	case 3:
		fileNotes, err := file.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error reading notes from file:", err)
			return
		}

		newNotes := notes.UpdateNote(fileNotes)

		
		err = file.WriteFile(fileName, &newNotes)

		if err != nil {
			fmt.Println("Error saving notes to file:", err)
		}
	case 4:
		fileNotes, err := file.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error reading notes from file:", err)
			return
		}

		newNotes := notes.DeleteNote(&fileNotes)

		err = file.WriteFile(fileName, &newNotes)

		if err != nil {
			fmt.Println("Error saving notes to file:", err)
		}
	case 5:
		fmt.Println("Exiting the program. Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
	}
}
