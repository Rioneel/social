package main

import (
	"net/http"
	"strconv"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/rioneel/social/internal/store"
)
type CreatePostPayload struct{
	Title string `json:"title"`
	Content string `json:"content"`
//	Image string `json:"image_path"`
//image addition
	Tags []string `json:"tags"`
}

func (app *application) CreatePostHandler(w http.ResponseWriter, r *http.Request){
	var payload CreatePostPayload
	if err := readJSON(w,r,&payload); err != nil{
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	

	post := &store.Post{
		Title: payload.Title,
		Content: payload.Content,
		// change after auth
		Tags: payload.Tags,
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil{
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil{
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter,r *http.Request){
	idParam := chi.URLParam(r, "postID")
	id ,err := strconv.ParseInt(idParam, 10 ,64)
	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id); if err!=nil{
		switch{
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	if err := writeJSON(w, http.StatusOK, post); err != nil{
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}



}