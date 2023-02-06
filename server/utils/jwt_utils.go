package utils

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWTAuthFunc func(next http.Handler, secret []byte) http.Handler

func verifyJWTToken(tokenString string, signingKey []byte) (jwt.Claims, error) {
	// TODO: verify jwt here
	return nil, nil
}

func JWTAuthMiddleware(next http.Handler, secret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Missing Authorization Header"))
			if err != nil {
				return
			}
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// TODO: validate JWT string here
		//claims, err := verifyJWTToken(tokenString, secret)
		//if err != nil {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	w.Write([]byte("Error verifying JWT token: " + err.Error()))
		//	return
		//}
		//
		//valid := true
		//if !valid {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	w.Write([]byte("Error missing required claims in JWT: " + err.Error()))
		//	return
		//}

		next.ServeHTTP(w, r)
	})
}
