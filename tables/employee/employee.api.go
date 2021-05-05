package employees

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GET
func GetEmployees(res http.ResponseWriter, req *http.Request) {
	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployees()
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, employees)
	}
}

//GET{ID}
func GetEmployeeByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	employeeid, _ := strconv.Atoi(id)
	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployeeByID(employeeid)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, employees)
	}
}

//POST
func CreateEmployee(res http.ResponseWriter, req *http.Request) {
	var employee Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var employeeModel EmployeeModel
		err2 := employeeModel.CreateEmployee(&employee)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, employee)
		}
	}
}

//Update
func UpdateEmployee(res http.ResponseWriter, req *http.Request) {
	var employee Employee
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var employeeModel EmployeeModel
		err2 := employeeModel.UpdateEmployee(&employee)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, employee)
		}
	}
}

//Delete
func DeleteEmployeeByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	employeeid, _ := strconv.Atoi(id)
	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployeeByID(employeeid)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		err2 := employeeModel.DeleteEmployee(employees)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, employees)
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
