package fileOperations

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

const vaultFile string = "vaults/vault.csv"
const jsonFile string = "vaults/vaultJson.json"

// Credential represents a username/password pair.
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Vault represents a collection of credentials associated with a master password.
type Vault struct {
	VaultName      string       `json:"vaultName"`
	MasterPassword string       `json:"masterPassword"`
	Creds          []Credential `json:"creds"`
}

type Json struct {
	VaultName      string       `json:"vaultName"`
	MasterPassword string       `json:"masterPassword"`
	Creds          []Credential `json:"creds"`
}

func FileExists(filename string) bool {
	log.SetPrefix("file exists: ")
	log.SetFlags(0)

	info, err := os.Stat(vaultFile)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(vaultName, password string) (bool, error) {
	log.SetPrefix("createfile: ")
	log.SetFlags(0)

	var updateFile string

	filePresent := FileExists(vaultFile)

	if filePresent {
		fmt.Print("Vault already exists. Would you like to signin to " + vaultName + " vault instead? ")
		fmt.Scan(&updateFile)

		if updateFile == "yes" {
			//TODO: Provide the user functionality to signin to vault.
			// _, err := UpdateFile(filename)

			// if err != nil {
			// 	panic(err)
			// }

			fmt.Println("File updated successfully!")

			return true, nil

		}

		return false, nil

	} else {

		data := make(map[string]string)
		var username string
		var creds []Credential
		fd := int(os.Stdin.Fd())

		var takeInput string = "y"

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

				// hash.Hash(password)

				data[username] = string(password)

				fmt.Print("Do you want to add more credentials [y/N]? ")
				fmt.Scan(&takeInput)
			} else {
				fmt.Println("Please enter a valid option.")
				fmt.Print("Do you want to add more credentials?[y/N] ")
				fmt.Scan(&takeInput)
			}

		}

		for key, value := range data {
			fmt.Println()
			creds = append(creds, Credential{Username: key, Password: value})
		}

		finalData := []Vault{
			{
				VaultName:      vaultName,
				MasterPassword: password,
				Creds:          creds,
			},
		}

		var result [][]string
		for _, vault := range finalData {
			// Add vault name and master password as the first entry
			entry := []string{vault.VaultName, vault.MasterPassword}

			// Add each credential as a separate entry
			for _, cred := range vault.Creds {
				entry = append(entry, cred.Username, cred.Password)
			}

			// Append the entry to the result
			result = append(result, entry)
		}

		// Create the vaults directory with.
		err := os.Mkdir("vaults/", os.FileMode(0777))
		if err != nil {
			log.Fatal(err)
			return false, errors.New("failed to create the vault directory.")
		}

		// Create the vault file
		file, err := os.Create(vaultFile)
		if err != nil {
			log.Fatal(err)
			return false, errors.New("failed to create the vault file")
		}
		defer file.Close()

		// initialize csv writer
		writer := csv.NewWriter(file)

		defer writer.Flush()

		// write all rows at once
		writer.WriteAll(result)

		// Marshal the vaults slice to JSON
		jsonData, err := json.MarshalIndent(finalData, "", "  ")

		if err != nil {
			log.Fatalf("Failed to create the JSON file.")
		}

		// write all rows at once
		os.WriteFile(jsonFile, jsonData, os.FileMode(0666))

		return true, nil
	}
}

// func DeleteFile(filename string) {

// }

func UpdateFile(filename string) (bool, error) {
	log.SetPrefix("update file: ")
	log.SetFlags(0)

	filePresent := FileExists(vaultFile)

	if filePresent {

		// Open the CSV file
		file, err := os.Open(vaultFile)
		if err != nil {
			return false, errors.New("failed to open the file with filename: " + filename)
		}
		defer file.Close()

		// Read the CSV data
		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1 // Allow variable number of fields
		data, err := reader.ReadAll()
		if err != nil {
			return false, errors.New("reading from the file failed")
		}

		// Print the CSV data
		for _, row := range data {
			for _, col := range row {
				fmt.Printf("%s,", col)
			}
			fmt.Println()
		}

		// Write the CSV data
		file2, err := os.Create(vaultFile)
		if err != nil {
			return false, errors.New("writing to the file failed")
		}
		defer file2.Close()

		writer := csv.NewWriter(file2)
		defer writer.Flush()
		// this defines the header value and data values for the new csv file
		headers := []string{"name", "age", "gender"}
		data1 := [][]string{
			{"Alice", "25", "Female"},
			{"Bob", "30", "Male"},
			{"Charlie", "35", "Male"},
		}

		writer.Write(headers)
		for _, row := range data1 {
			writer.Write(row)
		}

		return true, nil
	}

	return false, errors.New("No file present with filename: " + filename)

}
