package vaultCreation

import (
	"errors"
	"password-manager/internal/fileOperations"
)

func CreateVault(vaultName, masterPass, confirmPass string) (bool, error) {

	// If no name was given, return an error with a message.
	if vaultName == "" {
		return false, errors.New("empty vault name")
	}

	if masterPass != confirmPass {
		return false, errors.New("passwords don't match")
	}

	// Create the vault file with the password.
	file, err := fileOperations.CreateFile(vaultName + ".csv")

	if err != nil {
		return false, errors.New("file creation error")
	}

	if !file {
		return false, nil
	}

	return true, nil
}
