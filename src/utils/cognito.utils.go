package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"price-tracker-authentication/src/models"
)

func CreateSecretHash(userInformation *models.UserCredentials) string {
	secret := "lgp5bmngniq9hbno40qoob0db3oh5nn9pmgqeiamp2ebh4tc0jn"
	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(userInformation.UserName + "5k1nhiok061928quq6ql8lcg96"))
	hmacResult := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmacResult)
}
