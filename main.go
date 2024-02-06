package main

import (
	"log"
	"net/http"
	"price-tracker-authentication/src/Api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Create router
	r := chi.NewRouter()

	// Disable CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/api/getUser", handlers.HandleGetUser)
	r.Post("/api/signUp", handlers.HandleSignUpUser)
	r.Post("/api/logIn", handlers.HandleLogInUser)
	r.Post("/api/verifyUser", handlers.HandleVerifyUser)
	r.Post("/api/changePassword", handlers.HandleChangePassword)
	r.Post("/api/reSendVerificationCode", handlers.HandleReSendVerificationCode)
	r.Post("/api/signOut", handlers.HandleSignOut)
	r.Post("/api/set2FA", handlers.HandleSet2FA)
	r.Post("/api/associateToken", handlers.HandleAssociateSoftwareToken)
	r.Post("/api/verifyToken", handlers.HandleVerifyToken)
	r.Post("/api/respondChallenge", handlers.HandleRespondChallenge)
	r.Post("/api/forgotPassword", handlers.HandleForgotPassword)
	r.Post("/api/confirmForgotPassword", handlers.HandleConfirmForgotPassword)

	//Start server
	log.Println("Server starting at port 3001")
	http.ListenAndServe(":3001", r)
}
