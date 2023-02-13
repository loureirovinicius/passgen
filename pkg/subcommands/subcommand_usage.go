package subcommands

const GenerateUsage = `Usage of generate:
  -l, --length uint
	password length (default 8)
  -o, --output-file string
	path to file where you want to write the password

Examples:
  passgen generate 
  # Outputs a password with the default value (8).

  passgen generate -l 20
  # Outputs a password with 20 characters.

  passgen generate -o passwords.txt
  # Writes the password generated to "passwords.txt" file. It also follows the default value since no other was specified.

  passgen generate -l 20 -o passwords.txt
  # Writes the password of 20 characters to "passwords.txt" file.`

const EncryptUsage = `Usage of encrypt:
  -f, --raw-file string
	password file to be encrypted by this program
  -o, --output-file string
	encrypted password file path/name (command output) (default "passwords")
	
Examples:
  passgen encrypt -f passwords.txt -o passwords_encrypted
  # Encrypts the "passwords.txt" file and creates a new "passwords_encrypted.bin" file with the encrypted content.`

const DecryptUsage = `Usage of decrypt:
  -f, --encrypted-file string
    password file to be decrypted
  -k, --key-file string
	key file you will use to decrypt the encrypted file. It was generated in the encryption process
	
Examples:
  passgen decrypt -f passwords.bin -k key
  # Decrypts the "passwords.bin" file using the "key" file.`
