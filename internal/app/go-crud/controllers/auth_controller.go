package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tk-yuet/go-crud/internal/app/go-crud/encrypt"
	"github.com/tk-yuet/go-crud/internal/app/go-crud/models"
)

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type signInResponse struct {
	Token string `json:"token"`
}

func (env *ControllerEnv) SignIn(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	pw := r.FormValue("password")

	um := models.NewUserModel(env.DbClient, nil)
	um.FindByUsername(username)
	user := um.User
	if user == nil {
		fmt.Fprintf(w, "Username or password Incorrect")
		return
	}

	if encrypt.IsCorrectPassword(user.Password, pw) {
		token := encrypt.GenerateNewToken(1)

		resp := signInResponse{
			Token: token,
		}
		var jsonData []byte
		jsonData, err := json.Marshal(resp)
		if err != nil {
			log.Println("err: ", err)
		}
		fmt.Fprintf(w, "%s", string(jsonData))
		return
	}
	fmt.Fprintf(w, "Username or password Incorrect")

}

type signUpResponse struct {
	Success bool `json:"success"`
}

func (env *ControllerEnv) SignUp(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	pw := r.FormValue("password")

	newUser := models.NewUser(-1, username, encrypt.EncrpyptPw(pw))
	um := models.NewUserModel(env.DbClient, newUser)
	userId := um.Insert()

	resp := signUpResponse{
		Success: userId > 0,
	}

	var jsonData []byte
	jsonData, err := json.Marshal(resp)
	if err != nil {
		log.Println("err: ", err)
	}
	fmt.Fprintf(w, "%s", string(jsonData))
}
