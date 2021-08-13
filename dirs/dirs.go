package dirs

import (
	"fmt"
	"os"
	"path/filepath"
)

// CheckHome directory checks if a directory exists in the users home directory and the create parameter will determine if you
// want to or don't want to create the directory if it does not exist.
func CheckHome(directory string, create bool) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err.Error()
	}

	if _, err := os.Stat(filepath.Join(homedir, directory)); os.IsNotExist(err) {
		if create {
			fmt.Println("The Directory Does Not Exist.")
			fmt.Println("Creating One")
			err := os.Mkdir(filepath.Join(homedir, directory), 0755)
			fmt.Println("Done.")

			if err == nil {
				return filepath.Join(homedir, directory)
			} else {
				return ""
			}

		} else {
			return ""
		}
	}
	return err.Error()
}
