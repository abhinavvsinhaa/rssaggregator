package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abhinavvsinhaa/rssaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load environment")
	}

	port := os.Getenv("PORT")
	db := os.Getenv("DB_URL")

	if port == "" || db == "" {
		log.Fatal("Environment not configured properly")
	}

	conn, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.handleGetUser))

	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.handleCreateFeed))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.authMiddleware(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.authMiddleware(apiCfg.handleGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.authMiddleware(apiCfg.handleDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server starting on port %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
