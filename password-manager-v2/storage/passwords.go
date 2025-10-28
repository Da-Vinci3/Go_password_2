package storage

import (
	"encoding/json"
	"os"
)

// LoadPasswords reads encrypted password data from a JSON file.
// Returns empty slice if file doesn't exist (not an error - first use case).
func LoadPasswords(filename string) ([]map[string]string, error) {
	// 1. Try to read file
	file, err := os.ReadFile(filename)

	// 2. If file doesn't exist, return empty slice and nil error

	if os.IsNotExist(err) {
		return []map[string]string{}, nil
	}
	// 3. If other error, return it

	if err != nil {
		return nil, err
	}

	// 4. Unmarshal JSON into []map[string]string
	// with the map just let Unmarshal handle it no need to use make her
	var data []map[string]string

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	// 5. Return the data

	return data, nil
}

// SavePasswords writes encrypted password data to a JSON file.
func SavePasswords(filename string, passwords []map[string]string) error {
	// 1. Marshal to JSON with indentation (pretty print)
	data, err := json.MarshalIndent(passwords, "", "\t")

	if err != nil {
		return err
	}
	// 2. Write to file with 0600 permissions

	err = os.WriteFile(filename, data, 0600)
	if err != nil {
		return err
	}
	// 3. Return any error

	return nil
}
