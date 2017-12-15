package router

import "github.com/gorilla/mux"

// API adds api routes and returns the router.
func API() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Path("/val").Methods("POST").Name(Validate)
	return r
}
