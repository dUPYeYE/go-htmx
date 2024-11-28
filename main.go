package main

import (
	"database/sql"
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/dUPYeYE/go-htmx/internal/database"
)

type config struct {
	db *database.Queries
}

var staticFiles embed.FS

/*
- GO-htmx API
- Description: This is a simple API for go-htmx project.
- Version: 0.0.1
*/
func main() {
	// =================================================================================================
	// 1. Load the .env file and read variables
	// =================================================================================================
	if err := godotenv.Load(); err != nil {
		log.Printf("Error while reading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environmantal variable is not set!")
	}

	cfg := config{}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("DB_URL environmental variable is not set!")
		log.Println("App running without CRUD endpoints!")
	} else {
		db, err := sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatalf("Error while creating database connection: %v", err)
		}
		dbQueries := database.New(db)
		cfg.db = dbQueries

		log.Println("Database connection established!")
	}
	// =================================================================================================
	// 2. Create a new router and set up CORS
	// =================================================================================================
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// =================================================================================================
	// 3. Create a new API router and mount it to the main router
	// =================================================================================================
	apiRouter := chi.NewRouter()
	if cfg.db != nil {
		// apiRouter.Post("/users", handerCreateUser)
	}
	router.Mount("/api", apiRouter)
	// =================================================================================================
	// 4. Create a new static file server and mount it to the main router
	// =================================================================================================
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 50 * time.Second,
	}
	// =================================================================================================
	// 5. Start the server
	// =================================================================================================
	log.Printf("Server is running on port %s", port)
	log.Fatal(server.ListenAndServe())
}
