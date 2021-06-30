package webapp

import (
	"fmt"
	"go_course_thinknetika/14_search_engine/pkg/crawler"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

// ListenAndServe listens all TCP network and address ':8080',
// calls Serve to handle requests on incoming connections.
func ListenAndServe(address string, docs *[]crawler.Document, paths map[string]bool) error {
	r = mux.NewRouter()
	r.HandleFunc("/{source}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, docs, paths)
	}).Methods(http.MethodGet)
	return http.ListenAndServe(address, r)
}

// HTTP handler of /docs and /index routes
// returns to the client a content of []crawler.Document.
func handler(w http.ResponseWriter, r *http.Request, docs *[]crawler.Document, paths map[string]bool) {
	vars := mux.Vars(r)
	if !paths[vars["source"]] {
		w.WriteHeader(http.StatusNotFound) // or http.StatusMethodNotAllowed
		return
	}
	if len(*docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", *docs)
}
