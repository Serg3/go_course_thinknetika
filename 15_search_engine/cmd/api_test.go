package main

import (
	"bytes"
	"encoding/json"
	"go_course_thinknetika/15_search_engine/pkg/crawler"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var api *API

func TestMain(m *testing.M) {
	api = new(API)
	api.router = mux.NewRouter()
	api.Endpoints()
	os.Exit(m.Run())
}

func TestAPI_docs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong code: got %d, want %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_newDoc(t *testing.T) {
	data := crawler.Document{
		ID:    0,
		URL:   "https://google.com",
		Title: "Search",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs/new", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("error: got %d, want %d", rr.Code, http.StatusCreated)
	}

	got := docs[0]
	want := data
	if got != want {
		t.Errorf("error: got %v, want %v", got, want)
	}
}

func TestAPI_doc(t *testing.T) {
	data := crawler.Document{
		ID:    0,
		URL:   "https://google.com",
		Title: "Search",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs/0", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong code: got %d, want %d", rr.Code, http.StatusOK)
	}

	got := docs[0]
	want := data
	if got != want {
		t.Errorf("error: got %v, want %v", got, want)
	}
}

func TestAPI_editDoc(t *testing.T) {
	data := crawler.Document{
		ID:    1,
		URL:   "https://google.com",
		Title: "Search",
		Body:  "add a body",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/0/edit", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong code: got %d, want %d", rr.Code, http.StatusOK)
	}

	got := docs[0]
	want := data
	if got != want {
		t.Errorf("error: got %v, want %v", got, want)
	}
}

func TestAPI_deleteDoc(t *testing.T) {
	type fields struct {
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				router: tt.fields.router,
			}
			api.deleteDoc(tt.args.w, tt.args.r)
		})
	}
}
