package main

import (
	"fmt"
	"log"
	"price-tracker-authentication/src/models"
	"price-tracker-authentication/src/services"
)

func main() {
	// Create Cognito Client
	cognitoClient := services.NewCognitoClient("us-east-2", "5k1nhiok061928quq6ql8lcg96")

	// Create my user object
	user := &models.UserCredentials{
		Name:     "Andres",
		UserName: "Andres",
		Email:    "andres.ramirez9912@gmail.com",
		Password: "1234Andres@",
		Locale:   "es_CO",
	}

	// SigUp the User
	err, result := cognitoClient.SignUp(user)
	if err != nil {
		log.Fatalln("Signin failed", err)
	}
	fmt.Println("SignIn success:", result)

	// Confirm the User
	err, confirmed := cognitoClient.ConfirmSignUp(user, "516313")
	if err != nil {
		log.Fatalln("Confirmation failed", err)
	}
	if confirmed {
		fmt.Println("The user was confirmed")
	} else {
		fmt.Println("The user could not confirm")
	}
}
