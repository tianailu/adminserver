package crypto

import (
	"crypto/sha256"
	"fmt"
)

func GetSha256String(source, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(source + salt))
	bytes := hash.Sum(nil)
	target := fmt.Sprintf("%x", bytes)

	return target
}
