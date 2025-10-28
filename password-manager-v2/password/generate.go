package password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

var (
	LowercaseBytes = []byte("abcdefghijklmnopqrstuvwxyz")
	UppercaseBytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	DigitBytes     = []byte("0123456789")
	SymbolBytes    = []byte("!@#$%^&*()_+-=[]{}|;:,.<>?")
)

func Generate(length int) (string, error) {

	// 1. Validate length (min 8, max 128 - or whatever you want)
	if length < 8 {
		return "", errors.New("we cannot generate a password less than 8 characters")
	}

	if length > 128 {
		return "", errors.New("we cannot generate a password greater than 128 characters")
	}
	// 1.5 Instanciate byte for storage
	passwordGen := make([]byte, 0, length)

	// 2. Define character set (a-z, A-Z, 0-9, symbols)
	var AllCharBytes = append(append(append(LowercaseBytes, UppercaseBytes...), DigitBytes...), SymbolBytes...)

	// 3. Use crypto/rand to randomly select characters
	lengthChar := len(AllCharBytes)

	for range length {
		choice, err := rand.Int(rand.Reader, big.NewInt(int64(lengthChar)))

		if err != nil {
			return "", err
		}

		passwordGen = append(passwordGen, AllCharBytes[choice.Int64()])

	}
	// 4. Return the password string

	return string(passwordGen), nil
}
