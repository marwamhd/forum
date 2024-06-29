package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"pl/handlers"
	"runtime"
)

func main() {
	fmt.Println("the server is starting at https://localhost:5050/")
	OpenBrowser("https://localhost:5050/")
	StartServer()
}

func StartServer() {

	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/addpost", handlers.AddPostHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServeTLS("0.0.0.0:5050", "server.pem", "server.key", nil)

	log.Fatal(err)
}

func OpenBrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
