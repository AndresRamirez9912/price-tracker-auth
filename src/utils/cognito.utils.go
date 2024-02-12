package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"price-tracker-authentication/src/constants"
	"price-tracker-authentication/src/models"
)

func CreateSecretHash(userInformation *models.UserCredentials) string {
	secret := os.Getenv(constants.SECRET_HASH)
	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(userInformation.UserName + os.Getenv(constants.COGNITO_APPCLIENT_ID)))
	hmacResult := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmacResult)
}
