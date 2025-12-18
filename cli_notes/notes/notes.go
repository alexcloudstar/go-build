package notes

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	ID      int64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func newNote (title, content string) (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter note title: ")

	if scanner.Scan() {
		title = scanner.Text()
	}

	fmt.Print("Enter note content: ")
	if scanner.Scan() {
		content = scanner.Text()
	}

	return title, content
}

func AddNote(nextId int64) Note {
	var title string
	var content string

	title, content = newNote(title, content)


	fmt.Println("Note added successfully.")

	return Note{ID: int64(nextId), Title: title, Content: content}
}

func ListNotes(userNotes []Note) {
	fmt.Println("Notes:")
	for _, note := range userNotes {
		fmt.Printf("ID: %d, Title: %s, \nContent: %s\n", note.ID, note.Title, note.Content)
	}
}

func UpdateNote(userNotes []Note) []Note {
	fmt.Println("What note would you like to update? (Enter ID)")
	var id int64
	fmt.Scanln(&id)

	for _, note := range userNotes {
		if note.ID == id {
			var newTitle string
			var newContent string

			newTitle, newContent = newNote(newTitle, newContent)

			userNotes[note.ID-1].Title = newTitle
			userNotes[note.ID-1].Content = newContent

			fmt.Println("Note updated successfully.")
		}
	}

	return userNotes
}

func DeleteNote(userNotes *[]Note) []Note {
	fmt.Println("What note would you like to delete? (Enter ID)")
	var id int64
	fmt.Scanln(&id)

	for i, note := range *userNotes {
		if note.ID == id {
			*userNotes = append((*userNotes)[:i], (*userNotes)[i+1:]...)
			fmt.Println("Note deleted successfully.")
			break
		}
	}

	return *userNotes
}
