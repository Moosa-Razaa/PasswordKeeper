package controllers

import (
	"backend/services"
	"log"
	"net/http"
)

func CheckFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("request coming from %s", r.RemoteAddr)
	response, httpStatusCode := services.CheckFileExistence()

	if httpStatusCode != http.StatusOK {
		w.WriteHeader(httpStatusCode)
		_, _ = w.Write([]byte(response))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
