package main

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPI_docs(t *testing.T) {
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
			api.docs(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_newDoc(t *testing.T) {
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
			api.newDoc(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_doc(t *testing.T) {
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
			api.doc(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_editDoc(t *testing.T) {
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
			api.editDoc(tt.args.w, tt.args.r)
		})
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
