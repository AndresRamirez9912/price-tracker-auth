package apiModels

type ErrorResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
}

type SuccesResponse struct {
	Success  bool
	Response interface{}
}
