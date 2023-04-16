package helpers

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/sirArthurDayne/rest-ws/models"
	"github.com/sirArthurDayne/rest-ws/server"
)

func ValidateJwtAuthToken(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JwtSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
        return nil, err
	}
	return token, nil
}
