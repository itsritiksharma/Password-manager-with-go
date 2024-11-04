package fileOperations

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"password-manager/security/encryption"
	"strings"

	"golang.org/x/term"
)

const vaultDir string = "vaults/"

// Credential represents a username/password pair.
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Vault represents a collection of credentials associated with a master password.
type Vault struct {
	Creds []Credential `json:"creds"`
}

type Json struct {
	VaultName      string       `json:"vaultName"`
	MasterPassword string       `json:"masterPassword"`
	Creds          []Credential `json:"creds"`
}

func createCsvFile(vaultFile string, csvData [][]string) (bool, error) {
	// Create the vault file
	file, err := os.Create(vaultDir + vaultFile)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("failed to create the vault file")
	}
	defer file.Close()

	// initialize csv writer
	csvWriter := csv.NewWriter(file)

	defer csvWriter.Flush()

	// write all rows at once
	csvWriter.WriteAll(csvData)

	return true, nil
}

func createJsonFile(jsonFileData []Json) (bool, error) {

	// Marshal the vaults slice to JSON
	jsonData, err := json.MarshalIndent(jsonFileData, "", "  ")

	if err != nil {
		log.Fatalf("Failed to create the JSON file.")
	}

	jsonFileExists, err := FileExists("VaultsInfo.json")
	if err != nil {
		log.Fatal("Json file exists: ", err)
	}

	if jsonFileExists {
		readJsonFile, err := os.ReadFile(vaultDir + "VaultsInfo.json")
		if err != nil {
			log.Fatal("reading Json file error: ", err)
		}

		jsonFileData := string(readJsonFile)

		splitJsonFile := strings.Split(jsonFileData, "][")

		stringg := strings.Join(splitJsonFile, ",")

		trimmedJsonFileData := stringg[:len(stringg)-1]
		trimmedJoiningString := string(jsonData[1 : len(jsonData)-1])

		formattedJsonData := strings.Join([]string{trimmedJsonFileData, trimmedJoiningString}, ",") + "]"

		// If the file doesn't exist, create it, or append to the file
		openJsonFile, err := os.Create(vaultDir + "VaultsInfo.json")
		if err != nil {
			log.Fatal(err)
		}
		defer openJsonFile.Close()

		_, err = openJsonFile.Write([]byte(formattedJsonData))
		if err != nil {
			openJsonFile.Close() // ignore error; Write error takes precedence
			log.Fatal("Failed to write to file: ", err)
		}
		err = openJsonFile.Close()
		if err != nil {
			log.Fatal("Failed to close the json file: ", err)
		}

		return true, nil
	}

	// If the file doesn't exist, create it, or append to the file
	openJsonFile, err := os.Create(vaultDir + "VaultsInfo.json")
	if err != nil {
		log.Fatal(err)
	}
	defer openJsonFile.Close()

	_, err = openJsonFile.Write(jsonData)
	if err != nil {
		openJsonFile.Close() // ignore error; Write error takes precedence
		log.Fatal("Failed to write to file: ", err)
	}
	err = openJsonFile.Close()
	if err != nil {
		log.Fatal("Failed to close the json file: ", err)
	}

	return true, nil
}

func DirExists() bool {
	log.SetPrefix("file exists: ")
	log.SetFlags(0)

	files, err := os.ReadDir(vaultDir)
	if err != nil {
		return false
	}
	if len(files) <= 0 {
		return true
	}
	return true
}

func FileExists(filename string) (bool, error) {
	log.SetPrefix("file exists: ")
	log.SetFlags(0)

	files, err := os.ReadDir(vaultDir)
	if err != nil {
		return false, errors.New("vaults directory does not exist")
	}

	for _, file := range files {
		if file.Name() == filename {
			return true, nil
		}
	}

	return false, nil
}

