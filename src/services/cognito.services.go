package cognitoServices

import (
	"log"
	apiModels "price-tracker-authentication/src/Api/models"
	"price-tracker-authentication/src/constants"
	"price-tracker-authentication/src/models"
	"price-tracker-authentication/src/utils"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type awsCognitoClient struct {
	CognitoClient *cognito.CognitoIdentityProvider
	AppClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) *awsCognitoClient {
	conf := &aws.Config{Region: aws.String(cognitoRegion)}
	sess, err := session.NewSession(conf)
	if err != nil {
		log.Panic("Fatal error creating aws session", err)
	}
	client := cognito.New(sess)

	return &awsCognitoClient{
		CognitoClient: client,
		AppClientId:   cognitoAppClientID,
	}
}

func (cognitoClient *awsCognitoClient) SignUp(userInformation *models.UserCredentials) (error, *cognito.SignUpOutput) {
	secretHash := utils.CreateSecretHash(userInformation)

	user := &cognito.SignUpInput{
		Username:   aws.String(userInformation.UserName),
		Password:   aws.String(userInformation.Password),
		ClientId:   aws.String(cognitoClient.AppClientId),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String(constants.USER_ATTRIBUTE_NAME),
				Value: aws.String(userInformation.Name),
			},
			{
				Name:  aws.String(constants.USER_ATTRIBUTE_EMAIL),
				Value: aws.String(userInformation.Email),
			},
			{
				Name:  aws.String(constants.USER_ATTRIBUTE_LOCALE),
				Value: aws.String(constants.USER_ATTRIBUTE_ES_CO),
			},
			{
				Name:  aws.String(constants.USER_ATTRIBUTE_UPDATED_AT),
				Value: aws.String(strconv.FormatInt(time.Now().Unix(), 10)),
			},
		},
	}

	result, err := cognitoClient.CognitoClient.SignUp(user)
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (cognitoClient *awsCognitoClient) ConfirmUser(userInformation *models.UserCredentials, confirmationCode string) (error, bool) {
	secretHash := utils.CreateSecretHash(userInformation)

	confirmUser := &cognito.ConfirmSignUpInput{
		ClientId:         aws.String(cognitoClient.AppClientId),
		SecretHash:       aws.String(secretHash),
		ConfirmationCode: aws.String(confirmationCode),
		Username:         aws.String(userInformation.UserName),
	}
	_, err := cognitoClient.CognitoClient.ConfirmSignUp(confirmUser)
	if err != nil {
		log.Println("Error verifying the user", err)
		return err, false
	}
	return nil, true
}

func (cognitoClient *awsCognitoClient) LogIn(userInformation *models.UserCredentials) (error, *cognito.InitiateAuthOutput) {
	secretHash := utils.CreateSecretHash(userInformation)

	confirmUser := &cognito.InitiateAuthInput{
		AuthFlow: aws.String(constants.USER_PASSWORD_AUTH_FLOW),
		ClientId: aws.String(cognitoClient.AppClientId),
		AuthParameters: aws.StringMap(map[string]string{
			constants.USER_NAME_MAP: userInformation.UserName,
			constants.PASSWORD:      userInformation.Password,
			constants.SECRET_HASH:   secretHash,
		}),
	}
	signInResponse, err := cognitoClient.CognitoClient.InitiateAuth(confirmUser)
	if err != nil {
		log.Println("Error trying to login the user", err)
		return err, nil
	}
	return nil, signInResponse
}

func (cognitoClient *awsCognitoClient) ChangePassword(userInformation *models.UserCredentials, newPassword string) (error, bool) {
	err, logInResposne := cognitoClient.LogIn(userInformation)
	if err != nil {
		log.Println("Error Login the user in change password flow", err)
		return err, false
	}

	changePassword := &cognito.ChangePasswordInput{
		PreviousPassword: aws.String(userInformation.Password),
		ProposedPassword: aws.String(newPassword),
		AccessToken:      aws.String(*logInResposne.AuthenticationResult.AccessToken),
	}
	_, err = cognitoClient.CognitoClient.ChangePassword(changePassword)
	if err != nil {
		log.Println("Error changing the user's password", err)
		return err, false
	}
	return nil, true
}

