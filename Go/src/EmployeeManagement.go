package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Employee data
type Employee struct {
	ID          string `json:"id, omitempty"`
	Firstname   string `json:"first_name, omitempty"`
	Lastname    string `json:"last_name, omitempty"`
	Joiningdate string `json:"joining_date, omitempty"`
	Skills      string `json:"skills, omitempty"`
}

var allEmployee []Employee

//GetAllEmployee gives list of all Employees
func GetAllEmployee(w http.ResponseWriter, req *http.Request) {
	_, err := json.Marshal(allEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allEmployee)
}

//GetEmployee give Eployee details by emp id
func GetEmployee(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, item := range allEmployee {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

//AddEmployee will add new Eployee
func AddEmployee(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var employee Employee
	_ = json.NewDecoder(req.Body).Decode(&employee)
	employee.ID = params["id"]
	allEmployee = append(allEmployee, employee)
	json.NewEncoder(w).Encode(allEmployee)
}

//DeleteEmployee will remove existing Eployee details by emp id
func DeleteEmployee(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range allEmployee {
		if item.ID == params["id"] {
			allEmployee = append(allEmployee[:index], allEmployee[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(allEmployee)
}

func employeeDetails() {
	router := mux.NewRouter()
	employeeDB := make(map[int]Employee)

	emp1 := Employee{"X001", "XF001", "XL001", "05/09/2016", "Core Java"}
	emp2 := Employee{"X002", "XF002", "XL002", "12/09/2016", "J2EE"}
	emp3 := Employee{"X003", "XF003", "XL003", "19/09/2016", "Spring"}
	emp4 := Employee{"X004", "XF004", "XL004", "26/09/2016", "GO Lang"}
	emp5 := Employee{"X005", "XF005", "XL005", "03/10/2016", "Rest API"}

	employeeDB[1] = emp1
	employeeDB[2] = emp2
	employeeDB[3] = emp3
	employeeDB[4] = emp4
	employeeDB[5] = emp5

	for _, v := range employeeDB {
		allEmployee = append(allEmployee, v)
	}

	router.HandleFunc("/allEmployee", GetAllEmployee).Methods("GET")
	router.HandleFunc("/allEmployee/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/allEmployee/{id}", AddEmployee).Methods("POST")
	router.HandleFunc("/allEmployee/{id}", DeleteEmployee).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
