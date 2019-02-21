package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Not implemented!"))
}

func getSession(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Not implemented!"))
}
