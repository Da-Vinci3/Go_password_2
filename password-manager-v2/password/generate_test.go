package password_test

import (
	"password-manager-v2/password"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Test normal case
	pass, err := password.Generate(16)
	if err != nil || len(pass) != 16 {
		t.Errorf("Expected 16 chars, got %d", len(pass))
	}

	// Test min length
	_, err = password.Generate(7)
	if err == nil {
		t.Error("Should reject length < 8")
	}

	// Test max length
	_, err = password.Generate(129)
	if err == nil {
		t.Error("Should reject length > 128")
	}
}
