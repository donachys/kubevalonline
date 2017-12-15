package app

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/donachys/kubevalonline/router"
)

var (
	// StaticDir is the directory containing static assets.
	StaticDir = filepath.Join(defaultBase("github.com/donachys/kubevalonline/app"), "static")
)

// Handler sets up routes.
func Handler() *mux.Router {
	r := router.App()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))
	r.Get(router.Index).HandlerFunc(index)
	return r
}

func index(rw http.ResponseWriter, req *http.Request) {
	renderTemplate(rw, "index.html")
}
