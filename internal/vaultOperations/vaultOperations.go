package vaultOperations

import (
	"errors"
	"fmt"
	"log"
	"os"
	"password-manager/internal/fileOperations"
	"password-manager/security/encryption"
	"strconv"

	"golang.org/x/term"
)

/**
 * Function to allow the user to crate a password vault.
 */
func CreateVault() (bool, error) {

	var vaultName string
	var readingMasterPassword, readingConfirmPassword int = 1, 1
	var masterPass, confirmPass []byte
	var err error

	fd := int(os.Stdin.Fd())

	fmt.Println("Creating a new vault!")

	fmt.Print("Please provide a name for the vault: ")
	fmt.Scan(&vaultName)

	// If no name was given, return an error with a message.
	if vaultName == "" {
		return false, errors.New("empty vault name")
	}

	dirExists := fileOperations.DirExists()

	if !dirExists {
		// Create the vaults directory with and set 744 permission.
		err = os.Mkdir("vaults/", os.FileMode(0744))
		if err != nil {
			log.Fatal(err)
			return false, errors.New("failed to create the vault directory.")
		}
	}

	fileExists, err := fileOperations.FileExists(vaultName + ".csv")
	if err != nil {
		log.Fatalf("file exist error: %x", err)
	}

	if fileExists {
		// call another function which handles user input to determine
		// what to do next.
		status, err := ManageExistingVault(vaultName)
		return status, err
	}

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

	hashedMasterPassword := encryption.EncryptPassword([]byte(masterPass))

	// Create the vault file with the password.
	file, err := fileOperations.CreateFile(vaultName, hashedMasterPassword)

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
					_, err := CreateVault()

					// If an error was returned, print it to the console and
					// exit the program.
					if err != nil {
						return false, errors.New("error creating the vault")
					}

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
