package customers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//GET
func GetCutomers(res http.ResponseWriter, req *http.Request) {
	var customerModel CustomerModel
	customers, err := customerModel.GetCutomers()
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, customers)
	}
}

//GET{ID}
func GetCustomerByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var customerModel CustomerModel
	customers, err := customerModel.GetCustomerByID(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, customers)
	}
}

//POST
func CreateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var customerModel CustomerModel
		err2 := customerModel.CreateCustomer(&customer)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, customer)
		}
	}
}

//Update
func UpdateCustomer(res http.ResponseWriter, req *http.Request) {
	var customer Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var customerModel CustomerModel
		err2 := customerModel.UpdateCustomer(&customer)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, customer)
		}
	}
}

//Delete
func DeleteCustomerByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var customerModel CustomerModel
	customers, err := customerModel.GetCustomerByID(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		err2 := customerModel.DeleteCustomer(customers)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, customers)
		}
	}
}

//RespondWith
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applicaction/json")
	w.WriteHeader(code)
	w.Write(response)
}
