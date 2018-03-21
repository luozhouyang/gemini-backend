package token

import (
	"log"
	"io/ioutil"
)

func GetToken() string {
	b, err := ioutil.ReadFile("/home/allen/gemini_backend_user_token.txt")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(b)
}
