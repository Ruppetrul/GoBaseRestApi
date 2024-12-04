package main

import (
	_ "firstRest/database"
	"firstRest/handlers/api"
	"firstRest/handlers/web"
	"firstRest/workers"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	go workers.RegisterCoinGeckoWorker()

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/api/current", api.Current)
	http.HandleFunc("/robots.txt", web.Robots)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")

	if env == "" {
		log.Fatalf("APP_ENV not found.")
	} else {
		fmt.Printf("Run in ENV: %s\n", env)
	}
	switch env {
	case "production":
		certFile := "/etc/letsencrypt/live/crypto-visor.ru/cert.pem"
		keyFile := "/etc/letsencrypt/live/crypto-visor.ru/privkey.pem"

		if err := http.ListenAndServeTLS(":443", certFile, keyFile, nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	default:
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	}
}
