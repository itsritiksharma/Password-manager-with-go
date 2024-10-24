package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"password-manager/security/hash"
)

/**
 * Decrypts the password with AES-256 algorithm and GCM.
 */
func DecryptPassword(encodedEncryptedPassword []byte) string {

	var decryptedPassword []byte

	decodedCipher, _ := hex.DecodeString(string(encodedEncryptedPassword))

	hashedKey := hash.GetPasswordHashingKey()

	key, _ := hex.DecodeString(hashedKey)
	// inClearData := []byte("Some Clear Data")

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed DecodeString: %x", err)
	}

	gcmBlock, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalf("Failed create the GCM block: %x", err)
	}

	nonceSize := gcmBlock.NonceSize()
	if len(decodedCipher) < nonceSize {
		log.Fatal("Nonce and decoded ", err)
	}

	nonce, cipherText := decodedCipher[:nonceSize], decodedCipher[nonceSize:]

	decryptedPassword, err = gcmBlock.Open(nil, nonce, cipherText, nil)

	if err != nil {
		panic(err)
	}

	password := string(decryptedPassword)

	return password
}

/**
 * Decrypts the password with AES-256 algorithm and GCM.
 */
func DecryptFileData(encodedEncryptedFile []byte) string {

	var decryptedFileData []byte

	decodedCipher, _ := hex.DecodeString(string(encodedEncryptedFile))

	hashedKey := hash.GetFileHashingKey()

	key, _ := hex.DecodeString(hashedKey)
	// inClearData := []byte("Some Clear Data")

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed DecodeString: %x", err)
	}

	gcmBlock, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalf("Failed create the GCM block: %x", err)
	}

	nonceSize := gcmBlock.NonceSize()
	if len(decodedCipher) < nonceSize {
		log.Fatal("Nonce and decoded ", err)
	}

	nonce, cipherText := decodedCipher[:nonceSize], decodedCipher[nonceSize:]

	decryptedFileData, err = gcmBlock.Open(nil, nonce, cipherText, nil)

	if err != nil {
		panic(err)
	}

	fileData := string(decryptedFileData)

	return fileData
}
