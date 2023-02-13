package subcommands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"passgen/pkg/cryptography"
	"passgen/pkg/fileutils"
	"passgen/pkg/generation"
)

var (
	flagPasswordFile string
)

func Generate() {
	var flagLength uint
	var flagWriteToFile string

	scGenerate := flag.NewFlagSet("generate", flag.ExitOnError)
	scGenerate.UintVar(&flagLength, "l", 8, "password length")
	scGenerate.UintVar(&flagLength, "length", 8, "password length")
	scGenerate.StringVar(&flagWriteToFile, "o", "", "path to file where you want to write the password")
	scGenerate.StringVar(&flagWriteToFile, "output-file", "", "path to file where you want to write the password")

	scGenerate.Usage = func() { fmt.Println(GenerateUsage) }

	scGenerate.Parse(os.Args[2:])

	if flagLength <= 6 {
		fmt.Println("The password length must have at least 6 characters.")
		os.Exit(1)
	}

	if flagWriteToFile != "" {
		p := generation.GeneratePassword(flagLength)
		err := fileutils.WriteToFile(flagWriteToFile, p)
		if err != nil {
			log.Fatalf("there was an error writing the password to file: %v", err)
		}
		os.Exit(0)
	}
	fmt.Println(generation.GeneratePassword(flagLength))
}

func Encrypt() {
	var flagEncryptedFileName string

	scEncrypt := flag.NewFlagSet("encrypt", flag.ExitOnError)
	scEncrypt.StringVar(&flagPasswordFile, "f", "", "password file to be encrypted")
	scEncrypt.StringVar(&flagPasswordFile, "raw-file", "", "password file to be encrypted")
	scEncrypt.StringVar(&flagEncryptedFileName, "o", "passwords", "encrypted password file path/name (command output)")
	scEncrypt.StringVar(&flagEncryptedFileName, "output-file", "passwords", "encrypted password file path/name (command output)")

	scEncrypt.Usage = func() { fmt.Println(EncryptUsage) }

	scEncrypt.Parse(os.Args[2:])
	if flagPasswordFile != "" {
		err := cryptography.EncryptFile(flagPasswordFile, flagEncryptedFileName)
		if err != nil {
			log.Fatalf("there was an error encrypting the file: %v", err)
		}
	} else {
		fmt.Println("You must insert the password file you want to encrypt.")
		os.Exit(1)
	}
}

func Decrypt() {
	var flagKeyFile string

	scDecrypt := flag.NewFlagSet("decrypt", flag.ExitOnError)
	scDecrypt.StringVar(&flagPasswordFile, "f", "", "password file to be decrypted")
	scDecrypt.StringVar(&flagPasswordFile, "encrypted-file", "", "password file to be decrypted")
	scDecrypt.StringVar(&flagKeyFile, "k", "", "key file you will use to decrypt the encrypted file. It was generated in the encryption process")
	scDecrypt.StringVar(&flagKeyFile, "key-file", "", "key file you will use to decrypt the encrypted file. It was generated in the encryption process")

	scDecrypt.Usage = func() { fmt.Println(DecryptUsage) }

	scDecrypt.Parse(os.Args[2:])
	if flagPasswordFile != "" && flagKeyFile != "" {
		err := cryptography.DecryptFile(flagPasswordFile, flagKeyFile)
		if err != nil {
			log.Fatalf("there was an error decrypting the file: %v", err)
		}
	} else {
		fmt.Println("You must insert the password file you want to decrypt and the key file used to encrypt it.")
		os.Exit(1)
	}
}
