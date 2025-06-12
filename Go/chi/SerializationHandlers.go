package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleUserSerialization(w http.ResponseWriter, r *http.Request) {
	user, err := json.Marshal(StaticUser{
		Id:       1,
		Username: "JohnDoe",
		Email:    "johndoe@gmail.com",
		IsActive: true,
		Roles:    []string{"user", "admin"},
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(user)
}
