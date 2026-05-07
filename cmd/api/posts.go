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
		app.BadRequestResponse(w, r,err)
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
		app.InternalServerError(w, r, err)
		return 
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil{
		app.InternalServerError(w,r,err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter,r *http.Request){
	idParam := chi.URLParam(r, "postID")
	id ,err := strconv.ParseInt(idParam, 10 ,64);
	if err != nil{
		app.InternalServerError(w,r ,err )
	}
	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id); if err!=nil{
		switch{
		case errors.Is(err, store.ErrNotFound):
			app.NotFoundResponse(w,r,err)
		default:
			app.InternalServerError(w, r,err)
		}
		return
	}
	if err := writeJSON(w, http.StatusOK, post); err != nil{
		app.InternalServerError(w,r,err)
		return
	}



}