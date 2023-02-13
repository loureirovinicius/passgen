package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"strings"
)

func EncryptFile(passwordFile string, encryptedFileName string) error {
	pf, err := os.ReadFile(passwordFile)
	if err != nil {
		return fmt.Errorf("there was an error reading the file containing the passwords: %v", err)
	}

	key := make([]byte, 32)
	err = os.WriteFile("key", key, 0600)
	if err != nil {
		return fmt.Errorf("there was an error writing your key to a file: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("the following error stopped the generation of your block algorithm: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("the following error happened when generating the gcm: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("there was an error reading the nonce array: %v", err)
	}

	e := path.Ext(encryptedFileName)
	if e != "" {
		encryptedFileName = strings.TrimSuffix(encryptedFileName, e)
	}

	ciphertext := gcm.Seal(nonce, nonce, pf, nil)
	if err := os.WriteFile(encryptedFileName+".bin", ciphertext, 0644); err != nil {
		return fmt.Errorf("there was an error writing the key to the file: %v", err)
	}

	if err := os.Remove(passwordFile); err != nil {
		return fmt.Errorf("error removing password file: %v", err)
	}

	return nil
}

func DecryptFile(encryptedFile string, keyFile string) error {
	ciphertext, err := os.ReadFile(encryptedFile)
	if err != nil {
		return fmt.Errorf("there was an error reading the password file to decrypt: %v", err)
	}

	key, err := os.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("there was an error reading your key to a file: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("the following error stopped the generation of your block algorithm: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("the following error happened when generating the gcm: %v", err)
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("there was an error opening the encrypted file: %v", err)
	}

	decryptedFile := "passwords_decrypted.txt"
	err = os.WriteFile(decryptedFile, plaintext, 0600)
	if err != nil {
		return fmt.Errorf("there was an error writing the decrypted value to a file: %v", err)
	}

	return nil
}
