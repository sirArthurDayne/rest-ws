package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/segmentio/ksuid"
	"github.com/sirArthurDayne/rest-ws/helpers"
	"github.com/sirArthurDayne/rest-ws/models"
	"github.com/sirArthurDayne/rest-ws/repository"
	"github.com/sirArthurDayne/rest-ws/server"
)

type PostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InserPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.ValidateJwtAuthToken(s, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// decode and validate post request before writing to DB
			postRequest := PostRequest{}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer func() {
				err = r.Body.Close()
				if err != nil {
					fmt.Println(err)
				}
			}()
			jsonErr := json.Unmarshal(body, &postRequest)
			if jsonErr != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// generate random id for post
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// model representation
			postObj := models.Post{
				Id:          id.String(),
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			// write to database
			err = repository.InsertPost(r.Context(), &postObj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// send json response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:          postObj.Id,
				PostContent: postObj.PostContent,
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
