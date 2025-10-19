package utils

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuthClient *auth.Client

func InitFirebase() {
	// ambil JSON dari environment variable
	jsonKey := os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")
	if jsonKey == "" {
		log.Fatal("GOOGLE_SERVICE_ACCOUNT_JSON is not set")
	}

	// buat credential dari JSON string
	opt := option.WithCredentialsJSON([]byte(jsonKey))

	// inisialisasi Firebase app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	// dapatkan Auth client
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting auth client: %v", err)
	}

	FirebaseAuthClient = client
	log.Println("Firebase Auth initialized successfully")
}
