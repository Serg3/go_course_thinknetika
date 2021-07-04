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
func ListenAndServe(address string, docs *[]crawler.Document) error {
	r = mux.NewRouter()

	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		docsHandler(w, r, docs)
	}).Methods(http.MethodGet)

	r.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, docs)
	}).Methods(http.MethodGet)

	return http.ListenAndServe(address, r)
}

// HTTP handler of /docs route
// returns to the client a content of []crawler.Document.
func docsHandler(w http.ResponseWriter, r *http.Request, docs *[]crawler.Document) {
	if len(*docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var docsView string
	for i, d := range *docs {
		docsView += "<p>" + fmt.Sprint(i, ": ") + d.Title + "</p>"
	}

	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", docsView)
}

// HTTP handler of /index route
// returns to the client a content of []crawler.Document.
func indexHandler(w http.ResponseWriter, r *http.Request, docs *[]crawler.Document) {
	if len(*docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var indexView string
	for i, d := range *docs {
		indexView += "<p>" + fmt.Sprint(i, ": ", d.ID) + "</p>"
	}

	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", indexView)
}
