package vaultOperations

import (
	"errors"
	"fmt"
	"log"
	"os"
	"password-manager/internal/fileOperations"
	"password-manager/security/decryption"
	"strconv"
	"strings"

	"golang.org/x/term"
)

/**
 * Fetches all the records present in a vault.
 */
func FetchRecordsFromVault(recordsToFetch string, vaultName string, masterPassword []byte) {
	var credName string
	var credsFromVault []string

	if recordsToFetch == "all" {
		fmt.Println("Fetching all credentials from the vault...")
	} else {
		fmt.Println("Fetching the requested credential from vault...")
	}

	decodedFile, err := decryption.DecryptFile("vaults/"+vaultName+".csv", string(masterPassword))
	if err != nil {
		fmt.Println("Password doesn't match. Please try again.")
	}

	credsFromVault = strings.Split(decodedFile, ",")

	for i := 2; i < len(credsFromVault)-1; i = i + 2 {

		decodedUsername := decryption.DecryptPassword([]byte(credsFromVault[i]), string(masterPassword))
		decodedPassword := decryption.DecryptPassword([]byte(credsFromVault[i+1]), string(masterPassword))

		if recordsToFetch != "all" {

			fmt.Print("Enter cred to fetch: ")
			fmt.Scan(&credName)

			if credName == decodedUsername {
				fmt.Println("----------------------------------------------")
				fmt.Printf("Username: %s", decodedUsername)
				fmt.Println()
				fmt.Printf("Password: %s", decodedPassword)
				fmt.Println()
				break
			} else {
				fmt.Println("No cred found by this name. Try again.")
				i = i - 2
				continue
			}
		} else {
			fmt.Println("----------------------------------------------")
			fmt.Printf("Username: %s", decodedUsername)
			fmt.Println()
			fmt.Printf("Password: %s", decodedPassword)
			fmt.Println()
		}

	}
}

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
		fmt.Println("----------------------------------------------")
		fmt.Println("New vault created!!")
	}

	return true, nil
}

func DeleteVault(vaultName string) (bool, error) {

	// Create the vault file with the password.
	fileExists, err := fileOperations.FileExists(vaultName + ".csv")

	if !fileExists || err != nil {
		fmt.Println("No vault by this name.")
		return false, nil
	} else {
		fmt.Println("----------------------------------------------")
		fmt.Println("New vault created!!")
	}

	return true, nil
}

func ManageExistingVault(userProvidedVaultName string) (bool, error) {
	var selectedOption string
	var gettingInput string = "yes"
	var vaultName string
	var readingVaultName bool = true
	var readingMasterPassword, readingConfirmPassword bool = true, true
	var masterPass, confirmPass []byte
	fd := int(os.Stdin.Fd())

	fmt.Println("A vault with this name already exists. What would you like to do?")
	fmt.Println("1. Create a new vault with different name.")
	fmt.Println("2. Signin to \"" + userProvidedVaultName + "\" vault.")
	fmt.Println("3. Delete \"" + userProvidedVaultName + "\" vault")
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

					fmt.Println("Creating a new vault!")

					for readingVaultName {
						fmt.Print("Please provide a name for the vault: ")
						fmt.Scan(&vaultName)

						// If no name was given, return an error with a message.
						if vaultName != "" {
							readingVaultName = false
							break
						}

						fmt.Println("Invalid vault name. Please try again.")
						continue
					}

					dirExists := fileOperations.DirExists()

					if !dirExists {
						// Create the vaults directory with and set 744 permission.
						err = os.Mkdir("vaults/", os.FileMode(0744))
						if err != nil {
							log.Fatal(err)
						}
					}

					fileExists, err := fileOperations.FileExists(vaultName + ".csv")
					if err != nil {
						log.Fatalf("file exist error: %x", err)
					}

					if fileExists {
						continue
					}

					for readingMasterPassword {
						fmt.Print("Please provide a master password: ")
						masterPass, err = term.ReadPassword(fd)
						fmt.Println()
						if err != nil {
							panic(err)
						}
						if len(masterPass) < 8 {
							fmt.Println("Password shoud be minimum 8 characters long.")
							continue
						} else {
							readingMasterPassword = false
							break
						}
					}

					for readingConfirmPassword {
						fmt.Print("Please confirm the master password: ")
						confirmPass, err = term.ReadPassword(fd)
						fmt.Println()
						if err != nil {
							panic(err)
						}

						if string(masterPass) != string(confirmPass) {
							fmt.Println("Password doesn't match. Please enter the password again")
							continue
						} else {
							readingConfirmPassword = false
							break
						}
					}

					// If a vault is created vault will contain a boolean value
					// of true and error as nil else vault will be false and error
					// will contain error message.
					_, err = CreateVault(vaultName, string(masterPass))

					// If an error was returned, print it to the console and
					// exit the program.
					if err != nil {
						log.Fatalf("Vault creation error: %x", err)
					}
					fmt.Println("---------------------")

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
func SigninToVault() (string, string) {
	var vaultName string
	var readingMasterPassword bool = true
	var readingVaultName bool = true
	var goBack string = "no"
	var decodedFile string
	var credsFromVault []string
	var enteredMasterPass []byte
	var err error

	fd := int(os.Stdin.Fd())

	for readingVaultName {
		if goBack == "n" || goBack == "N" || goBack == "no" || goBack == "No" {
			fmt.Print("Enter vault name: ")
			fmt.Scan(&vaultName)

			if vaultName == "" {
				fmt.Println("Invalid vault name. Please try again.")
				continue
			}

			vaultFileExists, err := fileOperations.FileExists(vaultName + ".csv")
			if !vaultFileExists || err != nil {
				fmt.Println("No vault exists by this name. Please try again.")
				fmt.Print("Want to go back to main menu?[y/N] ")
				fmt.Scan(&goBack)
				if goBack == "y" || goBack == "Y" || goBack == "yes" || goBack == "Yes" {
					break
				} else if goBack == "n" || goBack == "N" || goBack == "no" || goBack == "No" {
					continue
				} else {
					fmt.Println("Invalid input, going back to main menu.")
					break
				}
			}
			break
		}
	}

	for readingMasterPassword {
		fmt.Print("Please enter master password: ")
		enteredMasterPass, err = term.ReadPassword(fd)
		fmt.Println()
		if err != nil {
			panic(err)
		}

		decodedFile, err = decryption.DecryptFile("vaults/"+vaultName+".csv", string(enteredMasterPass))
		if err != nil {
			fmt.Println("Password doesn't match. Please try again.")
			continue
		}

		credsFromVault = strings.Split(decodedFile, ",")

		decodedMasterPassword := decryption.DecryptPassword([]byte(credsFromVault[1]), string(enteredMasterPass))

		// if entered masterpass is equal to vault master pass continue else show the prompt.
		if decodedMasterPassword != string(enteredMasterPass) {
			fmt.Println("Password doesn't match. Please try again.")
			continue
		} else if len(enteredMasterPass) < 8 {
			fmt.Println("Password shoud be minimum 8 characters long.")
			continue
		} else {
			readingMasterPassword = false
			break
		}
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("Signed in successfully!!")

	return string(vaultName), string(enteredMasterPass)
}
