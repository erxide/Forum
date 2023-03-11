package data

import (
	"crypto/sha256"
	"encoding/hex"
)

// CheckPasswordHash vérifie si le mot de passe en clair correspond au hash stocké dans la base de données
func CheckPasswordHash(password string, hash string) bool {
	testpassword, err := HashPassword(password)
	if err != nil {
		return false
	}
	if testpassword == hash {
		return true
	}
	return false

}

// HashPassword hashe le mot de passe en utilisant bcrypt
func HashPassword(password string) (string, error) {
	// Convertir le mot de passe en une série de bytes
	passwordBytes := []byte(password)

	// Calculer le hash SHA-256
	hashBytes := sha256.Sum256(passwordBytes)

	// Convertir le hash en une chaîne hexadécimale
	hashString := hex.EncodeToString(hashBytes[:])

	return hashString, nil
}
