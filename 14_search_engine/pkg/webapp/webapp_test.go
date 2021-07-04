package webapp

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

var docs []crawler.Document

func TestMain(m *testing.M) {
	docs = []crawler.Document{}
	r = mux.NewRouter()
	r.HandleFunc("/{source}", func(w http.ResponseWriter, r *http.Request) {
		docsHandler(w, r, &docs)
	}).Methods(http.MethodGet)
	os.Exit(m.Run())
}

func Test_handler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("content-type", "plain/text")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("ncorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	req = httptest.NewRequest(http.MethodGet, "/invalid", nil)
	req.Header.Add("content-type", "plain/text")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Errorf("incorrect code: get %d, want %d", rr.Code, http.StatusNoContent)
	}

	var docsView string
	docs = []crawler.Document{
		{
			ID:    0,
			URL:   "https://golang.org/pkg",
			Title: "src - The Go Programming Language",
		},
	}
	wantDocs := []crawler.Document{
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
	get := string(body)
	for i, d := range wantDocs {
		docsView += "<p>" + fmt.Sprint(i, ": ") + d.Title + "</p>"
	}
	want := "<html><body><div>" + docsView + "</div></body></html>"
	if get != want {
		t.Errorf("incorrect body: get %v, want %v", get, want)
	}
}
