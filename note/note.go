package note

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Note struct {
	name     string
	text     string
	Location string
}

func (n *Note) CreateNote() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write to Note")

	for scanner.Scan() {
		n.text = scanner.Text()
		if scanner.Text() == '\n' {
			return
		}
	}

	fmt.Println("What to name this note?")
	scanner2 := bufio.NewScanner(os.Stdin)

	for scanner2.Scan() {
		n.name = scanner.Text()
	}

	dirs, err := os.ReadDir(n.Location)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range dirs {
		if val.Name() != n.name {
			file, err := os.Create(filepath.Join(n.Location, n.name))
			if err != nil {
				log.Fatal(err)
			}

			_, err = file.WriteString(n.text)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("The file already exists. Either delete it or rename it.")
		}
	}

}
