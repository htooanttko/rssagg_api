package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load()

 	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
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
	v1Router.Get("/healthz", handlerHealthz)
	v1Router.Get("/healthz/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}
	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}