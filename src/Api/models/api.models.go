package apiModels

import "price-tracker-authentication/src/models"

type ErrorResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
}

type SuccesResponse struct {
	Success  bool
	Response interface{}
}

type ChangePasswordRequest struct {
	UserInformation *models.UserCredentials `json:"userInformation"`
	NewPassword     string                  `json:"newPassword"`
}
