package response

import "net/http"

func SendServerError(w http.ResponseWriter, err error) {
	SendJson(w, http.StatusInternalServerError, Basic{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func SendBadRequestError(w http.ResponseWriter, err error) {
	SendJson(w, http.StatusBadRequest, Basic{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	})
}

func SendInvalidCredentials(w http.ResponseWriter) {
	SendJson(w, http.StatusUnauthorized, Basic{
		Status:  http.StatusUnauthorized,
		Message: "Invalid credentials",
	})
}
