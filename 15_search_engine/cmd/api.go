package main

import (
	"encoding/json"
	"fmt"
	"go_course_thinknetika/15_search_engine/pkg/crawler"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/new", api.newDoc).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.doc).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.editDoc).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.deleteDoc).Methods(http.MethodDelete, http.MethodOptions)
}

// curl localhost:8000/api/v1/docs
func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

// curl -XPOST localhost:8000/api/v1/docs/new -H 'application/json' -d \
//  '{"id":2,"url":"https://example.com","title":"Example"}'
func (api *API) newDoc(w http.ResponseWriter, r *http.Request) {
	d := crawler.Document{}
	json.NewDecoder(r.Body).Decode(&d)
	docs = append(docs, d)
	response, _ := json.Marshal(&d)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// curl localhost:8000/api/v1/docs/0
func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, doc := findDoc(id)
	viewDoc := string(doc.URL + " - " + doc.Title + ": " + doc.Body)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(viewDoc)
}

// curl -XPUT localhost:8000/api/v1/docs/0 -H 'application/json' -d \
//  '{"id":0,"url":"https://google.com","title":"Search","body":"information"}'
func (api *API) editDoc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	i, _ := findDoc(id)
	d := crawler.Document{}
	json.NewDecoder(r.Body).Decode(&d)
	docs[i] = d
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&d)
}

// curl -XDELETE localhost:8000/api/v1/docs/0
func (api *API) deleteDoc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	i, _ := findDoc(id)
	docs = append(docs[:i], docs[i+1:]...)
	w.WriteHeader(http.StatusOK)
}

func findDoc(id string) (int, crawler.Document) {
	for i, d := range docs {
		if fmt.Sprint(d.ID) == id {
			return i, d
		}
	}
	return -1, crawler.Document{}
}
