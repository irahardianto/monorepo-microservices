package hasher

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 will hash plain text with SHA 256
func SHA256(data string) string {
	hs := sha256.New()
	hs.Write([]byte(data))

	return fmt.Sprintf("%x", hs.Sum(nil))
}
