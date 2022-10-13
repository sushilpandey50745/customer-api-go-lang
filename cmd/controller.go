package main

import (
	//internal
	"customerapp/domain"
	"encoding/json"
	"fmt"
	"net/http"

	//external
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CustomerHandler struct {
	repo   domain.Repository
	Logger *zap.Logger
}

func (ctl CustomerHandler) Post(w http.ResponseWriter, r *http.Request) {
	defer ctl.Logger.Sync()
	var customer domain.Customer
	fmt.Println("Request here")
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create Customer
	if err := ctl.repo.Create(customer); err != nil {
		ctl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == domain.ErrCustomerExists {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctl.Logger.Info("created Customer",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

func (ctrl CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	defer ctrl.Logger.Sync()
	vars := mux.Vars(r)
	fmt.Println("vars", vars)
	id := vars["custid"]

	if customer, err := ctrl.repo.GetById(id); err != nil {
		ctrl.Logger.Error(err.Error(),
			zap.String("cust id", id),
			zap.String("url", r.URL.String()),
		)
		if err == domain.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(customer)
		if err != nil {
			ctrl.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	}
}
func (ctrl CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	defer ctrl.Logger.Sync()
	vars := mux.Vars(r)
	custId := vars["custid"]

	if err := ctrl.repo.Delete(custId); err != nil {
		ctrl.Logger.Error(err.Error(),
			zap.String("customer id", custId),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.Logger.Info("deleted customer",
		zap.String("customer id", custId),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
func (ctrl CustomerHandler) Put(w http.ResponseWriter, r *http.Request) {
	defer ctrl.Logger.Sync()
	vars := mux.Vars(r)
	custid := vars["custid"]
	var customer domain.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		ctrl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	// Update
	if err := ctrl.repo.Update(custid, customer); err != nil {
		ctrl.Logger.Error(err.Error(),
			zap.String("Customer id", custid),
			zap.String("URL", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ctrl.Logger.Info("Updated customer",
		zap.String("CustomerID", custid),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
func (ctrl CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if customers, err := ctrl.repo.GetAll(); err != nil {
		ctrl.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		if err == domain.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			ctrl.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
