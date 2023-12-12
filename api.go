package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)

	case "POST":
		return s.handleCreateAccount(w, r)

	case "DELETE":
		return s.handleDeleteAccount(w, r)

	default:
		fmt.Errorf("method not allowed %s", r.Method)
	}

	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return WriteJSON(w, http.StatusOK, Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleSearchStudent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleAddStudent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteStudent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

type APIServer struct {
	listenAddr string
}

type ApiError struct {
	Error string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// TODO: Handle error
			WriteJSON(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func NewAPIServer(listendaddr string) *APIServer {
	return &APIServer{
		listenAddr: listendaddr,
	}
}
