package handlers

import (
	"log"
	"net/http"
	"price-tracker-authentication/src/models"
	cognitoServices "price-tracker-authentication/src/services"
	"price-tracker-authentication/src/utils"
)

func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user, err := utils.GetUserBodyRequest(r)
	if err != nil {
		return
	}

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, signUpResponse := cognitoClient.SignUp(user)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		log.Println("Error creating the user", err)
		utils.SendErrorResponse(w, err)
	}

	utils.SendSuccessResponse(w, signUpResponse)
}

func HandleLogInUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user, err := utils.GetUserBodyRequest(r)
	if err != nil {
		return
	}

	cognitoClient := cognitoServices.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")
	err, logInResponse := cognitoClient.LogIn(user)

	w.Header().Add("content-Type", "application/json")
	if err != nil {
		log.Println("Error login the user", err)
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, logInResponse)
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
		log.Println("Error verifying the user", err)
		utils.SendErrorResponse(w, err)
		return
	}

	utils.SendSuccessResponse(w, confirmResponse)
}
