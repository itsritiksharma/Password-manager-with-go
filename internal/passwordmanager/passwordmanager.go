package passwordmanager

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/consoleprompts"
	"password-manager/internal/fileOperations"
	"password-manager/internal/vaultOperations"

	"golang.org/x/term"
)

func promptUser() {

	var vaultName string
	var prompt, readingVaultName, signedInToVault bool = true, true, false
	var readingMasterPassword, readingConfirmPassword bool = true, true
	var masterPass, confirmPass []byte
	var err error

	fd := int(os.Stdin.Fd())
	fmt.Println("Welcome to password manager.")
	for prompt {

		userSelectedOption := consoleprompts.PromptUserInput(signedInToVault)

		if signedInToVault {
			if userSelectedOption == 1 {

				if !signedInToVault {
					signedinVaultName, masterPassword := vaultOperations.SigninToVault()
					if signedinVaultName != "" && masterPassword != "" {
						signedInToVault = true
						masterPass = []byte(masterPassword)
						vaultName = signedinVaultName
					}
				}
				vaultOperations.FetchRecordsFromVault("all", vaultName, []byte(masterPass))

				continue
			} else if userSelectedOption == 2 {

				if !signedInToVault {
					signedinVaultName, masterPassword := vaultOperations.SigninToVault()
					if signedinVaultName != "" && masterPassword != "" {
						signedInToVault = true
						masterPass = []byte(masterPassword)
						vaultName = signedinVaultName
					}
				}
				vaultOperations.FetchRecordsFromVault("single", vaultName, []byte(masterPass))

				continue
			} else if userSelectedOption == 3 {

				fmt.Println("Deleting vault named " + vaultName + "...")

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
					// call another function which handles user input to determine
					// what to do next.
					vaultDeleted, err := vaultOperations.DeleteVault(vaultName + ".csv")
					if err != nil {
						fmt.Println(err.Error())
					}
					if vaultDeleted {
						signedInToVault = false
						vaultName = ""
						masterPass = []byte("")
					}
					continue
				}
				continue
			} else if userSelectedOption == 4 {
				// Sign out
				if signedInToVault {
					signedInToVault = false
					masterPass = []byte("")
					vaultName = ""
					fmt.Println("----------------------------------------------")
					fmt.Println("Signed out successfully.")
				} else {
					fmt.Println("----------------------------------------------")
					fmt.Println("You are already signed out.")
				}
				continue
			} else {
				prompt = false
				return
			}
		} else {
			if userSelectedOption == 1 {
				readingVaultName = true

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
					// call another function which handles user input to determine
					// what to do next.
					vaultOperations.ManageExistingVault(vaultName)
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
				_, err = vaultOperations.CreateVault(vaultName, string(masterPass))

				// If an error was returned, print it to the console and
				// exit the program.
				if err != nil {
					log.Fatalf("Vault creation error: %x", err)
				}
				vaultName = ""
				masterPass = []byte("")
				signedInToVault = false

				continue
			} else if userSelectedOption == 2 {
				fmt.Println("Signing in to vault...")

				if !signedInToVault {
					signedinVaultName, masterPassword := vaultOperations.SigninToVault()
					if signedinVaultName != "" && masterPassword != "" {
						signedInToVault = true
						masterPass = []byte(masterPassword)
						vaultName = signedinVaultName
					} else {
						signedInToVault = false
						masterPass = []byte("")
						vaultName = ""
						continue
					}
				} else {
					fmt.Println("----------------------------------------------")
					fmt.Println("You are already signed in.")
				}
				continue
			} else if userSelectedOption == 3 {

				if !signedInToVault {
					signedinVaultName, masterPassword := vaultOperations.SigninToVault()
					if signedinVaultName != "" && masterPassword != "" {
						signedInToVault = true
						masterPass = []byte(masterPassword)
						vaultName = signedinVaultName
					} else {
						continue
					}
				}
				vaultOperations.FetchRecordsFromVault("all", vaultName, []byte(masterPass))

				continue
			} else if userSelectedOption == 4 {

				if !signedInToVault {
					signedinVaultName, masterPassword := vaultOperations.SigninToVault()
					if signedinVaultName != "" && masterPassword != "" {
						signedInToVault = true
						masterPass = []byte(masterPassword)
						vaultName = signedinVaultName
					}
				}
				vaultOperations.FetchRecordsFromVault("single", vaultName, []byte(masterPass))

				continue
			} else {
				prompt = false
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
