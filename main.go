package main

import (
	"fmt"
	"log"
	"net/http"
	"pl/handlers"
)

func main() {
	fmt.Println("the server is starting at https://localhost:4343/")
	StartServer()
}

func StartServer() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServeTLS(":4343", "server.pem", "server.key", nil)
	log.Fatal(err)
}
