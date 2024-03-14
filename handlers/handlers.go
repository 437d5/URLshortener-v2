package handlers

import (
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/437d5/URLshortener-v2/model"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"

	"github.com/437d5/URLshortener-v2/db"
	"github.com/437d5/URLshortener-v2/shortener"
)

type URL struct {
	Repo *db.RedisRepo
}

const Socket = "127.0.0.1:3000/urls/"

// CreateURL get FullURL and pass it to u.Repo.Create
func (u *URL) CreateURL(w http.ResponseWriter, r *http.Request) {
	var token = shortener.CreateToken(shortener.Alphabet, shortener.TokenLen, shortener.AlphabetLen)
	var fullURL = r.FormValue("url")

	shorten := model.ShortenURL{
		ID:      rand.Uint64(),
		Token:   token,
		FullURL: fullURL,
	}

	err := u.Repo.Create(r.Context(), shorten)
	if err != nil {
		fmt.Println("failed to create:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// res, err := json.Marshal(shorten)
	// if err != nil {
	// 	fmt.Println("failed to marshal while creating:", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	// if _, err := w.Write(res); err != nil {
	// 	fmt.Println("error while writing:", err)
	// }

	responseHTML := `<h2>URL Shortener</h2> 
	<p>The token we created: %s</p>
	<a href="/urls/%s" target="_blank">Your URL</a>
	<p>URL to copy: %s%s</p>
	<p>The URL you gave: %s</p>
	<p>Input the URL you want to be short.</p>
	<form method="post" action="/urls">
	<input type="text" name="url" placeholder="Enter a URL">
	<input type="submit" value="Short"></form>`
	// TODO: refactor short.Token + Socket pair
	fmt.Fprintf(w, responseHTML, shorten.Token, shorten.Token, Socket, shorten.Token,  shorten.FullURL)
}

func (u *URL) GetURL(w http.ResponseWriter, r *http.Request) {
	tokenParam := chi.URLParam(r, "token")

	url, err := u.Repo.GetURL(r.Context(), tokenParam)
	if errors.Is(err, db.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to get url:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fullURL := url.FullURL

	// if _, err := w.Write([]byte(fullURL)); err != nil {
	// 	fmt.Println("error while writing:", err)
	// }
	http.Redirect(w, r, fullURL, http.StatusSeeOther)
	//	if err := json.NewEncoder(w).Encode(url); err != nil {
	//		fmt.Println("failed to marshal:", err)
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
}

func (u *URL) DeleteURL(w http.ResponseWriter, r *http.Request) {
	tokenParam := chi.URLParam(r, "token")

	err := u.Repo.DeleteURL(r.Context(), tokenParam)
	if errors.Is(err, db.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("error while deleting:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *URL) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	responseHTML := `<h2>URL Shortener</h2> 
	<p>Input the URL you want to be short.</p>
	<form method="post" action="/urls">
	<input type="text" name="url" placeholder="Enter a URL">
	<input type="submit" value="Short"></form>`
	fmt.Fprintf(w, responseHTML)
}
