package models

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserCredentials struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Locale   string `json:"lacale"`
}

type CognitoClient interface {
	SignUp(userInformation *UserCredentials) (error, string)
	ConfirmSignUp(userInformation *UserCredentials, verificationCode string) (error, bool)
	SignIn(userInformation *UserCredentials) (error, *cognito.InitiateAuthOutput)
}
