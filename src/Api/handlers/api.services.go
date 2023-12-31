package handlers

import (
	"net/http"
	apiModels "price-tracker-authentication/src/Api/models"
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

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, signUpResponse := cognitoClient.SignUp(user)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
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

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, logInResponse := cognitoClient.LogIn(user)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, logInResponse, http.StatusAccepted)
}

func HandleVerifyUser(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("userName")
	confirmationCode := r.URL.Query().Get("confirmationCode")
	user := &models.UserCredentials{
		UserName: userName,
	}

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, confirmResponse := cognitoClient.ConfirmUser(user, confirmationCode)

	w.Header().Add("content-Type", "application/json")
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

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, changePasswordResponse := cognitoClient.ChangePassword(changePassword.UserInformation, changePassword.NewPassword)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, changePasswordResponse, http.StatusOK)
}

func HandleReSendVerificationCode(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("userName")
	user := &models.UserCredentials{
		UserName: userName,
	}

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, reSendResponse := cognitoClient.ReSendConfirmationCode(user)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, reSendResponse, http.StatusOK)
}

func HandleSignOut(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("accessToken")

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, signOutResponse := cognitoClient.SignOut(accessToken)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, signOutResponse, http.StatusOK)
}

func HandleSet2FA(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("accessToken")

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, set2FAResponse := cognitoClient.Set2FAPreference(accessToken)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, set2FAResponse, http.StatusOK)
}

func HandleAssociateSoftwareToken(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("accessToken")

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, associateResponse := cognitoClient.AssociateSoftwareToken(accessToken)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, associateResponse, http.StatusOK)
}

func HandleVerifyToken(w http.ResponseWriter, r *http.Request) {
	verifyToken := &apiModels.Verify2FAToken{}
	err := utils.GetUserBodyRequest(r, verifyToken)
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, associateResponse := cognitoClient.Verify2FAToken(verifyToken)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, associateResponse, http.StatusOK)
}

func HandleRespondChallenge(w http.ResponseWriter, r *http.Request) {
	respondChallenge := &apiModels.RespondChallenge{}
	err := utils.GetUserBodyRequest(r, respondChallenge)
	defer r.Body.Close()

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, challengeResponse := cognitoClient.Respond2FAChallenge(respondChallenge)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, challengeResponse, http.StatusOK)
}
