package main

import (
	"fmt"
	"log"
	"password-manager/vaultCreation"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("vault: ")
	log.SetFlags(0)

	var vaultName, masterPass, confirmPass string

	fmt.Println("Creating a new vault!")

	fmt.Print("Please provide a name for the vault: ")
	fmt.Scan(&vaultName)

	fmt.Print("Please provide a master password: ")
	fmt.Scan(&masterPass)

	fmt.Print("Please confirm the master password: ")
	fmt.Scan(&confirmPass)

	// If a vault is created vault will contain a boolean value
	// of true and error as nil else vault will be false and error
	// will contain error message.
	vault, err := vaultCreation.CreateVault(vaultName, masterPass, confirmPass)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If vault print the valut cration confirmation message.
	if vault {
		fmt.Printf("New vault create and saved as %s.csv", vaultName)
	} else {
		fmt.Printf("No new vault created.")
	}
}
