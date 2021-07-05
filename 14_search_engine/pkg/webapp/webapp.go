package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	return mux.NewRouter()
}

// ListenAndServe listens all TCP network and passed address,
// calls Serve to handle requests on incoming connections.
func ListenAndServe(address string, r http.Handler) {
	http.ListenAndServe(address, r)
}
