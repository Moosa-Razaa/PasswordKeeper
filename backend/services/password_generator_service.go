package services

import (
	"backend/apis"
	"net/http"
)

func PasswordGeneratorService(w http.ResponseWriter, r *http.Request) {
	response, httpStatusCode := apis.GeneratePassword(r)

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte(response))
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
