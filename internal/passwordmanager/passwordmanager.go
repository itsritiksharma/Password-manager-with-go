package passwordmanager

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/consoleprompts"
	"password-manager/internal/vaultOperations"
	"password-manager/security/decryption"

	"golang.org/x/term"
)

func promptUser() {
	var prompt string = "prompt"
	fmt.Println("Welcome to password manager.")
	for prompt != "do-not-prompt" {
		if prompt == "prompt" {

			userSelectedOption := consoleprompts.PromptUserInput()

			if userSelectedOption == 1 {

				// If a vault is created vault will contain a boolean value
				// of true and error as nil else vault will be false and error
				// will contain error message.
				_, err := vaultOperations.CreateVault()

				// If an error was returned, print it to the console and
				// exit the program.
				if err != nil {
					log.Fatalf("Vault creation error: %x", err)
				}
				fmt.Println("---------------------")
				continue
			} else if userSelectedOption == 2 {
				fmt.Println("Signing in to vault...")
				// fmt.Print("Please provide a name for the vault: ")
				// fmt.Scan(&vaultName)

				// for readingMasterPassword != 0 {
				// 	if readingMasterPassword == 1 {
				// 		fmt.Print("Please provide a master password: ")
				// 		masterPass, err = term.ReadPassword(fd)
				// 		fmt.Println()
				// 		if err != nil {
				// 			panic(err)
				// 		}
				// 		if len(masterPass) < 8 {
				// 			fmt.Println("Password shoud be minimum 8 characters long.")
				// 			readingMasterPassword = 1
				// 		} else {
				// 			readingMasterPassword = 0
				// 		}
				// 	}
				// }
				// If a vault already exists, siging to the vault to view all
				// the contained creds.
				// _, err := vaultOperations.SigninToVault()
				continue

			} else if userSelectedOption == 3 {
				fmt.Println("Deleting vault...")
				// If a vault already exists, siging to the vault to view all
				// the contained creds.
				// _, err := vaultOperations.SigninToVault()
				continue

			} else if userSelectedOption == 4 {
				var vaultName string
				fd := int(os.Stdin.Fd())

				fmt.Println("Fetching all credentials from the vault...")
				fmt.Print("Enter vault name: ")
				fmt.Scan(&vaultName)
				fmt.Print("Please the master password: ")
				masterPass, err := term.ReadPassword(fd)
				fmt.Println()
				if err != nil {
					panic(err)
				}
				// hashedMasterPassword := encryption.EncryptPassword([]byte(masterPass), string(masterPass))

				decodedFile := decryption.DecryptFile("vaults/"+vaultName+".csv", string(masterPass))

				fmt.Println(decodedFile)
				continue

			} else if userSelectedOption == 5 {
				fmt.Println("Fetching the requested credential from vault...")
				// If a vault already exists, siging to the vault to view all
				// the contained creds.
				// _, err := vaultOperations.SigninToVault()
				continue

			} else if userSelectedOption == 6 {

				fd := int(os.Stdin.Fd())

				fmt.Println("Fetching all credentials from the JSON file...")
				fmt.Print("Please the master password: ")
				masterPass, err := term.ReadPassword(fd)
				fmt.Println()
				if err != nil {
					panic(err)
				}
				// hashedMasterPassword := encryption.EncryptPassword([]byte(masterPass), string(masterPass))

				decodedFile := decryption.DecryptFile("vaults/VaultsInfo.json", string(masterPass))
				fmt.Println(decodedFile)
				continue

			} else {
				prompt = "do-not-prompt"
				return
			}
		}
	}
}

func Manage() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vault: ")
	log.SetFlags(0)

	promptUser()

}
