package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	apiModels "price-tracker-authentication/src/Api/models"
	"price-tracker-authentication/src/models"
)

func GetUserBodyRequest(r *http.Request) (*models.UserCredentials, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error getting the body in SignUp request", err)
		return nil, err
	}

	user := &models.UserCredentials{}
	err = json.Unmarshal(body, user)
	if err != nil {
		log.Fatal("Error decoding the body in SignUp request", err)
		return nil, err
	}
	return user, nil
}

func SendErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	failResponse := &apiModels.ErrorResponse{
		ErrorCode:    400,
		ErrorMessage: err.Error(),
		Success:      false,
	}
	errorResponse, _ := json.Marshal(failResponse)
	w.Write(errorResponse)
	return
}

func SendSuccessResponse(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusCreated)
	successResponse := &apiModels.SuccesResponse{
		Success:  true,
		Response: response,
	}
	errorResponse, _ := json.Marshal(successResponse)
	w.Write(errorResponse)
}
