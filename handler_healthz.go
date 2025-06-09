package main

import (
	"net/http"
)

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	resWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	resWithErr(w, http.StatusBadRequest, "Something went wrong")
}