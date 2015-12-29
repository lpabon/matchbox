package api

import (
	"net/http"
	"path/filepath"
)

// Spec is a named group of configs.
type Spec struct {
	// spec identifier
	ID string `json:"id"`
	// boot kernel, initrd, and kernel options
	BootConfig *BootConfig `json:"boot"`
}

type specResource struct {
	store Store
}

func newSpecResource(mux *http.ServeMux, pattern string, store Store) {
	gr := &specResource{
		store: store,
	}
	mux.Handle(pattern, logRequests(requireGET(gr)))
}

func (r *specResource) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := filepath.Base(req.URL.Path)
	config, err := r.store.Spec(id)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	renderJSON(w, config)
}
