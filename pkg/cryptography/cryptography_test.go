package cryptography_test

import (
	"os"
	"passgen/pkg/cryptography"
	"testing"
)

func TestEncryptFile(t *testing.T) {
	err := os.WriteFile("passwords_test.txt", []byte("123456789test"), 0600)
	if err != nil {
		t.Fatalf("failed to create password file: %v", err)
	}

	cryptography.EncryptFile("passwords_test.txt", "encrypted_test.bin")

	if _, err := os.Stat("passwords_test.txt"); !os.IsNotExist(err) {
		t.Error("original password file wasn't deleted")
	}

	if _, err := os.Stat("encrypted_test.bin"); os.IsNotExist(err) {
		t.Error("encrypted password file doesn't exist")
	}

	if _, err := os.Stat("key"); os.IsNotExist(err) {
		t.Error("key file doesn't exist")
	}

	if c, _ := os.ReadFile("encrypted_test.bin"); string(c) == "" {
		t.Error("encrypted password file is empty")
	}

	if c, _ := os.ReadFile("key"); string(c) == "" {
		t.Error("key file is empty")
	}
}

func TestDecryptFile(t *testing.T) {
	if _, err := os.Stat("encrypted_test.bin"); os.IsNotExist(err) {
		t.Error("encrypted file doesn't exist")
	}

	if _, err := os.Stat("key"); os.IsNotExist(err) {
		t.Error("key file doesn't exist")
	}

	cryptography.DecryptFile("encrypted_test.bin", "key")

	if _, err := os.Stat("passwords_decrypted.txt"); os.IsNotExist(err) {
		t.Error("file with the decrypted password wasn't created")
	}

	if c, _ := os.ReadFile("passwords_decrypted.txt"); string(c) == "" {
		t.Error("decrypted password file is empty")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Remove("encrypted_test.bin")
	os.Remove("key")
	os.Remove("passwords_decrypted.txt")

	os.Exit(code)
}
