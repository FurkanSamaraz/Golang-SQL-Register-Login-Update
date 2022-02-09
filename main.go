package main

import (
	"golang/login"
	"golang/register"
	"golang/update"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/register", register.Register)
	mux.HandleFunc("/login", login.Login)
	mux.HandleFunc("/update", update.Update)
	http.ListenAndServe(":8080", mux)
}
