package app

import (
	"cr-auth/dto"
	"cr-auth/logger"
	"cr-auth/service"
	"encoding/json"

	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "App Is Working Fine")
}

// Signup Api

func (h AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest dto.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
		logger.Error("Error while decoding signup request: " + err.Error())
		writeResponse(w, 400, err.Error())
	} else {
		err := h.service.Signup(signupRequest)
		if err != nil {
			writeResponse(w, http.StatusBadGateway, dto.NewSignupResponse("Signup Failed", "", &err.Message))

		} else {
			writeResponse(w, http.StatusOK, dto.NewSignupResponse("Signup Success", signupRequest.Username, nil))

		}
	}

}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, appErr := h.service.Login(loginRequest)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

/*
	Sample URL string

http://localhost:8181/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=2000&account_id=95470
*/
func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	// converting from Query to map type
	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		appErr := h.service.Verify(urlParams)
		if appErr != nil {
			writeResponse(w, appErr.Code, notAuthorizedResponse(appErr.Message))
		} else {
			writeResponse(w, http.StatusOK, authorizedResponse())
		}
	} else {
		writeResponse(w, http.StatusForbidden, notAuthorizedResponse("missing token"))
	}
}

func (h AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest dto.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, appErr := h.service.Refresh(refreshRequest)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
