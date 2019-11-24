package server

import (
	"net/http"
)

// Define our struct
type authenticatorMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticatorMiddleware) Populate() {
	amw.tokenUsers = make(map[string]string)
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticatorMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if _, found := amw.tokenUsers[token]; found {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
