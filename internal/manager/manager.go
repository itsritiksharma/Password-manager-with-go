package manager

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/hash"
	"password-manager/internal/vaultOperations"

	"golang.org/x/term"
)

func Manage() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vault: ")
	log.SetFlags(0)

	var vaultName, hashedMasterPassword, hashedConfirmPassword string

	fd := int(os.Stdin.Fd())

	fmt.Println("Creating a new vault!")

	fmt.Print("Please provide a name for the vault: ")
	fmt.Scan(&vaultName)

	fmt.Print("Please provide a master password: ")
	masterPass, err := term.ReadPassword(fd)
	fmt.Println()

	if err != nil {
		panic(err)
	}

	fmt.Print("Please confirm the master password: ")
	confirmPass, err := term.ReadPassword(fd)
	fmt.Println()

	if err != nil {
		panic(err)
	}

	hashedMasterPassword = hash.HashPassword(masterPass)
	hashedConfirmPassword = hash.HashPassword(confirmPass)

	// If a vault is created vault will contain a boolean value
	// of true and error as nil else vault will be false and error
	// will contain error message.
	vault, err := vaultOperations.CreateVault(vaultName, string(hashedMasterPassword), string(hashedConfirmPassword))

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
