package main

import (
	"go-auth0-web/platform/authenticator"
	"go-auth0-web/platform/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to init authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on port 3000")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
