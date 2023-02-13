package fileutils_test

import (
	"os"
	"passgen/pkg/fileutils"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	err := os.WriteFile("writeToFile_test.txt", []byte("new test file"), 0600)
	if err != nil {
		t.Fatalf("failed to create file to write: %v", err)
	}

	err = fileutils.WriteToFile("writeToFile_test.txt", "new text created")
	if err != nil {
		t.Errorf("error when writing text to file: %v", err)
	}

	err = os.Remove("writeToFile_test.txt")
	if err != nil {
		t.Fatalf("failed to remove file used by tests: %v", err)
	}
}
