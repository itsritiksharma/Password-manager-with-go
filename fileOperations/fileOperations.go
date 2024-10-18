package fileOperations

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

func FileExists(filename string) bool {
	log.SetPrefix("file exists: ")
	log.SetFlags(0)

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(filename string) (bool, error) {
	log.SetPrefix("createfile: ")
	log.SetFlags(0)

	var updateFile string

	filePresent := FileExists(filename)

	if filePresent {
		fmt.Print("File already exists. Do you want to add new credentials instead? ")
		fmt.Scan(&updateFile)

		if updateFile == "yes" {
			_, err := UpdateFile(filename)

			if err != nil {
				panic(err)
			}

			fmt.Println("File updated successfully!")

			return true, nil

		}

		return false, nil

	} else {

		data := [][]string{
			{"vegetables", "asdf"},
			{"carrot", "banana"},
			{"potato", "strawberry"},
		}

		// create a file
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			return false, errors.New("failed to create the file")
		}
		defer file.Close()

		// initialize csv writer
		writer := csv.NewWriter(file)

		defer writer.Flush()

		// write all rows at once
		writer.WriteAll(data)

		return true, nil
	}
}

// func DeleteFile(filename string) {

// }

func UpdateFile(filename string) (bool, error) {
	log.SetPrefix("update file: ")
	log.SetFlags(0)

	filePresent := FileExists(filename)

	if filePresent {

		// Open the CSV file
		file, err := os.Open(filename)
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
		file2, err := os.Create(filename)
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
