package main

import (
	"net/http"

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