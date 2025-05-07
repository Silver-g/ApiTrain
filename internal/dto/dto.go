package dto

type ErrorResponse struct {
	ErrorCode string `json:"error"`
	Message   string `json:"message"`
}
type SuccessResponse struct {
	Message string `json:"message"`
}
