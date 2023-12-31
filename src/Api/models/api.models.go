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

type Verify2FATokenRequest struct {
	AccessToken        string `json:"accessToken"`
	FriendlyDeviceName string `json:"friendlyDeviceName"`
	Session            string `json:"session"`
	UserCode           string `json:"userCode"`
}

type RespondChallengeRequest struct {
	ChallengeName string `json:"challengeName"`
	Session       string `json:"session"`
	UserName      string `json:"userName"`
	Token2FA      string `json:"token2FA"`
}
