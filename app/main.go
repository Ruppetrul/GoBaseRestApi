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
	http.HandleFunc("/favicon.ico", web.Favicon)
	http.HandleFunc("/build_graph", web.BuildGraph)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var env string
	if env = os.Getenv("APP_ENV"); env == "" {
		log.Fatalf("APP_ENV not found.")
	} else {
		fmt.Printf("Run in ENV: %s\n", env)
	}
	switch env {
	case "production":
		baseUri := "/etc/letsencrypt/live/crypto-visor.ru/"
		if err := http.ListenAndServeTLS(":443", baseUri+"cert.pem", baseUri+"privkey.pem", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	default:
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	}
}
