//    \\ SPIKE: Keep your secrets secret with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package route

import (
	"fmt"
	"net/http"
)

func factory(p, a, m string) handler {
	fmt.Println("factory:", "p", p, "a", a, "m", m)

	switch {
	// Route to fetch the Keystone status.
	// The status can be "pending" or "ready".
	case m == http.MethodPost && a == "" && p == urlInit:
		return routeInit
	case m == http.MethodPost && a == "" && p == urlSecrets:
		return routePostSecret
	case m == http.MethodPost && a == "get" && p == urlSecrets:
		return routeGetSecret
	// Fallback route.
	default:
		return routeFallback
	}
}

func Route(r *http.Request, w http.ResponseWriter) {
	factory(r.URL.Path, r.URL.Query().Get("action"), r.Method)(r, w)
}

/*
 The most balanced way is to keep the root key locally in a folder that the
 user configures; restrict access to the key by setting proper file and folder
 permissions and ask for password (instead of key) while logging in.

 	 # First time setup
 	 spike encrypt-config --password-prompt
 	 Enter password: ******
 	 # Encrypts the root key using a key derived from the password
 	 # Saves the encrypted key + salt to config

 	 # Regular usage
 	 spike login --password-prompt
 	 Enter password: ******
 	 # Internally:
 	 # 1. Derives encryption key from password + stored salt
 	 # 2. Decrypts the root key
 	 # 3. Uses root key to get session key


*/

/*
	package config

	import (
		"fmt"
		"os"
		"path/filepath"
		"runtime"
		"syscall"

		"gopkg.in/yaml.v3"
	)

	type Config struct {
		RootKey          string `yaml:"root_key,omitempty"`
		EncryptedRootKey string `yaml:"encrypted_root_key,omitempty"`
		Salt             []byte `yaml:"salt,omitempty"`
		IsEncrypted      bool   `yaml:"is_encrypted"`
	}

	type ConfigManager struct {
		configPath string
		config     Config
	}

	// NewConfigManager creates a new configuration manager
	func NewConfigManager(customPath string) (*ConfigManager, error) {
		configPath, err := resolveConfigPath(customPath)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve config path: %w", err)
		}

		cm := &ConfigManager{
			configPath: configPath,
		}

		// Load existing config if it exists
		if err := cm.loadConfig(); err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load config: %w", err)
		}

		return cm, nil
	}

	// resolveConfigPath determines the configuration file path
	func resolveConfigPath(customPath string) (string, error) {
		if customPath != "" {
			absPath, err := filepath.Abs(customPath)
			if err != nil {
				return "", err
			}
			return absPath, nil
		}

		// Get user's home directory
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}

		// Default path is ~/.spike/config.yaml
		return filepath.Join(home, ".spike", "config.yaml"), nil
	}

	// ensureConfigDir ensures the configuration directory exists with proper permissions
	func (cm *ConfigManager) ensureConfigDir() error {
		configDir := filepath.Dir(cm.configPath)

		// Create directory with restricted permissions
		err := os.MkdirAll(configDir, 0700)
		if err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}

		// On Unix-like systems, explicitly set directory permissions
		if runtime.GOOS != "windows" {
			if err := os.Chmod(configDir, 0700); err != nil {
				return fmt.Errorf("failed to set directory permissions: %w", err)
			}
		}

		return nil
	}

	// loadConfig loads the configuration from file
	func (cm *ConfigManager) loadConfig() error {
		data, err := os.ReadFile(cm.configPath)
		if err != nil {
			return err
		}

		return yaml.Unmarshal(data, &cm.config)
	}

	// saveConfig saves the configuration to file with proper permissions
	func (cm *ConfigManager) saveConfig() error {
		if err := cm.ensureConfigDir(); err != nil {
			return err
		}

		// Create or truncate the config file with restricted permissions
		file, err := os.OpenFile(cm.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return fmt.Errorf("failed to open config file: %w", err)
		}
		defer file.Close()

		// On Unix-like systems, explicitly set file permissions
		if runtime.GOOS != "windows" {
			if err := file.Chmod(0600); err != nil {
				return fmt.Errorf("failed to set file permissions: %w", err)
			}
		}

		data, err := yaml.Marshal(cm.config)
		if err != nil {
			return fmt.Errorf("failed to marshal config: %w", err)
		}

		if _, err := file.Write(data); err != nil {
			return fmt.Errorf("failed to write config: %w", err)
		}

		return nil
	}

	// SetRootKey saves the root key in plain text
	func (cm *ConfigManager) SetRootKey(rootKey string) error {
		cm.config.RootKey = rootKey
		cm.config.IsEncrypted = false
		// Clear any encrypted data
		cm.config.EncryptedRootKey = ""
		cm.config.Salt = nil

		return cm.saveConfig()
	}

	// GetRootKey retrieves the root key
	func (cm *ConfigManager) GetRootKey() (string, error) {
		if cm.config.IsEncrypted {
			// Handle encrypted case
			password, err := promptPassword("Enter password to decrypt config: ")
			if err != nil {
				return "", fmt.Errorf("failed to read password: %w", err)
			}

			key := deriveKey(password, cm.config.Salt)
			encrypted, err := base64.StdEncoding.DecodeString(cm.config.EncryptedRootKey)
			if err != nil {
				return "", fmt.Errorf("failed to decode encrypted data: %w", err)
			}

			return decrypt(key, encrypted)
		}

		return cm.config.RootKey, nil
	}

	// EnableEncryption converts plain text storage to encrypted storage
	func (cm *ConfigManager) EnableEncryption() error {
		if cm.config.IsEncrypted {
			return fmt.Errorf("encryption is already enabled")
		}

		if cm.config.RootKey == "" {
			return fmt.Errorf("no root key to encrypt")
		}

		password, err := promptPassword("Enter password to encrypt config: ")
		if err != nil {
			return fmt.Errorf("failed to read password: %w", err)
		}

		// Generate a random salt
		salt := make([]byte, saltLength)
		if _, err := rand.Read(salt); err != nil {
			return fmt.Errorf("failed to generate salt: %w", err)
		}

		// Derive encryption key and encrypt
		key := deriveKey(password, salt)
		encrypted, err := encrypt(key, cm.config.RootKey)
		if err != nil {
			return fmt.Errorf("encryption failed: %w", err)
		}

		// Update config with encrypted data
		cm.config.EncryptedRootKey = base64.StdEncoding.EncodeToString(encrypted)
		cm.config.Salt = salt
		cm.config.IsEncrypted = true
		cm.config.RootKey = "" // Clear plain text key

		return cm.saveConfig()
	}

	// DisableEncryption converts encrypted storage to plain text
	func (cm *ConfigManager) DisableEncryption() error {
		if !cm.config.IsEncrypted {
			return fmt.Errorf("encryption is not enabled")
		}

		// Get the decrypted key first
		rootKey, err := cm.GetRootKey()
		if err != nil {
			return fmt.Errorf("failed to decrypt key: %w", err)
		}

		// Clear encrypted data and store as plain text
		cm.config.EncryptedRootKey = ""
		cm.config.Salt = nil
		cm.config.IsEncrypted = false
		cm.config.RootKey = rootKey

		return cm.saveConfig()
	}

*/

