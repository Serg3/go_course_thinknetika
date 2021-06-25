package webapp

import (
	"fmt"
	"go_course_thinknetika/14_search_engine/pkg/crawler"
	"go_course_thinknetika/14_search_engine/pkg/crawler/spider"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// cache for using scan() only once
// the program starts
var crwDocs []crawler.Document

type SearchParams struct {
	urls  []string
	depth int
}

func PerformScan(urls []string, depth int) {
	sp := SearchParams{}
	sp.urls = urls
	sp.depth = depth
	go docs(sp)
}

// ListenAndServe listens all TCP network and address ':8080',
// calls Serve to handle requests on incoming connections.
func ListenAndServe() {
	mux := mux.NewRouter()
	mux.HandleFunc("/{source}", diHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", mux)
}

// HTTP handler of /docs and /index routes
// returns to the client a content with []crawler.Document.
func diHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["source"] != "docs" && vars["source"] != "index" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", crwDocs)
}

func docs(sp SearchParams) {
	if len(crwDocs) == 0 {
		crwDocs = scan(sp)
	}
}

// Function scan() uses package 'crawler'
// to search through Go sites by word
// and returs []crawler.Document result
func scan(sp SearchParams) (docs []crawler.Document) {
	scn := spider.New()
	for _, url := range sp.urls {
		res, err := scn.Scan(url, sp.depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, res...)
	}
	return docs
}
