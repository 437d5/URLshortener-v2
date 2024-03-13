package handlers

import (
	"fmt"
	"net/http"
)

type URL struct{}

func (u *URL) CreateURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create url")
}

func (u *URL) GetURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the url")
}

func (u *URL) DeleteURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete the url")
}

func (u *URL) ShowURLList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show the url list")
}

func (u *URL) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `create new url page`)
}