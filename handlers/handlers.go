package handlers

import (
	"encoding/json"
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

// CreateURL get FullURL and pass it to u.Repo.Create
func (u *URL) CreateURL(w http.ResponseWriter, r *http.Request) {
	var token = shortener.CreateToken(shortener.Alphabet, shortener.AlphabetLen, shortener.TokenLen)

	var body struct {
		FullURL string `json:"fullURL"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shorten := model.ShortenURL{
		ID:      rand.Uint64(),
		Token:   token,
		FullURL: body.FullURL,
	}

	err := u.Repo.Create(r.Context(), shorten)
	if err != nil {
		fmt.Println("failed to create:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(shorten)
	if err != nil {
		fmt.Println("failed to marshal while creating:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(res); err != nil {
		fmt.Println("error while writing:", err)
	}
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

	if err := json.NewEncoder(w).Encode(url); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	fmt.Fprintf(w, `create new url page`)
}
