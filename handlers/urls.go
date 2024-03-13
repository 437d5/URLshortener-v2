package handlers

import (
	"fmt"
	"net/http"
)

type URL struct{}

func (u *URL) CreateURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an url")
}

func (u *URL) GetURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the url")
}

func (u *URL) DeleteURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete the url")
}

func (u *URL) ShowURLList(w http.ResponseController, r *http.Request) {
	fmt.Println("show the url list")
}
