package passwordmanager

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/consoleprompts"
	"password-manager/internal/fileOperations"
	"password-manager/internal/vaultOperations"
	"password-manager/security/decryption"
	"strings"

	"golang.org/x/term"
)

func promptUser() {

	var vaultName string
	var prompt string = "prompt"
	var readingVaultName bool = true
	var readingMasterPassword, readingConfirmPassword bool = true, true
	var masterPass, confirmPass []byte
	var err error

	fd := int(os.Stdin.Fd())
	fmt.Println("Welcome to password manager.")
	for prompt != "do-not-prompt" {
		if prompt == "prompt" {

			userSelectedOption := consoleprompts.PromptUserInput()

			if userSelectedOption == 1 {

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
				var readingMasterPassword bool = true
				var enteredMasterPass []byte
				var readingVaultName bool = true
				var goBack string = "no"
				var decodedFile string
				var credsFromVault []string

				fd := int(os.Stdin.Fd())

				fmt.Println("Fetching all credentials from the vault...")

				for readingVaultName {
					if goBack == "no" {
						fmt.Print("Enter vault name: ")
						fmt.Scan(&vaultName)

						if vaultName == "" {
							fmt.Println("Invalid vault name. Please try again.")
							continue
						}

						vaultFileExists, err := fileOperations.FileExists(vaultName + ".csv")
						if !vaultFileExists || err != nil {
							fmt.Println("No vault exists by this name. Please try again.")
							fmt.Print("Want to go back to main menu? [y/N]")
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
							}
						}

						for i := 2; i < len(credsFromVault)-1; i = i + 2 {
							decodedUsername := decryption.DecryptPassword([]byte(credsFromVault[i]), string(enteredMasterPass))
							decodedPassword := decryption.DecryptPassword([]byte(credsFromVault[i+1]), string(enteredMasterPass))

							fmt.Println("---------------------------------------------------------------------------")
							fmt.Printf("Username: %s", decodedUsername)
							fmt.Println()
							fmt.Printf("Password: %s", decodedPassword)
							fmt.Println()
							fmt.Println("---------------------------------------------------------------------------")
						}
						readingVaultName = false
						break
					}
				}

				continue

			} else if userSelectedOption == 5 {
				var vaultName string
				var credName string
				var credFound bool = false

				fmt.Println("Fetching the requested credential from vault...")

				fd := int(os.Stdin.Fd())

				fmt.Print("Enter vault name: ")
				fmt.Scan(&vaultName)
				fmt.Print("Please the master password: ")
				masterPass, err := term.ReadPassword(fd)
				fmt.Println()
				if err != nil {
					panic(err)
				}

				fmt.Print("Enter cred to fetch: ")
				fmt.Scan(&credName)
				// hashedMasterPassword := encryption.EncryptPassword([]byte(masterPass), string(masterPass))

				decodedFile, err := decryption.DecryptFile("vaults/"+vaultName+".csv", string(masterPass))

				creds := strings.Split(decodedFile, ",")

				for i := 0; i < len(creds)-1; i = i + 2 {

					decodedUsername := decryption.DecryptPassword([]byte(creds[i]), string(masterPass))

					if decodedUsername == credName {

						decodedPassword := decryption.DecryptPassword([]byte(creds[i+1]), string(masterPass))

						fmt.Printf("%s,%s", decodedUsername, decodedPassword)
						fmt.Println()

						credFound = true
						break
					}

				}

				if !credFound {
					fmt.Println("No cred found by this name.")
				}

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

				decodedFile, err := decryption.DecryptFile("vaults/VaultsInfo.json", string(masterPass))
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
