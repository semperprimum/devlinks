package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/semperprimum/devlinks/server/db"
	"github.com/semperprimum/devlinks/server/handlers"
	"github.com/semperprimum/devlinks/server/middleware"
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

	router.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /user", handlers.GetUser)
	authRouter.HandleFunc("PUT /user/update", handlers.UpdateUserInfo)
	authRouter.HandleFunc("POST /user/link", handlers.AddLink)
	authRouter.HandleFunc("PUT /user/link", handlers.UpdateLinksOrder)
	authRouter.HandleFunc("DELETE /user/link/{id}", handlers.DeleteLink)
	authRouter.HandleFunc("POST /user/upload", handlers.UploadPicture)

	router.Handle("/", middleware.Auth(authRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	fmt.Println(`
|\   ____\|\   __  \         
\ \  \___|\ \  \|\  \        
 \ \  \  __\ \  \\\  \       
  \ \  \|\  \ \  \\\  \      
   \ \_______\ \_______\     
    \|_______|\|_______|     
 ________  ________  ___     
|\   __  \|\   __  \|\  \    
\ \  \|\  \ \  \|\  \ \  \   
 \ \   __  \ \   ____\ \  \  
  \ \  \ \  \ \  \___|\ \  \ 
   \ \__\ \__\ \__\    \ \__\
    \|__|\|__|\|__|     \|__|`)
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
