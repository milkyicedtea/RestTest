package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func HandleDbReadTest(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
	}
	db, err := GetPgConn()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer db.Release()

	var user ReadUser
	err = db.QueryRow(
		context.Background(),
		"SELECT username, email FROM users WHERE id = $1",
		userId,
	).Scan(&user.Username, &user.Email)
	response, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Write(response)
}

func HandleDbWriteTest(w http.ResponseWriter, r *http.Request) {
	var user ReadUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	db, err := GetPgConn()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer db.Release()
	_, err = db.Exec(
		context.Background(),
		"INSERT INTO users (username, email) VALUES ($1, $2)",
		user.Username,
		user.Email,
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}