/*
	package config

	import (
		"crypto/aes"
		"crypto/cipher"
		"crypto/rand"
		"encoding/base64"
		"encoding/json"
		"fmt"
		"golang.org/x/crypto/argon2"
		"io"
		"os"
		"syscall"
		"golang.org/x/term"
	)

	// SecureConfig holds our configuration data
	type SecureConfig struct {
		EncryptedRootKey string `json:"encrypted_root_key"`
		Salt            []byte `json:"salt"`
		// Add other fields as needed
	}

	// Parameters for Argon2 key derivation
	const (
		keyLength  = 32 // for AES-256
		saltLength = 16
		time       = 1
		memory     = 64 * 1024
		threads    = 4
	)

	// promptPassword securely prompts for password
	func promptPassword(prompt string) (string, error) {
		fmt.Print(prompt)
		password, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Add newline after input
		if err != nil {
			return "", err
		}
		return string(password), nil
	}

	// deriveKey generates an encryption key from a password using Argon2
	func deriveKey(password string, salt []byte) []byte {
		return argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	}

	// encrypt encrypts data using AES-GCM
	func encrypt(key []byte, plaintext string) ([]byte, error) {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return nil, err
		}

		ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
		return ciphertext, nil
	}

	// decrypt decrypts data using AES-GCM
	func decrypt(key []byte, ciphertext []byte) (string, error) {
		block, err := aes.NewCipher(key)
		if err != nil {
			return "", err
		}

		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return "", err
		}

		if len(ciphertext) < gcm.NonceSize() {
			return "", fmt.Errorf("ciphertext too short")
		}

		nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return "", err
		}

		return string(plaintext), nil
	}

	// EncryptConfig encrypts the root key with a password
	func EncryptConfig(rootKey string) error {
		password, err := promptPassword("Enter password to encrypt config: ")
		if err != nil {
			return fmt.Errorf("failed to read password: %v", err)
		}

		// Generate a random salt
		salt := make([]byte, saltLength)
		if _, err := rand.Read(salt); err != nil {
			return fmt.Errorf("failed to generate salt: %v", err)
		}

		// Derive encryption key from password
		key := deriveKey(password, salt)

		// Encrypt the root key
		encrypted, err := encrypt(key, rootKey)
		if err != nil {
			return fmt.Errorf("encryption failed: %v", err)
		}

		config := SecureConfig{
			EncryptedRootKey: base64.StdEncoding.EncodeToString(encrypted),
			Salt:            salt,
		}

		// Save to file
		file, err := os.OpenFile(".spike-config", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return fmt.Errorf("failed to open config file: %v", err)
		}
		defer file.Close()

		return json.NewEncoder(file).Encode(config)
	}

	// DecryptConfig decrypts and returns the root key
	func DecryptConfig() (string, error) {
		// Read config file
		file, err := os.Open(".spike-config")
		if err != nil {
			return "", fmt.Errorf("failed to open config file: %v", err)
		}
		defer file.Close()

		var config SecureConfig
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return "", fmt.Errorf("failed to decode config: %v", err)
		}

		// Get password from user
		password, err := promptPassword("Enter password to decrypt config: ")
		if err != nil {
			return "", fmt.Errorf("failed to read password: %v", err)
		}

		// Derive key from password and salt
		key := deriveKey(password, config.Salt)

		// Decode base64 encrypted data
		encrypted, err := base64.StdEncoding.DecodeString(config.EncryptedRootKey)
		if err != nil {
			return "", fmt.Errorf("failed to decode encrypted data: %v", err)
		}

		// Decrypt the root key
		rootKey, err := decrypt(key, encrypted)
		if err != nil {
			return "", fmt.Errorf("decryption failed: %v", err)
		}

		return rootKey, nil
	}

*/

