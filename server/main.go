package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/handlers"
	"github.com/semperprimum/devlinks/server/middleware"
	"log"
	"net/http"
)

func main() {
	db.InitDB()

	db.CreateDB()

	defer db.DB.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("POST /login", handlers.LoginUser)
	router.HandleFunc("POST /register", handlers.CreateUser)

	router.HandleFunc("GET /links/{id}", handlers.GetUserLinks)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /user", handlers.GetUser)
	authRouter.HandleFunc("PUT /user/update", handlers.UpdateUserInfo)
	authRouter.HandleFunc("POST /user/link", handlers.AddLink)
	authRouter.HandleFunc("PUT /user/link", handlers.UpdateLinksOrder)
	authRouter.HandleFunc("DELETE /user/link/{id}", handlers.DeleteLink)

	router.Handle("/", middleware.Auth(authRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
