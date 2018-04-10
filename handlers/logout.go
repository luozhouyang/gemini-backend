package handlers

import (
	"net/http"
	"backend/models"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gemini_cookie")
	if err != http.ErrNoCookie {
		sess := models.Session{Uuid: cookie.Value}
		sess.DeleteByUUID()
	}
}
