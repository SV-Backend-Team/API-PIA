package employees

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(res http.ResponseWriter, req *http.Request) {
	var employeeModel EmployeeModel
	employees, err := employeeModel.FindAll()
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, employees)
	}
}

func FindByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	employeeid, _ := strconv.Atoi(id)
	var employeeModel EmployeeModel
	employees, err := employeeModel.FindByID(employeeid)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, employees)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applicaction/json")
	w.WriteHeader(code)
	w.Write(response)
}
