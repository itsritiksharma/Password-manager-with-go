package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"password-manager/security/hash"
)

/**
 * Encrypts the password with AES-256 algo and a hash key generated using sha-256.
 */
func EncryptPassword(password []byte) string {

	hashedKey := hash.GetHashedKey()

	key, _ := hex.DecodeString(hashedKey)
	// inClearData := []byte("Some Clear Data")

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

	cipheredText := gcmBlock.Seal(nonce, nonce, password, nil)

	encodedEncryptedPassword := hex.EncodeToString(cipheredText)

	return encodedEncryptedPassword
}
