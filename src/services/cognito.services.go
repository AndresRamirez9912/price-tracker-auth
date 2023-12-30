package services

import (
	"log"
	"price-tracker-authentication/src/models"
	"price-tracker-authentication/src/utils"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AwsCognitoClient struct {
	CognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	AppClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) *AwsCognitoClient {
	conf := &aws.Config{Region: aws.String(cognitoRegion)}
	sess, err := session.NewSession(conf)
	if err != nil {
		log.Panic("Fatal error creating aws session", err)
	}
	client := cognito.New(sess)

	return &AwsCognitoClient{
		CognitoClient: client,
		AppClientId:   cognitoAppClientID,
	}
}

func (cognitoClient *AwsCognitoClient) SignUp(userInformation *models.UserCredentials) (error, string) {
	secretHash := utils.CreateSecretHash(userInformation)

	// Create the signUp object
	user := &cognitoidentityprovider.SignUpInput{
		Username:   aws.String(userInformation.UserName),
		Password:   aws.String(userInformation.Password),
		ClientId:   aws.String(cognitoClient.AppClientId),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(userInformation.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(userInformation.Email),
			},
			{
				Name:  aws.String("locale"),
				Value: aws.String(userInformation.Locale),
			},
			{
				Name:  aws.String("updated_at"),
				Value: aws.String(strconv.FormatInt(time.Now().Unix(), 10)),
			},
		},
	}

	// SingUp the user
	result, err := cognitoClient.CognitoClient.SignUp(user)
	if err != nil {
		return err, ""
	}
	return nil, result.String()
}

func (cognitoClient *AwsCognitoClient) ConfirmSignUp(userInformation *models.UserCredentials, verificationCode string) (error, bool) {
	secretHash := utils.CreateSecretHash(userInformation)

	confirmUser := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(cognitoClient.AppClientId),
		SecretHash:       aws.String(secretHash),
		ConfirmationCode: aws.String(verificationCode),
		Username:         aws.String(userInformation.UserName),
	}
	_, err := cognitoClient.CognitoClient.ConfirmSignUp(confirmUser)
	if err != nil {
		return err, false
	}
	return nil, true
}
