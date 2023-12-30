package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	apiModels "price-tracker-authentication/src/Api/models"
)

func GetUserBodyRequest(r *http.Request, model any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error getting the body in request", err)
		return err
	}

	user := model
	err = json.Unmarshal(body, user)
	if err != nil {
		log.Fatal("Error decoding the body in request", err)
		return err
	}
	return nil
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

func SendSuccessResponse(w http.ResponseWriter, response any, statusCode int) {
	w.WriteHeader(statusCode)
	successResponse := &apiModels.SuccesResponse{
		Success:  true,
		Response: response,
	}
	errorResponse, _ := json.Marshal(successResponse)
	w.Write(errorResponse)
}
