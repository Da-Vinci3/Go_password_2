package storage

import (
	"encoding/json"
	"os"
)

// This is our config struct honestly it's quite simple all it does is store the salt, the salt doesn't really matter if people can see it, its derivation is complex
// You can reverse engineer it easily
type Config struct {
	Salt          string
	SetupComplete bool
}

// Here we need to save our salt for future uses otherwise i can't guarantee you get your data back
// and that is a big no, no especially for a password app, imagine you get locked out?
func SaveConfig(filename string, cfg *Config) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, bytes, 0600)
	if err != nil {
		return err
	}

	return nil

}

// This one loads the config if you have a config it loads it otherwise we get an error
// Its really important, like the key to you house, we need to make sure you have a house first
func LoadConfig(filename string) (*Config, error) {

	cfg := Config{}
	file, err := os.ReadFile(filename)

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
