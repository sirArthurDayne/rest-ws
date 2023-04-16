package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"github.com/sirArthurDayne/rest-ws/models"
	"github.com/sirArthurDayne/rest-ws/repository"
	"github.com/sirArthurDayne/rest-ws/server"
)

type UserSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := UserSignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}

		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserSignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
