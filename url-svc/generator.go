package url_svc

import (
	"crypto/rand"
	"encoding/base64"
)
// IDGenerator defines the interface for generating unique IDs
type IDGenerator interface {
	Generate() (string, error)
}

// RandomIDGenerator generates random base64 IDs
type RandomIDGenerator struct{}

// NewRandomIDGenerator creates a new random ID generator
func NewRandomIDGenerator() *RandomIDGenerator {
	return &RandomIDGenerator{}
}

// Generate creates a random 6-byte base64 string
func (g *RandomIDGenerator) Generate() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Use RawURLEncoding to avoid padding and URL-unsafe characters.
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// HashIDGenerator generates IDs based on URL hash (for consistent URLs)
type HashIDGenerator struct{}

// NewHashIDGenerator creates a new hash-based ID generator
func NewHashIDGenerator() *HashIDGenerator {
	return &HashIDGenerator{}
}

// Generate creates a hash-based ID (placeholder for future implementation)
func (g *HashIDGenerator) Generate() (string, error) {
	// This could use SHA-256 hash of the URL and take first 6 bytes
	// For now, we'll use the random generator
	return "", nil
}
