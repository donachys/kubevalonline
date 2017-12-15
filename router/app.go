package router

import "github.com/gorilla/mux"

// App adds app routes and returns the router.
func App() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Path("/").Methods("GET").Name(Index)
	return r
}
