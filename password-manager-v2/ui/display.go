package ui

import (
	"fmt"
	myc "password-manager-v2/crypto"
	myp "password-manager-v2/password"
	myst "password-manager-v2/storage"
)

// CRUD
// This section deals with the creation of passwords and their management

func AddPassword(key []byte) {
	// we are adding a password to our file, luckily write file will create it if it doesn't exist
	var site string
	var passLength int

	stored_text, err := myst.LoadPasswords("passwords.json")
	if err != nil {
		fmt.Printf("error loading: %s\n", err)
		return
	}

	// 1. Get the site and generate password
	fmt.Println("Enter site name: ")
	_, err = fmt.Scan(&site)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	fmt.Println("Enter desired password length: ")
	_, err = fmt.Scan(&passLength)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	password, err := myp.Generate(passLength)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// 2. Assuming we have succesfully gotten the site and password now we encrypt

	site_key, err := myc.Encrypt(key, site)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}
	password_key, err := myc.Encrypt(key, password)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// 3. next we need to add them to our passwords.json file

	// we need to store it in our data structure and add that to stored_text
	// why a map and not a struct well out Config struct is already exposed and in GO
	// we try not to expose too much except where inevitable like with config.
	n := map[string]string{site_key: password_key}
	stored_text = append(stored_text, n)

	err = myst.SavePasswords("passwords.json", stored_text)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// 4. Lets tell our use what happened if we made it this far we did it
	fmt.Printf("\n✅ Password saved for %s: %s\n", site, password)

}

func ViewPasswords(key []byte) {
	// 1. Load passwords
	passwords, err := myst.LoadPasswords("passwords.json")
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}
	// 2. For each entry, decrypt site and password
	// First loop: iterate over the slice
	for _, entry := range passwords {
		// entry is map[string]string

		// Second loop: iterate over the map (usually just 1 key-value pair)
		for encryptedSite, encryptedPass := range entry {
			// Decrypt both
			site, err := myc.Decrypt(key, encryptedSite)
			if err != nil {
				fmt.Printf("error decrypting site: %s\n", err)
				continue
			}

			pass, err := myc.Decrypt(key, encryptedPass)
			if err != nil {
				fmt.Printf("error decrypting password: %s\n", err)
				continue
			}

			// Display
			fmt.Printf("Site: %s\nPassword: %s\n\n", site, pass)
		}
	}

}

func UpdatePassword(key []byte) {
	var scansite string
	// 1. Show passwords that exist

	ViewPasswords(key)

	// 2. We need to know what site we are looking for
	fmt.Println("Enter site name: ")
	_, err := fmt.Scan(&scansite)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}
	// 3. We need to Load and Decrypt so we can actually find the site

	passwords, err := myst.LoadPasswords("passwords.json")
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// Decrypt and update

	// Build new slice for what we want to keep
	var keptPasswords []map[string]string

	for _, entry := range passwords {
		// entry is map[string]string

		// Second loop: iterate over the map (usually just 1 key-value pair)
		for encryptedSite := range entry {
			// Decrypt both
			site, err := myc.Decrypt(key, encryptedSite)
			if err != nil {
				fmt.Printf("error decrypting site: %s\n", err)
				continue
			}

			// Update password
			if scansite != site {
				keptPasswords = append(keptPasswords, entry)
			} else {
				newpass, err := myp.Generate(12)
				if err != nil {
					fmt.Printf("error decrypting password: %s\n", err)
				}
				newencryptedpass, err := myc.Encrypt(key, newpass)
				if err != nil {
					fmt.Printf("error decrypting password: %s\n", err)
				}
				n := map[string]string{
					encryptedSite: newencryptedpass,
				}
				keptPasswords = append(keptPasswords, n)
			}
		}
	}

	// In save passwords we always use WriteFile which turncates all data
	err = myst.SavePasswords("passwords.json", keptPasswords)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	fmt.Printf("\n✅ Password for %s updated\n", scansite)

}

func DeletePassword(key []byte) {
	var scansite string
	// 1. Show passwords that exist
	ViewPasswords(key)
	// 2. We need to know what site we are looking for
	fmt.Println("Enter site name: ")
	_, err := fmt.Scan(&scansite)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}
	// 3. Build new slice for what we want to keep
	var keptPasswords []map[string]string

	// 4. Load the passwords we have Decrypt them, find what we want to delete and append the result
	// to our new slice then overwrite the file

	passwords, err := myst.LoadPasswords("passwords.json")
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// Decrypyt

	for _, entry := range passwords {
		for encryptedSite := range entry {
			site, err := myc.Decrypt(key, encryptedSite)
			if err != nil {
				continue
			}

			if scansite != site {
				keptPasswords = append(keptPasswords, entry)
			}

		}
	}

	// In save passwords we always use WriteFile which turncates all data
	err = myst.SavePasswords("passwords.json", keptPasswords)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	fmt.Printf("\n✅ Password for %s deleted\n", scansite)
}

func SearchPasswords(key []byte) {
	var scansite string

	// 1. Load passwords
	passwords, err := myst.LoadPasswords("passwords.json")
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// 2. Prompt the user for the site they want
	fmt.Println("Enter site name: ")
	_, err = fmt.Scan(&scansite)
	if err != nil {
		fmt.Printf("we encountered an error: %s \t", err)
	}

	// 3. For each entry, decrypt site and password
	// First loop: iterate over the slice
	found := false

	for _, entry := range passwords {
		for encryptedSite, encryptedPass := range entry {
			site, err := myc.Decrypt(key, encryptedSite)
			if err != nil {
				continue
			}

			pass, err := myc.Decrypt(key, encryptedPass)
			if err != nil {
				continue
			}

			if scansite == site {
				fmt.Printf("Site: %s\nPassword: %s\n\n", site, pass)
				found = true
				break // Found it, stop searching
			}

		}
	}

	if !found {
		fmt.Println("We couldn't find that password, you can always add it")
	}

}
