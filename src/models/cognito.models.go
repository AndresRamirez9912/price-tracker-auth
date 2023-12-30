package models

type UserCredentials struct {
	Name     string
	UserName string
	Email    string
	Password string
	Locale   string
}

type CognitoClient interface {
	SignUp(userInformation *UserCredentials) (error, string)
	ConfirmSignUp(userInformation *UserCredentials, verificationCode string) (error, bool)
}
