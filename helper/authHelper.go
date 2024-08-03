package helper

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func VerifyPassword(password, hashedPassword string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	return hashedPassword == base64.StdEncoding.EncodeToString(h.Sum(nil))
}
