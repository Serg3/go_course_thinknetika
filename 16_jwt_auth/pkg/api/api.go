package api

import (
	"context"
	"encoding/json"
	"fmt"
	jwtauth "go_course_thinknetika/16_jwt_auth/pkg/auth"
	users "go_course_thinknetika/16_jwt_auth/pkg/db"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// API's information structure.
type API struct {
	router *mux.Router
	users  *users.Users
}

// LogInfo is a struct for users' credentials.
type LogInfo struct {
	Usr string
	Psw string
}

// New creates API object.
func New() *API {
	api := API{}
	api.users = users.New()
	api.router = mux.NewRouter()
	api.endpoints()
	api.router.Use(api.authorized)
	api.router.Use(api.requestID)
	api.router.Use(api.logger)
	return &api
}

// Router returns API's router.
func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) endpoints() {
	api.router.HandleFunc("/", api.mainHandler).Methods(http.MethodGet)
	api.router.HandleFunc("/auth", api.authentication).Methods(http.MethodPost)
}

func (a *API) mainHandler(w http.ResponseWriter, r *http.Request) {
	if admin, ok := r.Context().Value("admin").(bool); !ok || !admin {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}
}

func (api *API) authentication(w http.ResponseWriter, r *http.Request) {
	var log LogInfo
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr := api.users.User(log.Usr, log.Psw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	token, err := jwtauth.NewToken(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(token)
}

func (api *API) authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth" {
			next.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
			return
		}

		claims, err := jwtauth.VerifyToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "admin", claims.Admin)
		ctx = context.WithValue(ctx, "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestID", rand.Intn(1_000_000))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID int
		var requestID int

		if val, ok := r.Context().Value("userID").(int); ok {
			userID = val
		}
		if val, ok := r.Context().Value("requestID").(int); ok {
			requestID = val
		}

		fmt.Printf("Method: %v, Addr: %v, URI: %v, userID: %v, requestID: %v\r\n", r.Method, r.RemoteAddr, r.RequestURI, userID, requestID)
		next.ServeHTTP(w, r)
	})
}
