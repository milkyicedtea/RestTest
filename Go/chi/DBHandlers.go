package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strconv"
)

type ReadUser struct {
	//Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func HandleDbReadTest(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
	}
	db, err := GetPgConn()
	if err != nil {
		http.Error(w, `{"error": "Invalid ID"}`, http.StatusBadRequest)
		return
	}
	defer db.Release()

	var user ReadUser
	err = db.QueryRow(
		context.Background(),
		"SELECT username, email FROM users WHERE id = $1",
		userId,
	).Scan(&user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		} else {
			log.Println(err)
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
