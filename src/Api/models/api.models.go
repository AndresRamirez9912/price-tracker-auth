package apiModels

import "price-tracker-authentication/src/models"

type GeneralResponse struct {
	Success      bool        `json:"success"`
	Response     interface{} `json:"response"`
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
}

type ChangePasswordRequest struct {
	UserInformation *models.UserCredentials `json:"userInformation"`
	NewPassword     string                  `json:"newPassword"`
}

type GetUserRequest struct {
	AccessToken string `json:"AccessToken"`
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
