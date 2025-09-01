package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/dfodeker/learn-go/baby-server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	db             *database.Queries
	platform       string
	signingKey     string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_Url Must BE SET")

	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM Must BE SET")
	}
	signingKey := os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("SIGNING_KEY Must BE SET")
	}

	mux := http.NewServeMux()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error Loading DB, %s", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	apiCFg := apiConfig{
		db:       dbQueries,
		platform: platform,
	}
	fileServer := http.FileServer(http.Dir("."))

	mux.Handle("/app/", apiCFg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	mux.HandleFunc("GET /admin/metrics", apiCFg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCFg.handlerReset)

	mux.HandleFunc("POST /api/users", apiCFg.CreateUserHandler)
	mux.HandleFunc("POST /api/login", apiCFg.UserLoginHandler)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/chirps", apiCFg.CreateChirpHandler)
	mux.HandleFunc("GET /api/chirps", apiCFg.GetAllChirpsHandler)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCFg.getChirpByID)

	// mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
