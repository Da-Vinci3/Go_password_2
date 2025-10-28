# Go Password Manager v2

A secure, command-line password manager built in Go with modern cryptographic practices.

## Features

- **Argon2id Key Derivation**: Uses a master password instead of a random key
- **AES-256-CBC Encryption**: Industry-standard encryption for all stored passwords
- **Secure Password Generation**: Cryptographically secure random passwords
- **Full CRUD Operations**: Add, view, update, delete, and search passwords
- **Persistent Storage**: Encrypted JSON storage with proper configuration management

## What's New in v2

v2 is a complete rewrite with significant improvements over v1:

- ✅ **Argon2id key derivation** - Use a memorable master password instead of a random 32-character string
- ✅ **Proper project structure** - Separated into logical packages (crypto, storage, password, ui)
- ✅ **Configuration management** - Salt stored in config file instead of counter.txt hack
- ✅ **Comprehensive test suite** - All core functionality is tested
- ✅ **Better error handling** - Consistent error checking throughout
- ✅ **Improved UX** - Clear screen, better prompts, user-friendly interface

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/password-manager-v2.git
cd password-manager-v2

# Build the application
go build -o password-manager

# Run
./password-manager
```

## Usage

### First Run (Setup)

On first run, you'll be prompted to create a master password:

```
It's your first time here, okay, I just need a password you will remember: 
```

This password will be used to derive your encryption key via Argon2id. **Remember this password** - if you lose it, your data cannot be recovered.

### Menu Options

```
=== Password Manager ===
1. Add Password 
2. View Passwords 
3. Update Password 
4. Delete Password
5. Search Password
6. Exit
```

### Adding a Password

1. Select option 1
2. Enter the site name (e.g., "Facebook")
3. Enter desired password length (8-128 characters)
4. Password is automatically generated and displayed

### Viewing Passwords

Select option 2 to see all stored passwords with their associated sites.

### Updating a Password

1. Select option 3
2. View your existing passwords
3. Enter the site name to update
4. A new password is generated and displayed

### Deleting a Password

1. Select option 4
2. View your existing passwords
3. Enter the site name to delete
4. Confirmation message displayed

### Searching for a Password

1. Select option 5
2. Enter the site name to search for
3. Password is displayed if found

## Project Structure

```
password-manager-v2/
├── main.go              # Entry point and menu logic
├── crypto/
│   ├── key.go          # Argon2 key derivation
│   ├── cipher.go       # AES-256-CBC encryption/decryption
│   ├── key_test.go     # Key derivation tests
│   └── cipher_test.go  # Encryption tests
├── storage/
│   ├── config.go       # Configuration file management
│   ├── passwords.go    # Password storage management
│   ├── config_test.go  # Config tests
│   └── passwords_test.go # Storage tests
├── password/
│   ├── generate.go     # Secure password generation
│   └── generate_test.go # Generation tests
└── ui/
    └── crud.go         # CRUD operation implementations
```

## Security Features

### Key Derivation
- **Algorithm**: Argon2id (winner of the Password Hashing Competition)
- **Parameters**: 
  - Time cost: 1 iteration
  - Memory cost: 64 MB
  - Parallelism: 4 threads
  - Output: 32 bytes (256 bits)

### Encryption
- **Algorithm**: AES-256-CBC
- **Key Size**: 256 bits
- **IV**: Random 16-byte IV per encryption
- **Padding**: PKCS7

### Password Generation
- **Entropy Source**: `crypto/rand` (cryptographically secure)
- **Character Set**: 95 printable ASCII characters (a-z, A-Z, 0-9, symbols)
- **Length**: User-configurable (8-128 characters)

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test ./... -v
```

Run tests for a specific package:
```bash
go test ./crypto -v
go test ./storage -v
go test ./password -v
```

## Files

- `myconfig.json` - Stores salt and setup completion flag (NOT secret)
- `passwords.json` - Stores encrypted site names and passwords (encrypted, but backup recommended)

**Important**: Both files are created automatically. Do not manually edit them.

## Security Considerations

### What's Protected
✅ All passwords are encrypted with AES-256
✅ Master password is never stored
✅ Salt is randomly generated per installation
✅ Each encryption uses a unique random IV

### What's NOT Protected
❌ File metadata (file existence, modification times)
❌ Number of stored passwords
❌ Application usage patterns

### Recommendations
- Use a strong, unique master password
- Store `passwords.json` in an encrypted filesystem for additional protection
- Regularly backup `myconfig.json` and `passwords.json` together
- Never commit these files to version control

## Limitations

- No cloud sync
- No password strength checking
- No password expiration reminders
- No import/export functionality
- No multi-user support

## Future Improvements

- [ ] Password strength meter
- [ ] Import/export (encrypted backup)
- [ ] Password expiration tracking
- [ ] Clipboard integration
- [ ] Two-factor authentication for master password
- [ ] Cloud sync option

## License

MIT License - See LICENSE file for details

## Contributing

Contributions welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## Acknowledgments

- Built as a learning project to understand cryptography and secure password management
- Inspired by the need for a simple, transparent password manager
- Special thanks to the Go crypto library maintainers

## Disclaimer

This is an educational project. While it implements proper cryptographic practices, it has not undergone a professional security audit. For mission-critical password storage, consider using established solutions like 1Password, Bitwarden, or KeePass.

---

**Remember**: Your master password is the key to all your data. Choose wisely and never forget it!
