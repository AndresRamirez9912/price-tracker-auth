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