func (cognitoClient *awsCognitoClient) ReSendConfirmationCode(userInformation *models.UserCredentials) (error, *cognito.ResendConfirmationCodeOutput) {
	secretHash := utils.CreateSecretHash(userInformation)

	resendConfirmation := &cognito.ResendConfirmationCodeInput{
		ClientId:   aws.String(cognitoClient.AppClientId),
		SecretHash: aws.String(secretHash),
		Username:   aws.String(userInformation.UserName),
	}
	response, err := cognitoClient.CognitoClient.ResendConfirmationCode(resendConfirmation)
	if err != nil {
		log.Println("Error re-sending the user's verification code", err)
		return err, nil
	}
	return nil, response
}

func (cognitoClient *awsCognitoClient) SignOut(accessToken string) (error, bool) {
	signOut := &cognito.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}
	_, err := cognitoClient.CognitoClient.GlobalSignOut(signOut)
	if err != nil {
		log.Println("Error signing out the user", err)
		return err, false
	}
	return nil, true
}

func (cognitoClient *awsCognitoClient) Set2FAPreference(accessToken string) (error, bool) {
	set2FA := &cognito.SetUserMFAPreferenceInput{
		AccessToken: aws.String(accessToken),
		SoftwareTokenMfaSettings: &cognito.SoftwareTokenMfaSettingsType{
			Enabled:      aws.Bool(true),
			PreferredMfa: aws.Bool(true),
		},
	}
	_, err := cognitoClient.CognitoClient.SetUserMFAPreference(set2FA)
	if err != nil {
		log.Println("Error setting the 2FA method", err)
		return err, false
	}
	return nil, true
}

func (cognitoClient *awsCognitoClient) AssociateSoftwareToken(accessToken string) (error, *cognito.AssociateSoftwareTokenOutput) {
	associateToken := &cognito.AssociateSoftwareTokenInput{
		AccessToken: aws.String(accessToken),
	}
	associateResponse, err := cognitoClient.CognitoClient.AssociateSoftwareToken(associateToken)
	if err != nil {
		log.Println("Error setting the 2FA method", err)
		return err, nil
	}
	return nil, associateResponse
}

func (cognitoClient *awsCognitoClient) Verify2FAToken(verifyInformation *apiModels.Verify2FAToken) (error, *cognito.VerifySoftwareTokenOutput) {
	verifyToken := &cognito.VerifySoftwareTokenInput{
		AccessToken:        aws.String(verifyInformation.AccessToken),
		UserCode:           aws.String(verifyInformation.UserCode),
		FriendlyDeviceName: aws.String(verifyInformation.FriendlyDeviceName),
		Session:            aws.String(verifyInformation.Session),
	}
	verifyResponse, err := cognitoClient.CognitoClient.VerifySoftwareToken(verifyToken)
	if err != nil {
		log.Println("Error verifying the 2FA token", err)
		return err, nil
	}
	return nil, verifyResponse
}

func (cognitoClient *awsCognitoClient) Respond2FAChallenge(challengeResponse *apiModels.RespondChallenge) (error, *cognito.RespondToAuthChallengeOutput) {
	userInformation := &models.UserCredentials{UserName: challengeResponse.UserName}
	secretHash := utils.CreateSecretHash(userInformation)

	authChallengeParameters := &cognito.RespondToAuthChallengeInput{
		Session:       aws.String(challengeResponse.Session),
		ChallengeName: aws.String(challengeResponse.ChallengeName),
		ChallengeResponses: aws.StringMap(map[string]string{
			constants.USER_NAME_MAP:           challengeResponse.UserName,
			constants.SOFTWARE_TOKEN_MFA_CODE: challengeResponse.Token2FA,
			constants.SECRET_HASH:             secretHash,
		}),
		ClientId: aws.String(cognitoClient.AppClientId),
	}
	authChallengeResponse, err := cognitoClient.CognitoClient.RespondToAuthChallenge(authChallengeParameters)
	if err != nil {
		log.Println("Error sending the 2FA challenge token", err)
		return err, nil
	}
	return nil, authChallengeResponse
}
