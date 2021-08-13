package main

import (
	"fmt"

	"github.com/D7682/notes/dirs"
	"github.com/D7682/notes/note"
)

func main() {
	n := &note.Note{}
	n.Location = dirs.CheckHome("notes", true)
	// dirs.CheckHome() returns a string to the new directories location, so if the value is nil then that means no directory was created or located.
	if n.Location == "" {
		fmt.Println("There was no location created")
		return
	}

	n.CreateNote()
}
