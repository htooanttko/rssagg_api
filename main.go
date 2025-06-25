package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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
	v1Router.Post("/user", apiCfg.handlerUserCreate)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerUserGet))

	// feed
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Get("/feed", apiCfg.middlewareAuth(apiCfg.handlerFeedGet))
	
	// feed follow
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowGet))
	v1Router.Delete("/feed_follows/{FeedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	v1Router.Get("/posts",apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries,collectionConcurrency,collectionInterval)
	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}