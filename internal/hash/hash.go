package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/**
 * Loads the env file and returns the value of env variable
 */
func getEnvVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

/**
 * Hashed password with provided salt using sha256 algorithm.
 */
func hash(password string, salt string) string {
	// declares new sha256 hash variable.
	h := sha256.New()

	// writes the byts from the combination of password and salt to sha256 hash variable.
	h.Write([]byte(string(password) + salt))

	// append the current hash to nil and assign it to pass variable.
	pass := h.Sum(nil)

	// Encodes the hashed password to string
	passHex := hex.EncodeToString(pass)

	return passHex
}

/**
 * Hashes the password twice first with salt then with pepper.
 */
func HashPassword(password []byte) string {

	var hashedPassword string

	pepper := getEnvVariables("PEPPER")
	passSalt := getEnvVariables("SALT")

	hashedPassword = hash(string(password), passSalt)
	hashedPassword = hash(string(hashedPassword), pepper)

	return hashedPassword

}
