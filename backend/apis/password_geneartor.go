package apis

import (
	"log"
	"net/http"
)

func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	log.Printf("request coming from %s", r.RemoteAddr)
	_, _ = w.Write([]byte("password"))
}
