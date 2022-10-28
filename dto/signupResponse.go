package dto

type SignupResponse struct {
	Message  string  `json:"message"`
	Username string  `json:"username"`
	Error    *string `json:"error"`
}

func NewSignupResponse(message, username string, error *string) *SignupResponse {
	return &SignupResponse{
		Message:  message,
		Username: username,
		Error:    error,
	}

}