package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"password-manager/security/hash"
)

/**
 * Encrypts the password with AES-256 algo and a hash key generated using sha-256.
 */
func EncryptPassword(password []byte, masterPassword string) string {

	hashedKey := hash.GetPasswordHashingKey(masterPassword)

	key, _ := hex.DecodeString(hashedKey)

	dataVerifier := []byte(masterPassword)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create new cipher: %x", err)
	}

	gcmBlock, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalf("Failed to create new GCM block: %x", err)
	}

	nonce := make([]byte, gcmBlock.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {

		panic(err.Error())
	}

	cipheredText := gcmBlock.Seal(nonce, nonce, password, dataVerifier)

	encodedEncryptedPassword := hex.EncodeToString(cipheredText)

	return encodedEncryptedPassword
}

/**
 * Encrypts the File with AES-256 algo and a hash key generated using sha-256.
 */
func EncryptFile(fileName string, masterPassword string) (bool, error) {

	hashedKey := hash.GetFileHashingKey(masterPassword)

	key, _ := hex.DecodeString(hashedKey)

	dataVerifier := []byte(masterPassword)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create new cipher: %x", err)
	}

	gcmBlock, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalf("Failed to create new GCM block: %x", err)
	}

	nonce := make([]byte, gcmBlock.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	readData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	cipheredText := gcmBlock.Seal(nonce, nonce, readData, dataVerifier)

	// If the file doesn't exist, create it, or append to the file
	err = os.WriteFile(fileName, []byte(cipheredText), 0644)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("error writing to file.")
	}

	return true, nil

}
