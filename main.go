package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/htooanttko/rssagg_api/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main () {
	godotenv.Load()

 	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("Database conn error:",err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig {
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	// health check
	v1Router.Get("/healthz", handlerHealthz)
	v1Router.Get("/healthz/err", handlerErr)

	// user
	v1Router.Post("/user",apiCfg.handlerCreateUser)
	v1Router.Get("/user",apiCfg.handlerUserGet)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}
	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}