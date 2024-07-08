package main

import (
	"log"
	"net/http"

	"github.com/gofor-little/env"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/handlers"
	"github.com/semperprimum/devlinks/server/middleware"
)

func main() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	db.InitDB()

	db.CreateDB()

	defer db.DB.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("POST /login", handlers.LoginUser)
	router.HandleFunc("POST /register", handlers.CreateUser)

	router.HandleFunc("GET /links/{id}", handlers.GetUserLinks)

	router.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /user", handlers.GetUser)
	authRouter.HandleFunc("PUT /user/update", handlers.UpdateUserInfo)
	authRouter.HandleFunc("POST /user/link", handlers.AddLink)
	authRouter.HandleFunc("PUT /user/link", handlers.UpdateLinksOrder)
	authRouter.HandleFunc("DELETE /user/link/{id}", handlers.DeleteLink)
	authRouter.HandleFunc("POST /user/upload", handlers.UploadPicture)
	authRouter.HandleFunc("PUT /user/link/{id}", handlers.UpdateLink)

	router.Handle("/", middleware.Auth(authRouter))
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(corsHandler),
	}

	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
