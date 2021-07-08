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
	api.router.HandleFunc("/api/v1/docs/{id}/edit", api.editDoc).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}/delete", api.deleteDoc).Methods(http.MethodDelete, http.MethodOptions)
}

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) newDoc(w http.ResponseWriter, r *http.Request) {
	// var b crawler.Document
	// err := json.NewDecoder(r.Body).Decode(&b)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// docs = append(docs, b)
}

func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doc := findDoc(vars["id"])
	res := string(doc.URL + " - " + doc.Title + ": " + doc.Body)
	json.NewEncoder(w).Encode(res)
}

func (api *API) editDoc(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// doc := findDoc(vars["id"])
	// fmt.Println(doc)
	// editDoc := &crawler.Document{
	// 	ID:    0,
	// 	URL:   "google.com",
	// 	Title: "search",
	// 	Body:  "list of result",
	// }
	// fmt.Println(&doc)
	// json.NewDecoder(r.Body).Decode(&doc)
	// doc = editDoc
	// fmt.Println(doc)
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(&doc)
}

func (api *API) deleteDoc(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// doc := findDoc(vars["id"])
	// append(docs[:doc.ID], docs[doc.ID+1:]...)
}

func findDoc(id string) *crawler.Document {
	for _, d := range docs {
		if fmt.Sprint(d.ID) == id {
			return &d
		}
	}
	return &crawler.Document{}
}
