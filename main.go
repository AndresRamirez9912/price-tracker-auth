package main

import (
	"log"
	"net/http"
	"price-tracker-authentication/src/Api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Create router
	r := chi.NewRouter()
	r.Post("/signUp", handlers.HandleSignUpUser)
	r.Post("/logIn", handlers.HandleLogInUser)
	r.Post("/verifyUser", handlers.HandleVerifyUser)
	r.Post("/changePassword", handlers.HandleChangePassword)
	r.Post("/reSendVerificationCode", handlers.HandleReSendVerificationCode)
	r.Post("/signOut", handlers.HandleSignOut)
	r.Post("/set2FA", handlers.HandleSet2FA)
	r.Post("/associateToken", handlers.HandleAssociateSoftwareToken)
	r.Post("/verifyToken", handlers.HandleVerifyToken)
	r.Post("/respondChallenge", handlers.HandleRespondChallenge)

	//Start server
	log.Println("Server starting at port 3000")
	http.ListenAndServe(":3000", r)
}
