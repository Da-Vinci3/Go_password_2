package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// addPadding adds PKCS7 padding to data
func addPadding(data []byte, blockSize int) []byte {
	// Calculate how many bytes of padding needed
	// r means remainder so whatever we get from modulo
	r := len(data) % blockSize
	// Create padding bytes (each byte = padding length)
	// pl means padding length
	pl := blockSize - r
	// Append padding to data
	for i := 0; i < pl; i++ {
		data = append(data, byte(pl))
	}
	// Return padded data
	return data
}

// removePadding removes PKCS7 padding from data
func removePadding(data []byte) ([]byte, error) {
	// Lets check we actually have something to remove
	if data == nil || len(data) == 0 {
		return nil, nil
	}
	// Get the last byte (tells you padding length)
	// pc means padding character, so we know what we have
	pc := data[len(data)-1]

	// Validate padding length is reasonable
	pl := int(pc)

	// Validate all padding bytes are correct
	err := checkPaddingIsValid(data, pl)
	if err != nil {
		return nil, err
	}
	// Return data without padding
	return data[:len(data)-pl], nil
}

func checkPaddingIsValid(data []byte, pl int) error {
	// pl is padding length, I didn't want to be verbose
	if len(data) < pl {
		return errors.New("invalid padding")
	}
	// give me the padding thats what p is so start at the padding and go from there really cool huh?
	// honestly you should read articles on padding it's boring until you know it then its fun, like super fun.
	p := data[len(data)-(pl):]
	for _, pc := range p {
		// this is type casting, make them both uints and compare, so 8=8 etc
		// not type casting like in holywood but the good kind
		if uint(pc) != uint(len(p)) {
			return errors.New("invalid padding")
		}
	}
	return nil
}

// Encrypt encrypts plaintext using AES-256-CBC
func Encrypt(key []byte, plaintext string) (string, error) {
	// 1. Validate key length is 32 bytes

	if len(key) != 32 {
		return "", errors.New("the key must be exactly 32 bytes")
	}
	// 2. Create AES cipher block

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	// 3. Convert plaintext to bytes and add padding
	plaintextBytes := []byte(plaintext)
	paddedPlaintext := addPadding(plaintextBytes, aes.BlockSize)

	// 4. Generate random IV (16 bytes)
	iv := make([]byte, 16)
	_, err = rand.Read(iv)

	if err != nil {
		return "", err
	}

	// 5. Create CBC encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// 6. Encrypt the padded data
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// 7. Prepend IV to ciphertext

	result := append(iv, ciphertext...)
	// 8. Hex encode and return
	hexString := hex.EncodeToString(result)

	return hexString, nil
}

// Decrypt decrypts ciphertext using AES-256-CBC
func Decrypt(key []byte, ciphertextHex string) (string, error) {
	// 1. Validate key length is 32 bytes
	if len(key) != 32 {
		return "", errors.New("the key must be exactly 32 bytes")
	}
	// 2. Hex decode the input
	hexString, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	// 3. Split: first 16 bytes = IV, rest = ciphertext
	usedHex := hexString[16:]
	iv := hexString[:16]

	// 4. Create AES cipher block
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	// 5. Create CBC decrypter
	mode := cipher.NewCBCDecrypter(block, iv)

	// 6. Decrypt the ciphertext
	plaintext := make([]byte, len(usedHex))
	mode.CryptBlocks(plaintext, usedHex)

	// 7. Remove padding
	result, err := removePadding(plaintext)

	if err != nil {
		return "", err
	}
	// 8. Convert to string and return
	usedResult := string(result)

	return usedResult, nil
}
