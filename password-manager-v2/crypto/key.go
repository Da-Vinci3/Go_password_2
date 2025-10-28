package crypto

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

// This generates our super secure salt, it's cryptographically secure
// Why well the use of crypto rand is really important here, its a package built by cryptography experts to ensure secure randomness

func GenerateSalt() (string, error) {
	// In salt generation please use a random salt like genuinely use a random salt, this is where you should use
	// rand.Read it'll do it better than you can and ensure true random behaviour also its easier, don't reinvent the wheel.
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

// Once we have our salt its pretty good but we need to derive our key, that's really important
// this allows us to take our password eventually with Argon2 and derive a decryption key from it cool huh?
func DeriveKey(password, salthex string) ([]byte, error) {
	salt, err := hex.DecodeString(salthex)
	if err != nil {
		return nil, err
	}

	// Derive a 32-byte (256-bit) key using Argon2id
	key := argon2.IDKey(
		[]byte(password),
		salt,
		1,       // Time cost (iterations)
		64*1024, // Memory cost in KB (64 MB)
		4,       // Parallelism (number of threads)
		32,      // Output key length (256 bits)
	)

	return key, nil
}
