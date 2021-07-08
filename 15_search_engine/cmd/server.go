package main

import (
	"go_course_thinknetika/15_search_engine/pkg/crawler"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api    *API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &API{router: srv.router}
	srv.api.Endpoints()
	http.ListenAndServe(":8000", srv.router)
}

var docs = []crawler.Document{
	{
		ID:    0,
		URL:   "google.com",
		Title: "search",
	},
}
