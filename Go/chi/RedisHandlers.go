package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func HandleCacheUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userId := chi.URLParam(r, "id")
	cacheKey := "user:" + userId

	client := GetRedisClient()

	cached, err := client.Get(ctx, cacheKey).Result()
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cached))
		return
	}

	// Fallback to DB
	db, err := GetPgConn()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	defer db.Release()

	var user ReadUser
	err = db.QueryRow(
		ctx,
		"SELECT username, email FROM users WHERE id = $1",
		userId,
	).Scan(&user.Username, &user.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	redisClient.Set(ctx, cacheKey, userJson, 60*time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}
