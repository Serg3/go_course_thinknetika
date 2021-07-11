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

func (api *API) endpoints() {
	api.router.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs", api.newDoc).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.doc).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.editDoc).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/api/v1/docs/{id}", api.deleteDoc).Methods(http.MethodDelete, http.MethodOptions)
}

// curl localhost:8000/api/v1/docs
func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(docs); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// curl -XPOST localhost:8000/api/v1/docs/new -H 'application/json' -d \
//  '{"id":2,"url":"https://example.com","title":"Example"}'
func (api *API) newDoc(w http.ResponseWriter, r *http.Request) {
	d := crawler.Document{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}

	docs = append(docs, d)
	response, err := json.Marshal(&d)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// curl localhost:8000/api/v1/docs/0
func (api *API) doc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index, doc := findDoc(id)
	if index < 0 {
		http.Error(w, "Doc not found", http.StatusNotFound)
		return
	}

	viewDoc := string(doc.URL + " - " + doc.Title + ": " + doc.Body)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(viewDoc); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// curl -XPUT localhost:8000/api/v1/docs/0 -H 'application/json' -d \
//  '{"id":0,"url":"https://google.com","title":"Search","body":"information"}'
func (api *API) editDoc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index, _ := findDoc(id)
	if index < 0 {
		http.Error(w, "Doc not found", http.StatusNotFound)
		return
	}

	d := crawler.Document{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}

	docs[index] = d
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&d); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// curl -XDELETE localhost:8000/api/v1/docs/0
func (api *API) deleteDoc(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index, _ := findDoc(id)
	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	docs = append(docs[:index], docs[index+1:]...)
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
