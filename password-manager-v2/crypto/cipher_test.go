package crypto

import (
	"crypto/rand"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// Test round-trip works
	text := "Hello Gophers"

	// 1. Create a key
	key_1 := make([]byte, 32)
	_, err := rand.Read(key_1)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 2. Encrypt a message
	msg, err := Encrypt(key_1, text)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}

	// 3. Decrypt it
	msg2, err := Decrypt(key_1, msg)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 4. Check it matches original

	if text != msg2 {
		t.Error("The two messages should match")
	}
}

func TestEncryptDifferentIVs(t *testing.T) {
	// Test same input produces different ciphertexts (due to random IV)
	text := "Hello Gophers"
	key_1 := make([]byte, 32)
	_, err := rand.Read(key_1)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 1. Encrypt same message twice with same key
	msg, err := Encrypt(key_1, text)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}

	msg2, err := Encrypt(key_1, text)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 2. Check ciphertexts are different

	if msg == msg2 {
		t.Error("IVs should be random something went wrong")
	}
}

func TestDecryptWrongKey(t *testing.T) {
	// Test decryption fails gracefully with wrong key
	// 1. Encrypt with key1
	text := "Hello Gophers"
	key_1 := make([]byte, 32)
	key_2 := make([]byte, 32)

	_, err := rand.Read(key_1)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	_, err = rand.Read(key_2)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}

	msg, err := Encrypt(key_1, text)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 2. Try decrypt with key2
	_, err = Decrypt(key_2, msg)
	// 3. Should get error or garbage (not crash)

	if err == nil {
		t.Error("Something should've gone wrong")
	}
}

func TestEncryptEmptyString(t *testing.T) {
	// Test encrypting empty string
	text := ""

	// 1. Create a key
	key_1 := make([]byte, 32)
	_, err := rand.Read(key_1)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}
	// 2. Encrypt a message
	msg, err := Encrypt(key_1, text)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}

	// 3. Decrypt it
	msg2, err := Decrypt(key_1, msg)
	if err != nil {
		t.Errorf("We got this err: %v", err)
	}

	// 4. Check it matches original

	if text != msg2 {
		t.Error("The two messages should match")
	}
}

func TestInvalidKeyLength(t *testing.T) {
	// Test that non-32-byte keys return error
	key := []byte{22, 33, 44, 55, 66, 11, 22}

	_, err := Encrypt(key, "Hello Gophers")
	if err == nil {
		t.Error("Expected error for short key in Encrypt")
	}

	_, err = Decrypt(key, "Hello Gophers")
	if err == nil {
		t.Error("Expected error for short key in Decrypt")
	}

}
