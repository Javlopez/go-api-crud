package handler

import (
	repository "wiki/app/repository"
	"wiki/app/driver"
	post "wiki/app/repository/post"
	"net/http"
	"golang.org/x/net/context"
	"encoding/json"
	"wiki/app/models"
	"fmt"
	"strconv"
	"github.com/go-chi/chi"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post
type Post struct {
	repo repository.PostRepo
}

func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(context.Background(), 5)
	respondWithJSON(w, http.StatusOK, payload)
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}

	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(context.Background(), &post)
	fmt.Println(newID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

func (p *Post) Update(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	data := models.Post{ID: int64(id)}

	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(context.Background(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError,  "Server Error")
	} else {
		respondWithJSON(w, http.StatusOK, payload)
	}
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := p.repo.GetByID(context.Background(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(context.Background(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		respondWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
	}
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
