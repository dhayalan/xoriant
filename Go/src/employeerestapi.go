package main

import (
	"encoding/json"
	"net/http"
    //"sort"
)

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string
   // Age int
}
/*
func (e Employee) String() string {
    return fmt.Sprintf("%s: %d", e.FirstName, e.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Employee

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Example() {
    people := []Employee{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)

    // Output:
    // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
    // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
}*/

func main() {
	http.HandleFunc("/", GetEmployees)
	http.HandleFunc("/version", GetVersion)
	http.ListenAndServe(":8080", nil)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	s := "1.0"
	w.Write([]byte(s))
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	employeeDB := make(map[int]Employee)

	emp1 := Employee{"Rob", "Pike",}
	emp2 := Employee{"Robert", "Griesemer"}
	emp3 := Employee{"Ken", "Thompson"}

	employeeDB[1] = emp1
	employeeDB[2] = emp2
	employeeDB[3] = emp3

	var employees []Employee

	for _, v := range employeeDB {

		employees = append(employees, v)

	}

	js, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
