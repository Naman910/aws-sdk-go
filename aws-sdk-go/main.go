package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	ctx       = context.Background()
	rdb       *redis.Client
	logrusLog = logrus.New()
)

func init() {
	logrusLog.Formatter = &logrus.JSONFormatter{}
	logrusLog.Level = logrus.InfoLevel
}

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		logrusLog.Fatal("REDIS_ADDR environment variable is not set")
	}

	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // set if needed
		DB:       0,
	})

	router := mux.NewRouter()
	router.HandleFunc("/qa", CreateQA).Methods("POST")
	router.HandleFunc("/search", SearchQA).Methods("GET")

	logrusLog.Info("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
