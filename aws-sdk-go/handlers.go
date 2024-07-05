package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type QA struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func CreateQA(w http.ResponseWriter, r *http.Request) {
	var qa QA
	err := json.NewDecoder(r.Body).Decode(&qa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(qa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rdb.Set(ctx, qa.Question, data, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchQA(w http.ResponseWriter, r *http.Request) {
	question := r.URL.Query().Get("question")
	if question == "" {
		http.Error(w, "Missing question parameter", http.StatusBadRequest)
		return
	}

	answer, err := rdb.Get(ctx, question).Result()
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(answer))
}
