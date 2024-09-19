package middlewares

import "net/http"

//func VerifyGet(next http.Handler) http.Handler {
//	return RestMethodCheckerMiddleware(next, http.MethodGet)
//}

func VerifyPost(next http.Handler) http.Handler {
	return RestMethodCheckerMiddleware(next, http.MethodPost)
}

func VerifyDelete(next http.Handler) http.Handler {
	return RestMethodCheckerMiddleware(next, http.MethodDelete)
}

func VerifyPatch(next http.Handler) http.Handler {
	return RestMethodCheckerMiddleware(next, http.MethodPatch)
}

func RestMethodCheckerMiddleware(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
