package handlers

import (
	"net/http"
	"backend/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	name := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	user := &models.User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	if err := user.Create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
