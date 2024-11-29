package main

import (
	"database/sql"
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/dUPYeYE/go-htmx/internal/database"
	"github.com/dUPYeYE/go-htmx/views"
)

//go:embed static views/globals.css static/images/*
var staticFiles embed.FS

type config struct {
	db        *database.Queries
	jwtSecret string
}

/*
- GO-htmx API
- Description: This is a simple API for go-htmx project.
- Version: 0.0.1
*/
func main() {
	// =================================================================================================
	// 1. Load the .env file and read variables
	// =================================================================================================
	log.Println("Starting the server...")
	log.Println("Reading the .env file...")
	if err := godotenv.Load(); err != nil {
		log.Printf("Error while reading .env file: %v", err)
	}
	log.Println("Reading the .env file... Done!")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environmental variable is not set!")
	}

	cfg := config{}

	log.Println("Initializing database connection...")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("DB_URL environmental variable is not set!")
		log.Println("App running without CRUD endpoints!")
	} else {
		db, err := sql.Open("postgres", dbURL)
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
	// 3. Serve static files
	// =================================================================================================
	router.Handle("/htmx", templ.Handler(views.Index(nil)))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := staticFiles.Open("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if _, err = io.Copy(w, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	router.Get("/images/*", func(w http.ResponseWriter, r *http.Request) {
		content, err := staticFiles.ReadFile("static" + r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(content)
	})
	router.Get("/globals.css", func(w http.ResponseWriter, r *http.Request) {
		content, err := staticFiles.ReadFile("views/globals.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/css")
		w.Write(content)
	})
	// =================================================================================================
	// 4. Create a new API router and mount it to the main router
	// =================================================================================================
	apiRouter := chi.NewRouter()
	if cfg.db != nil {
		apiRouter.Post("/users", cfg.handlerCreateUser)
		apiRouter.Get("/users", cfg.handlerGetAllUsers)
		apiRouter.Get("/users/{id}", cfg.middlewareAuth(cfg.handlerGetOneUser))
		apiRouter.Delete("/users/{id}", cfg.middlewareAuth(cfg.handlerDeleteUser))

		apiRouter.Post("/auth/login", cfg.handlerLogin)
		apiRouter.Get("/auth/refresh", cfg.middlewareAuth(cfg.handlerRefreshToken))
		apiRouter.Post("/auth/revoke", cfg.middlewareAuth(cfg.handlerRevokeToken))
	}
	router.Mount("/api", apiRouter)
	// =================================================================================================
	// 5. Create a new static file server and mount it to the main router
	// =================================================================================================
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 50 * time.Second,
	}
	// =================================================================================================
	// 6. Start the server
	// =================================================================================================
	log.Printf("Server is running on port %s", port)
	log.Fatal(server.ListenAndServe())
}