func CreateFile(vaultName, masterPassword string) (bool, error) {
	log.SetPrefix("createfile: ")
	log.SetFlags(0)

	var vault []string
	var username string
	var takeInput string = "y"

	var creds []Credential
	var csvData [][]string
	var vaultFile = vaultName + ".csv"

	fd := int(os.Stdin.Fd())
	credData := make(map[string]string)

	// Prompt user to enter username and password.
	for takeInput != "N" {
		if takeInput == "y" {
			fmt.Print("Enter username: ")
			fmt.Scan(&username)

			fmt.Print("Enter password for " + username + ": ")
			password, err := term.ReadPassword(fd)
			fmt.Println()

			if err != nil {
				panic(err)
			}

			credData[username] = string(password)

			fmt.Print("Do you want to add more credentials [y/N]? ")
			fmt.Scan(&takeInput)
		} else {
			fmt.Println("Please enter a valid option.")
			fmt.Print("Do you want to add more credentials?[y/N] ")
			fmt.Scan(&takeInput)
		}
	}

	// Create a mapping of credential from the data entered by user.
	for key, value := range credData {
		encryptedCredPassword := encryption.EncryptPassword([]byte(value), masterPassword)

		creds = append(creds, Credential{Username: key, Password: encryptedCredPassword})
	}

	// Hash the username for creds.
	encryptedVaultName := encryption.EncryptPassword([]byte(vaultName), masterPassword)

	// Hash the password for creds.
	encryptedMasterPassword := encryption.EncryptPassword([]byte(masterPassword), masterPassword)

	vault = append(vault, encryptedVaultName, encryptedMasterPassword)

	// Add each credential as a separate entry
	for _, cred := range creds {

		// Hash the username for creds.
		encryptedUsername := encryption.EncryptPassword([]byte(cred.Username), masterPassword)

		vault = append(vault, encryptedUsername, cred.Password)

	}

	// Append the entry to the result
	csvData = append(csvData, vault)

	_, err := createCsvFile(vaultFile, csvData)
	if err != nil {
		return false, errors.New("failed to create the csv file")
	}

	hashedMasterPassword := encryption.EncryptPassword([]byte(masterPassword), string(masterPassword))

	jsonFileData := []Json{{
		VaultName:      vaultName,
		MasterPassword: hashedMasterPassword,
		Creds:          creds,
	}}

	_, err = createJsonFile(jsonFileData)
	if err != nil {
		return false, errors.New("failed to create the json file")
	}

	// Encrypt csv
	csvFileExists, err := FileExists(vaultFile)
	if err != nil {
		log.Fatal("Json file exists error: ", err)
	}

	if !csvFileExists {
		return false, errors.New("csv file does not exist for encryption")
	}
	_, err = encryption.EncryptFile(vaultDir+vaultFile, masterPassword)
	if err != nil {
		panic(err.Error())
	}

	// Encrypt JSON
	// jsonFileExists, err := FileExists("VaultsInfo.json")
	// if err != nil {
	// 	log.Fatal("Json file exists error: ", err)
	// }

	// if !jsonFileExists {
	// 	return false, errors.New("csv file does not exist for encryption")
	// }
	// _, err = encryption.EncryptFile(vaultDir+"VaultsInfo.json", masterPassword)
	// if err != nil {
	// 	panic(err.Error())
	// }

	return true, nil
}

// func DeleteFile(filename string) {

// }

// func UpdateFile(filename string) (bool, error) {
// 	log.SetPrefix("update file: ")
// 	log.SetFlags(0)

// 	filePresent := FileExists(vaultFile)

// 	if filePresent {
// If the file doesn't exist, create it, or append to the file

// TODO Code to append to file
// f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// if err != nil {
// 	log.Fatal(err)
// }
// if _, err := f.Write([]byte("appended some data\n")); err != nil {
// 	f.Close() // ignore error; Write error takes precedence
// 	log.Fatal(err)
// }
// if err := f.Close(); err != nil {
// 	log.Fatal(err)
// }

// 		// Open the CSV file
// file, err := os.Open(vaultFile)
// 		if err != nil {
// 			return false, errors.New("failed to open the file with filename: " + filename)
// 		}
// 		defer file.Close()

// 		// Read the CSV data
// 		reader := csv.NewReader(file)
// 		reader.FieldsPerRecord = -1 // Allow variable number of fields
// 		data, err := reader.ReadAll()
// 		if err != nil {
// 			return false, errors.New("reading from the file failed")
// 		}

// 		// Print the CSV data
// 		for _, row := range data {
// 			for _, col := range row {
// 				fmt.Printf("%s,", col)
// 			}
// 			fmt.Println()
// 		}

// 		// Write the CSV data
// 		file2, err := os.Create(vaultFile)
// 		if err != nil {
// 			return false, errors.New("writing to the file failed")
// 		}
// 		defer file2.Close()

// 		writer := csv.NewWriter(file2)
// 		defer writer.Flush()
// 		// this defines the header value and data values for the new csv file
// 		headers := []string{"name", "age", "gender"}
// 		data1 := [][]string{
// 			{"Alice", "25", "Female"},
// 			{"Bob", "30", "Male"},
// 			{"Charlie", "35", "Male"},
// 		}

// 		writer.Write(headers)
// 		for _, row := range data1 {
// 			writer.Write(row)
// 		}

// 		return true, nil
// 	}

// 	return false, errors.New("No file present with filename: " + filename)

// }
