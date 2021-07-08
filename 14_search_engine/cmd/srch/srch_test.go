package main

import (
	"fmt"
	"go_course_thinknetika/14_search_engine/pkg/crawler"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var crw crwDocs
var r *mux.Router

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	r.HandleFunc("/docs", crw.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index", crw.indexHandler).Methods(http.MethodGet)
	os.Exit(m.Run())
}

func Test_crwDocs_docsHandler(t *testing.T) {
	crw = crwDocs{docs: []crawler.Document{}}
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	var docsView string
	crw.docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://golang.org/pkg",
			Title: "src - The Go Programming Language",
		},
	}
	req = httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for i, d := range crw.docs {
		docsView += "<p>" + fmt.Sprint(i, ": ") + d.Title + "</p>"
	}
	want := "<html><body><div>" + docsView + "</div></body></html>"
	if got != want {
		t.Errorf("incorrect body: get %v, want %v", got, want)
	}
}

func Test_crwDocs_indexHandler(t *testing.T) {
	crw = crwDocs{docs: []crawler.Document{}}
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	var indexView string
	crw.docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://golang.org/pkg",
			Title: "src - The Go Programming Language",
		},
	}
	req = httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	for i, d := range crw.docs {
		indexView += "<p>" + fmt.Sprint(i, ": ", d.ID) + "</p>"
	}
	want := "<html><body><div>" + indexView + "</div></body></html>"
	if got != want {
		t.Errorf("incorrect body: get %v, want %v", got, want)
	}
}

func Test_invalidHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNotFound)
	}
}
