package controllers

import (
	"backend/services"
	"net/http"
)

func AddNewPassword(w http.ResponseWriter, r *http.Request) {
	response, httpStatusCode := services.AddNewPassword(r)

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte(response))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
