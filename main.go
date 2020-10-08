package main

import (
	"go-postgres/router"
	"log"
	"net/http"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./go-postgres.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
}

func main() {
	r := router.Router()
	log.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
