package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

const credentialsHashPassword = "e3d4b909-fded-489f-99b5-ecab0c04e989"

// GetPasswordHash returns a signed hash of the requested password
func GetPasswordHash(password string) string {
	claimsUUID, _ := uuid.FromString(credentialsHashPassword)
	passwordHash := uuid.NewV5(claimsUUID, password)
	return passwordHash.String()
}

func NewShopSecretToken() string {
	return GetRandomString(6)
}

func GetRandomString(length int) string {
	uuidStr := uuid.NewV4().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	uuidStr = strings.ToUpper(uuidStr)
	return uuidStr[len(uuidStr)-length:]
}
