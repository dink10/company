// Package handler is responsible for API to the whole service.
// It's api layer.
//
// Doesn't matter which type of api is created: http, websocket, mq.
// All of them have to be on this layer.
//
// The main idea of the layer is a door to our application.
// On this layer we get reuqests from users.
// On this layer we prepare data (convert it from transport format to go structures),
// prepare for business logic.
// This layer MUST communicate with buisiness logic layer (controller) only.
// All needed interfaces for buisiness logic layer MUST be created in this layer.
package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tty2/company/model"
)

type (
	// Service is a controller interface.
	Service interface {
		CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
		GetEmployee(ctx context.Context, id int) (model.Employee, error)
		RaiseSalary(ctx context.Context, id int, amount int) error
		DeleteEmployee(ctx context.Context, id int) error
	}
	// API is a set of api geteways.
	API struct {
		Mux     *mux.Router
		service Service
	}
)

// New sets new api and returns it.
func New(serv Service) API {
	api := API{
		Mux:     mux.NewRouter(),
		service: serv,
	}

	api.Mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		toJSON(w, http.StatusOK, struct {
			Status string `json:"status"`
		}{Status: "ok"})
	})
	api.Mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		toJSON(w, http.StatusOK, struct {
			Status string `json:"status"`
		}{Status: "ok"})
	})

	es := mux.NewRouter().PathPrefix("/v1/employee").Subrouter()
	es.Path("").
		Methods(http.MethodPost).
		HandlerFunc(api.Create)
	es.Path("/{id}").
		Methods(http.MethodGet).
		HandlerFunc(api.GetEmployee)
	es.Path("/raise").
		Methods(http.MethodPost).
		HandlerFunc(api.RaiseSalary)
	es.Path("/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(api.DeleteEmployee)
	api.Mux.PathPrefix("/v1/employee").Handler(es)

	return api
}

// Create method.
func (a API) Create(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Printf("error on close request body: %v", err)
		}
	}()
	err = emp.Validate()
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	newEmp, err := a.service.CreateEmployee(r.Context(), emp)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	toJSON(w, http.StatusOK, newEmp)
}

// GetEmployee method.
func (a API) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		toJSONError(w, 400, errors.New("parameter id wasn't passed"))
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		toJSONError(w, 400, fmt.Errorf("couldn't parse id not int: %v", err))
		return
	}
	emp, err := a.service.GetEmployee(r.Context(), id)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	toJSON(w, http.StatusOK, emp)
}

// RaiseSalary method.
func (a API) RaiseSalary(w http.ResponseWriter, r *http.Request) {
	raise := struct {
		Amount int `json:"amount"`
		ID     int `json:"id"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&raise)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Printf("error on close request body: %v", err)
		}
	}()
	err = a.service.RaiseSalary(r.Context(), raise.ID, raise.Amount)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	toJSON(w, http.StatusNoContent, nil)
}

// DeleteEmployee method.
func (a API) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		toJSONError(w, 400, errors.New("parameter id wasn't passed"))
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		toJSONError(w, 400, fmt.Errorf("couldn't parse id not int: %v", err))
		return
	}
	err = a.service.DeleteEmployee(r.Context(), id)
	if err != nil {
		toJSONError(w, 400, err)
		return
	}
	toJSON(w, http.StatusNoContent, nil)
}
