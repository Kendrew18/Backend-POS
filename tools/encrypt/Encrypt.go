package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
)

func encryptStruct(data interface{}, key []byte) ([]byte, error) {
	// Convert the data to a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a new nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the JSON data
	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	// Return the encrypted data
	return ciphertext, nil
}
