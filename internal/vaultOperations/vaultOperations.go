package vaultOperations

import (
	"errors"
	"password-manager/internal/fileOperations"
	"password-manager/security/encryption"
)

/**
 * Function to allow the user to crate a password vault.
 */
func CreateVault(vaultName, masterPass string) (bool, error) {

	// If no name was given, return an error with a message.
	if vaultName == "" {
		return false, errors.New("empty vault name")
	}

	// if len(masterPass) < 8 {
	// 	return false, errors.New("password should be 8 characters long")
	// }

	hashedMasterPassword := encryption.EncryptPassword([]byte(masterPass))

	// Create the vault file with the password.
	file, err := fileOperations.CreateFile(vaultName, hashedMasterPassword)

	if err != nil {
		return false, errors.New("file creation error")
	}

	if !file {
		return false, nil
	}

	return true, nil
}

/**
 * Function to allow the user to signin to the vault.
 */
// func SigninToVault(vaultName, enteredMasterPass string) (bool, error) {
// 	if vaultName == "" {
// 		return false, errors.New("empty vault name")
// 	}

// 	if enteredMasterPass != masterPass {
// 		return false, errors.New("passwords don't match")
// 	}
// }
