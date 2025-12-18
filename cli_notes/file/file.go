package file

import (
	"encoding/json"
	"os"

	"github.com/alexcloudstar/go_build_cli_notes/notes"
)

func CreateFile(fileName string) error {
	err := os.WriteFile(fileName, []byte("[]"), 0644)

	if err != nil {
		return err
	}

	return nil
}

func WriteFile(fileName string, userNotes *[]notes.Note) error {
	notesJson, _ := json.Marshal(userNotes)


	err := os.WriteFile(fileName, notesJson, 0644)

	return err
}

func ReadFile(fileName string) ([]notes.Note, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var userNotes []notes.Note

	err = json.Unmarshal(data, &userNotes)

	if err != nil {
		return nil, err
	}

	return userNotes, nil
}
