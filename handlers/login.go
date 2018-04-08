package handlers

import (
	"net/http"
	"backend/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	name := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	user, err := models.UserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if user.Name == name && user.Password == models.Encrypt(password) {
		session, err := user.CreateSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{
			Name:     "gemini_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		return
	} else {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}
