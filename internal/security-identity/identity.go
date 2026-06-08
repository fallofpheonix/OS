package identity

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

// Identity represents a validator's cryptographic identity.
type Identity struct {
	PrivateKey ed25519.PrivateKey `json:"private_key"`
	PublicKey  ed25519.PublicKey  `json:"public_key"`
}

// EncryptedContainer defines the on-disk format for secured identities.
type EncryptedContainer struct {
	Salt   []byte `json:"salt"`
	Nonce  []byte `json:"nonce"`
	Cipher []byte `json:"cipher"`
}

// GenerateNewIdentity creates a fresh Ed25519 keypair.
func GenerateNewIdentity() (*Identity, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	return &Identity{
		PrivateKey: priv,
		PublicKey:  pub,
	}, nil
}

// SaveEncrypted stores the identity to a file encrypted with a passphrase.
// CONSENSUS-015 Hardening: Argon2id + XChaCha20-Poly1305.
func (id *Identity) SaveEncrypted(path string, passphrase string) error {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return err
	}

	// Derive key using Argon2id
	// Parameters: Time=3, Memory=64MB, Threads=4
	key := argon2.IDKey([]byte(passphrase), salt, 3, 64*1024, 4, 32)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	plaintext, err := json.Marshal(id)
	if err != nil {
		return err
	}

	ciphertext := aead.Seal(nil, nonce, plaintext, nil)

	container := EncryptedContainer{
		Salt:   salt,
		Nonce:  nonce,
		Cipher: ciphertext,
	}

	data, err := json.Marshal(container)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

// LoadEncrypted retrieves and decrypts an identity from a file.
func LoadEncrypted(path string, passphrase string) (*Identity, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var container EncryptedContainer
	if err := json.Unmarshal(data, &container); err != nil {
		return nil, err
	}

	// Derive key using same salt and parameters
	key := argon2.IDKey([]byte(passphrase), container.Salt, 3, 64*1024, 4, 32)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	plaintext, err := aead.Open(nil, container.Nonce, container.Cipher, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %v (incorrect passphrase?)", err)
	}

	var id Identity
	if err := json.Unmarshal(plaintext, &id); err != nil {
		return nil, err
	}

	return &id, nil
}

// ExportPublicKey returns the hex-encoded public key.
func (id *Identity) ExportPublicKey() string {
	return hex.EncodeToString(id.PublicKey)
}
