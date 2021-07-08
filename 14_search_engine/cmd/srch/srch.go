package main

import (
	"fmt"
	"go_course_thinknetika/14_search_engine/pkg/crawler"
	"go_course_thinknetika/14_search_engine/pkg/crawler/spider"
	"go_course_thinknetika/14_search_engine/pkg/webapp"
	"log"
	"net/http"
)

type crwDocs struct {
	docs []crawler.Document
}

func main() {
	crw := crwDocs{}
	go func() {
		crw.docs = scan([]string{"https://golang.org", "https://go.dev"}, 2)
	}()

	r := webapp.Router()
	r.HandleFunc("/docs", crw.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index", crw.indexHandler).Methods(http.MethodGet)
	webapp.ListenAndServe(":8000", r)
}

func (crw *crwDocs) docsHandler(w http.ResponseWriter, r *http.Request) {
	if len(crw.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var docsView string
	for i, d := range crw.docs {
		docsView += "<p>" + fmt.Sprint(i, ": ") + d.Title + "</p>"
	}

	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", docsView)
}

func (crw *crwDocs) indexHandler(w http.ResponseWriter, r *http.Request) {
	if len(crw.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var indexView string
	for i, d := range crw.docs {
		indexView += "<p>" + fmt.Sprint(i, ": ", d.ID) + "</p>"
	}

	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", indexView)
}

func scan(urls []string, depth int) (docs []crawler.Document) {
	scn := spider.New()
	for _, url := range urls {
		res, err := scn.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, res...)
	}
	return docs
}
