package main

import (
	"net/http"
	"log"
)
func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error){
	log.Printf("internal server error :%s path :%s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (app *application) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error){
	log.Printf("bad request error :%s path :%s error: %s", r.Method, r.URL.Path, err)
	
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) NotFoundResponse(w http.ResponseWriter, r *http.Request, err error){
	log.Printf("not found error :%s path :%s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound,err.Error())
}
