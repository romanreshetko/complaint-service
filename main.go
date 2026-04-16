package main

import (
	DB "complaint-service/db"
	"complaint-service/handlers"
	"complaint-service/middlewares"
	"log"
	"net/http"
	"os"
)

func main() {
	cnf := DB.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := DB.ConnectWithRetry(cnf)
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := middlewares.LoadPublicKey("./keys/public.pem")
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware := middlewares.AuthMiddleware(publicKey)

	h := handlers.New(db)
	mux := http.NewServeMux()
	mux.Handle("complaint/create/review", authMiddleware(http.HandlerFunc(h.CreateReviewComplaintHandler)))
	handlerWithCors := middlewares.CorsMiddleware(mux)
	log.Println("Complaint service started on port 8080")
	log.Println(http.ListenAndServe(":8080", handlerWithCors))
}
