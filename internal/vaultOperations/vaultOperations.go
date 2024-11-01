package vaultOperations

import (
	"errors"
	"fmt"
	"password-manager/internal/fileOperations"
	"strconv"
)

/**
 * Function to allow the user to crate a password vault.
 */
func CreateVault(vaultName string, masterPass string) (bool, error) {

	// Create the vault file with the password.
	file, err := fileOperations.CreateFile(vaultName, string(masterPass))

	if err != nil {
		return false, errors.New("file creation error")
	}

	if !file {
		fmt.Println("No new vault created!!")
		return false, nil
	} else {
		fmt.Println("New vault created!!")
	}

	return true, nil
}

func ManageExistingVault(vaultName string) (bool, error) {
	var selectedOption string
	var gettingInput string = "yes"

	fmt.Println("A vault with this name already exists. What would you like to do?")
	fmt.Println("1. Create a new vault with different name.")
	fmt.Println("2. Signin to \"" + vaultName + "\" vault.")
	fmt.Println("3. Delete \"" + vaultName + "\" vault")
	fmt.Println("4. Go back to initial options.")
	fmt.Println("Quit (enter q or quit or press \"ctrl+c\")")

	for gettingInput != "no" {
		if gettingInput == "yes" {

			fmt.Scan(&selectedOption)

			if selectedOption == "q" || selectedOption == "quit" {
				fmt.Println("Exiting")
				gettingInput = "no"
				return true, nil
			}
			intOption, err := strconv.ParseInt(selectedOption, 10, 32)
			if err != nil {
				fmt.Println("Please enter a valid option.")
				gettingInput = "yes"
				continue
			}

			if intOption <= 0 || intOption >= 5 {
				fmt.Println("Please enter a valid option.")
				gettingInput = "yes"
				continue
			} else {
				if intOption == 1 {

					// If a vault is created vault will contain a boolean value
					// of true and error as nil else vault will be false and error
					// will contain error message.
					// _, err := CreateVault(vaultName)

					// If an error was returned, print it to the console and
					// exit the program.
					// if err != nil {
					// 	return false, errors.New("error creating the vault")
					// }

					return true, nil
				} else if intOption == 2 {
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
					return true, nil

				} else if intOption == 3 {
					fmt.Println("Deleting vault...")
					// If a vault already exists, siging to the vault to view all
					// the contained creds.
					// _, err := vaultOperations.SigninToVault()
					return true, nil

				} else if intOption == 4 {
					fmt.Println("Going back...")
					// If a vault already exists, siging to the vault to view all
					// the contained creds.
					// _, err := vaultOperations.SigninToVault()
					return true, nil

				} else {
					gettingInput = "no"
					return false, errors.New("unexpected input")
				}
			}
		}
	}

	return true, nil
}

/**
 * Function to allow the user to signin to the vault.
 */
// func SigninToVault(vaultName, vaultMasterPassword string) (bool, error) {
// 	if vaultName == "" {
// 		return false, errors.New("empty vault name")
// 	}

// 	if enteredMasterPass != masterPass {
// 		return false, errors.New("passwords don't match")
// 	}
// }
