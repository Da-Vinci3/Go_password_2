package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	myc "password-manager-v2/crypto"
	myst "password-manager-v2/storage"
	myui "password-manager-v2/ui"
)

func main() {

	var choice int
	key, err := getMasterKey()
	if err != nil {
		fmt.Printf("we encountered an error: %v", err)
	}

	for {

		// Display menu
		fmt.Println("\n=== Password Manager ===")
		fmt.Println("Here you can choose from the options available(1 - 6): ")
		fmt.Println("1. Add Password ")
		fmt.Println("2. View Passwords ")
		fmt.Println("3. Update Password ")
		fmt.Println("4. Delete Password")
		fmt.Println("5. Search Password")
		fmt.Println("6. Exit")
		// Get choice
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("we encountered an error: %v", err)
		}

		// Switch on choice (1-6)

		switch choice {
		case 1:
			myui.AddPassword(key)
		case 2:
			myui.ViewPasswords(key)
		case 3:
			myui.UpdatePassword(key)
		case 4:
			myui.DeletePassword(key)
		case 5:
			myui.SearchPasswords(key)
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again")
		}

		waitForEnter()
		clearScreen()
	}
}

// handleSetup prompts user to create master password on first run
// These are helper functions they make life a lot easier when dealing with large codebases
func handleSetup() ([]byte, error) {
	// 1. Prompt user to create master password
	var password string

	fmt.Println("It's your first time here, okay, I just need a password you will remember: ")
	fmt.Scan(&password)
	// 2. Generate salt using crypto.GenerateSalt()
	salt, err := myc.GenerateSalt()

	if err != nil {
		return nil, err
	}
	// 3. Derive key using crypto.DeriveKey()
	key, err := myc.DeriveKey(password, salt)
	if err != nil {
		return nil, err
	}
	// 4. Save config with salt and SetupComplete=true
	cfg := myst.Config{
		Salt:          salt,
		SetupComplete: true,
	}

	err = myst.SaveConfig("myconfig.json", &cfg)
	if err != nil {
		return nil, err
	}

	// 5. Return the key
	return key, nil
}

// handleLogin prompts user to enter existing master password
func handleLogin() ([]byte, error) {
	// 1. Load config using storage.LoadConfig()
	cfg, err := myst.LoadConfig("myconfig.json")

	if err != nil {
		return nil, err
	}
	// 2. Prompt user for master password
	var password string
	fmt.Println("Welcome back, do you remember your password: ")
	fmt.Scan(&password)

	// 3. Derive key using crypto.DeriveKey() with stored salt
	key, err := myc.DeriveKey(password, cfg.Salt)

	if err != nil {
		return nil, err
	}
	// 4. Return the key

	return key, nil
}

// getMasterKey determines if first run or returning user
func getMasterKey() ([]byte, error) {
	// 1. Try to load config
	cfg, err := myst.LoadConfig("myconfig.json")
	if err != nil {
		return nil, err
	}
	// 2. If config is nil (doesn't exist), call handleSetup()
	if cfg == nil {
		key, err := handleSetup()
		if err != nil {
			return nil, err
		}

		return key, nil
	} else {
		// 3. If config exists, call handleLogin()
		key, err := handleLogin()
		if err != nil {
			return nil, err
		}

		return key, nil
	}

}

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func waitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPress 0 to continue...")
	reader.ReadString('0')
}
