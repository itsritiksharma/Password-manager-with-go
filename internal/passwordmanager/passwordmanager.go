package passwordmanager

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/vaultOperations"

	"golang.org/x/term"
)

func Manage() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vault: ")
	log.SetFlags(0)

	var vaultName string
	var readingMasterPassword, readingConfirmPassword int = 1, 1
	var masterPass, confirmPass []byte
	var err error

	fd := int(os.Stdin.Fd())

	fmt.Println("Creating a new vault!")

	fmt.Print("Please provide a name for the vault: ")
	fmt.Scan(&vaultName)

	for readingMasterPassword != 0 {
		if readingMasterPassword == 1 {
			fmt.Print("Please provide a master password: ")
			masterPass, err = term.ReadPassword(fd)
			fmt.Println()
			if err != nil {
				panic(err)
			}
			if len(masterPass) < 8 {
				fmt.Println("Password shoud be minimum 8 characters long.")
				readingMasterPassword = 1
			} else {
				readingMasterPassword = 0
			}
		}
	}

	for readingConfirmPassword != 0 {
		if readingConfirmPassword == 1 {
			fmt.Print("Please confirm the master password: ")
			confirmPass, err = term.ReadPassword(fd)
			fmt.Println()
			if err != nil {
				panic(err)
			}

			if string(masterPass) != string(confirmPass) {
				fmt.Println("Password doesn't match. Please enter the password again")
				readingMasterPassword = 1
			} else {
				readingConfirmPassword = 0
			}
		}
	}

	// If a vault is created vault will contain a boolean value
	// of true and error as nil else vault will be false and error
	// will contain error message.
	vault, err := vaultOperations.CreateVault(vaultName, string(masterPass))

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If vault print the valut cration confirmation message.
	if vault {
		fmt.Printf("New vault create and saved as %s.csv", vaultName)
		fmt.Println()
	} else {
		fmt.Printf("No new vault created.")
		fmt.Println()

	}
}
