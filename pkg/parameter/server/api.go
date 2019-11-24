package server

import (
	"encoding/json"
	"net/http"

	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter"
	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository parameter.Repository
}

// Server interface
type Server interface {
	Router() http.Handler
}

// New API server
func New(repo parameter.Repository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/health", a.health).Methods(http.MethodGet)
	s := r.PathPrefix("/parameters").Subrouter()
	s.HandleFunc("", a.fetchParameters).Methods(http.MethodGet)
	s.HandleFunc("/{ID:[a-zA-Z0-9_]+}", a.updateParameter).Methods(http.MethodPut)
	s.HandleFunc("/{ID:[a-zA-Z0-9_]+}", a.deleteParameter).Methods(http.MethodDelete)
	s.HandleFunc("/{ID:[a-zA-Z0-9_]+}", a.fetchParameter).Methods(http.MethodGet)

	//r.Use(loggingMiddleware)

	amw := authenticationMiddleware{}
	amw.Populate()

	s.Use(amw.Middleware)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (a *api) updateParameter(w http.ResponseWriter, r *http.Request) {
	var param parameter.Parameter
	_ = json.NewDecoder(r.Body).Decode(&param)

	err := a.repository.UpdateParameter(&param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (a *api) deleteParameter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := a.repository.DeleteParameter(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (a *api) fetchParameters(w http.ResponseWriter, r *http.Request) {
	parameters, _ := a.repository.FetchParameters()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parameters)
}

func (a *api) fetchParameter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param, err := a.repository.FetchParameterByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Parameter Not found")
		return
	}

	json.NewEncoder(w).Encode(param)
}
