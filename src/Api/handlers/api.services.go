package handlers

import (
	"net/http"
	"os"
	apiModels "price-tracker-authentication/src/Api/models"
	"price-tracker-authentication/src/constants"
	"price-tracker-authentication/src/models"
	cognitoServices "price-tracker-authentication/src/services"
	"price-tracker-authentication/src/utils"
)

func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user := &models.UserCredentials{}
	err := utils.GetUserBodyRequest(r, user)
	if err != nil {
		return
	}

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, signUpResponse := cognitoClient.SignUp(user)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, signUpResponse, http.StatusCreated)
}

func HandleLogInUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user := &models.UserCredentials{}
	err := utils.GetUserBodyRequest(r, user)
	if err != nil {
		return
	}

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, logInResponse := cognitoClient.LogIn(user)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, logInResponse, http.StatusAccepted)
}

func HandleVerifyUser(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get(constants.USER_NAME)
	confirmationCode := r.URL.Query().Get(constants.CONFIRMATION_CODE)
	user := &models.UserCredentials{
		UserName: userName,
	}

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, confirmResponse := cognitoClient.ConfirmUser(user, confirmationCode)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, confirmResponse, http.StatusOK)
}

func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	changePassword := &apiModels.ChangePasswordRequest{}
	err := utils.GetUserBodyRequest(r, changePassword)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, changePasswordResponse := cognitoClient.ChangePassword(changePassword.UserInformation, changePassword.NewPassword)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, changePasswordResponse, http.StatusOK)
}

func HandleReSendVerificationCode(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get(constants.USER_NAME)
	user := &models.UserCredentials{
		UserName: userName,
	}

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, reSendResponse := cognitoClient.ReSendConfirmationCode(user)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, reSendResponse, http.StatusOK)
}

func HandleSignOut(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get(constants.ACCESS_TOKEN)

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, signOutResponse := cognitoClient.SignOut(accessToken)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, signOutResponse, http.StatusOK)
}

func HandleSet2FA(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get(constants.ACCESS_TOKEN)

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, set2FAResponse := cognitoClient.Set2FAPreference(accessToken)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, set2FAResponse, http.StatusOK)
}

func HandleAssociateSoftwareToken(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get(constants.ACCESS_TOKEN)

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, associateResponse := cognitoClient.AssociateSoftwareToken(accessToken)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, associateResponse, http.StatusOK)
}

func HandleVerifyToken(w http.ResponseWriter, r *http.Request) {
	verifyToken := &apiModels.Verify2FATokenRequest{}
	err := utils.GetUserBodyRequest(r, verifyToken)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, associateResponse := cognitoClient.Verify2FAToken(verifyToken)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, associateResponse, http.StatusOK)
}

func HandleRespondChallenge(w http.ResponseWriter, r *http.Request) {
	respondChallenge := &apiModels.RespondChallengeRequest{}
	err := utils.GetUserBodyRequest(r, respondChallenge)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, challengeResponse := cognitoClient.Respond2FAChallenge(respondChallenge)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, challengeResponse, http.StatusOK)
}

func HandleForgotPassword(w http.ResponseWriter, r *http.Request) {
	respondChallenge := &models.UserCredentials{}
	err := utils.GetUserBodyRequest(r, respondChallenge)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, forgotPasswordResponse := cognitoClient.ForgotPassword(respondChallenge)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, forgotPasswordResponse, http.StatusOK)
}

func HandleConfirmForgotPassword(w http.ResponseWriter, r *http.Request) {
	confirmationCode := r.URL.Query().Get(constants.CONFIRMATION_CODE)
	respondForget := &apiModels.ChangePasswordRequest{}
	err := utils.GetUserBodyRequest(r, respondForget)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, ConfirmforgotPasswordResponse := cognitoClient.ConfirmForgotPassword(respondForget, confirmationCode)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, ConfirmforgotPasswordResponse, http.StatusOK)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	getUserBody := &apiModels.GetUserRequest{}
	err := utils.GetUserBodyRequest(r, getUserBody)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient(os.Getenv(constants.AWS_COGNITO_REGION), os.Getenv(constants.COGNITO_APPCLIENT_ID))
	err, getUserResponse := cognitoClient.GetUserByJWT(getUserBody.AccessToken)

	w.Header().Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, getUserResponse, http.StatusOK)
}
