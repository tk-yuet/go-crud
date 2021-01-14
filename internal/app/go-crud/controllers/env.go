package controllers

import (
	"io"
	"net/http"
	"strings"

	"github.com/tk-yuet/go-crud/internal/app/go-crud/database"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/encrypt"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/models"
)

type ControllerEnv struct {
	DbClient *database.MySqlClient
	User     *models.User
}

func NewControllerEnv() *ControllerEnv {
	return &ControllerEnv{
		DbClient: nil,
		User:     nil,
	}
}

func (env *ControllerEnv) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		splitToken := strings.Split(auth, "Jwt ")
		if len(splitToken) < 2 {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"unauthorized"}`)
			return
		}

		token := splitToken[1]
		userId := encrypt.GetUserIdByJwtToken(token)

		if userId == 0 {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"unauthorized"}`)
			return
		}

		um := models.NewUserModel(env.DbClient, nil)
		um.Select(userId)
		env.User = um.User

		if um.User == nil {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"unauthorized"}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}
