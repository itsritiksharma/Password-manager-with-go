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
func hash(hashingString string, salt string) string {
	// declares new sha256 hash variable.
	h := sha256.New()

	// writes the byts from the combination of password and salt to sha256 hash variable.
	h.Write([]byte(string(hashingString) + salt))

	// append the current hash to nil and assign it to pass variable.
	hashedString := h.Sum(nil)

	// Encodes the hashed password to string
	stringHex := hex.EncodeToString(hashedString)

	return stringHex
}

/**
 * Hashes the password with salt.
 */
func GetHashedKey() string {

	var hashedKey string

	encKey := getEnvVariables("ENCRYPTION_KEY")
	salt := getEnvVariables("SALT")

	hashedKey = hash(string(encKey), salt)

	return hashedKey

}