/*
Create encrypted backup of the root key.

type KeyBackup struct {
    Version          int       `json:"version"`
    EncryptedKey     string    `json:"encrypted_key"`
    KeyID            string    `json:"key_id"`
    Timestamp        time.Time `json:"timestamp"`
    EncryptionParams struct {
        Algorithm string `json:"algorithm"`
        KDF      string `json:"kdf"`
        Salt     []byte `json:"salt"`
    } `json:"encryption_params"`
}

func createBackup(rootKey []byte, recoveryPassword string) (*KeyBackup, error) {
    // Generate strong encryption parameters
    salt := make([]byte, 32)
    rand.Read(salt)

    // Derive key from recovery password
    key := deriveKey(recoveryPassword, salt)

    // Encrypt root key
    encrypted, err := encrypt(key, rootKey)
    if err != nil {
        return nil, err
    }

    return &KeyBackup{
        Version:      1,
        EncryptedKey: base64.StdEncoding.EncodeToString(encrypted),
        KeyID:        generateKeyID(rootKey),
        Timestamp:    time.Now(),
        EncryptionParams: struct {
            Algorithm string `json:"algorithm"`
            KDF      string `json:"kdf"`
            Salt     []byte `json:"salt"`
        }{
            Algorithm: "AES-256-GCM",
            KDF:      "Argon2id",
            Salt:     salt,
        },
    }, nil
}


*/

/*
func recoverRootKey(backup *KeyBackup, recoveryPassword string) ([]byte, error) {
    // Derive key from recovery password using stored params
    key := deriveKey(recoveryPassword, backup.EncryptionParams.Salt)

    // Decode and decrypt
    encrypted, err := base64.StdEncoding.DecodeString(backup.EncryptedKey)
    if err != nil {
        return nil, err
    }

    rootKey, err := decrypt(key, encrypted)
    if err != nil {
        return nil, err
    }

    // Verify key ID matches
    if generateKeyID(rootKey) != backup.KeyID {
        return nil, errors.New("key verification failed")
    }

    return rootKey, nil
}
*/

/*
Use a strong recovery password (high entropy)
Consider splitting recovery password using Shamir's Secret Sharing
Regular testing of recovery process
Maintain audit log of backup access
Version control for backup format
Backup rotation strategy

*/
