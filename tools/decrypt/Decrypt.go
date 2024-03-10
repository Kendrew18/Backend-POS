package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func decryptStruct(c echo.Context) error {
	// Get the encrypted data from the request body
	body := c.Request().Body
	defer body.Close()

	// Decode the encrypted data from hexadecimal format

	decode, _ := string(io.ReadAll(body))

	encryptedData, err := hex.DecodeString(strings.TrimSpace(decode))
	if err != nil {
		return err
	}

	// Create a new Cipher Block from the key
	key := []byte("mysecretkey")
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Get the nonce and ciphertext
	nonceSize := gcm.NonceSize()
	nonce := encryptedData[:nonceSize]
	ciphertext := encryptedData[nonceSize:]

	// Decrypt the JSON data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Convert the JSON data to a struct
	var dataInterface interface{}
	err = json.Unmarshal(plaintext, &dataInterface)
	if err != nil {
		return err
	}

	// Return the decrypted struct
	return c.JSON(http.StatusOK, dataInterface)
}
