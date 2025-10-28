package crypto

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	// Your test here
	_, err := GenerateSalt()
	if err != nil {
		t.Errorf("expected no error got: %v", err)
	}

}

func TestDeriveKey(t *testing.T) {
	// Your test here
	password := "GoGophers"
	saltHex := "840bf3abbc4e706e81522bbcb607a567"
	_, err := DeriveKey(password, saltHex)

	if err != nil {
		t.Errorf("expected no error got: %v", err)
	}
}
