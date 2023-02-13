package fileutils

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, text string) error {
	var err error

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("the following error happened when opening the file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(text + string('\n'))
	if err != nil {
		return fmt.Errorf("an error occured when writing the password to a file: %v", err)
	}

	return nil
}
