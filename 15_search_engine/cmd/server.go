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
	srv.api.endpoints()
	http.ListenAndServe(":8000", srv.router)
}

var docs = []crawler.Document{
	{
		ID:    0,
		URL:   "https://google.com",
		Title: "Search",
	},
	{
		ID:    1,
		URL:   "https://golang.org",
		Title: "Go",
		Body:  "programming",
	},
}
