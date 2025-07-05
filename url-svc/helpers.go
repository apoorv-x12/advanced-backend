package url_svc

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomBase64String returns a URL-safe base64 string of 6 random bytes (8 chars).
func GenerateRandomBase64String() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Use RawURLEncoding to avoid padding and URL-unsafe characters.
	return base64.RawURLEncoding.EncodeToString(b), nil
}
