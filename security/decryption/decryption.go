package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"
	"password-manager/security/hash"
)

/**
 * Decrypts the password with AES-256 algorithm and GCM.
 */
func DecryptPassword(encodedEncryptedPassword []byte, masterPassword string) string {

	var decryptedPassword []byte

	decodedCipher, _ := hex.DecodeString(string(encodedEncryptedPassword))

	hashedKey := hash.GetPasswordHashingKey(masterPassword)

	key, _ := hex.DecodeString(hashedKey)

	dataVerifier := []byte(masterPassword)

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

	decryptedPassword, err = gcmBlock.Open(nil, nonce, cipherText, dataVerifier)

	if err != nil {
		panic(err)
	}

	password := string(decryptedPassword)

	return password
}

/**
 * Decrypts the password with AES-256 algorithm and GCM.
 */
func DecryptFile(fileName string, masterPassword string) string {

	var decryptedFileData []byte

	readData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	hashedKey := hash.GetFileHashingKey(masterPassword)

	key, _ := hex.DecodeString(hashedKey)
	dataVerifier := []byte(masterPassword)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed DecodeString: %x", err)
	}

	gcmBlock, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalf("Failed create the GCM block: %x", err)
	}

	nonceSize := gcmBlock.NonceSize()
	if len(readData) < nonceSize {
		log.Fatal("Nonce and decoded ", err)
	}

	nonce, cipherText := readData[:nonceSize], readData[nonceSize:]

	decryptedFileData, err = gcmBlock.Open(nil, nonce, cipherText, dataVerifier)

	if err != nil {
		panic(err)
	}

	fileData := string(decryptedFileData)

	return fileData
}
