package main

import (
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	// reference: https://golangbyexample.com/http-basic-auth-golang/

	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if u != username {
		fmt.Printf("Username provided is wrong: %s\n", u)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if p != password {
		fmt.Printf("Password provided is wrong: %s\n", p)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Printf("Username: %s\n", u)
	fmt.Printf("Password: %s\n", p)
	w.WriteHeader(http.StatusOK)
	return
}
