package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	//Log --
	Log     *log.Logger
	logpath = flag.String("logpath", "D:/Projects/Go/src/mypackage/empReport.log", "Log Path")
	//DB - Global database variable DB
	DB *sql.DB
	allEmployee []Employee
)

// NewLog --
func crateLogs(logpath string) {
	println("LogFile: " + logpath)
	file, err := os.Create(logpath)
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

//Employee data struct
type Employee struct {
	ID          string //`json:"id, omitempty"`
	Username    string //`json:"username, omitempty"`
	Password    string //`json:"password, omitempty"`
	Firstname   string //`json:"first_name, omitempty"`
	Lastname    string //`json:"last_name, omitempty"`
	Joiningdate string //`json:"joining_date, omitempty"`
	Skills      string //`json:"skills, omitempty"`
}

func main() {
	flag.Parse()
	crateLogs(*logpath)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter DB Username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("Enter DB Password: ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Print("Enter DB PortNo: ")
	scanner.Scan()
	portNo := scanner.Text()
	Log.Println("DB PortNo: ", portNo)
	fmt.Print("Enter Database Name: ")
	scanner.Scan()
	dbName := scanner.Text()
	Log.Println("Database Name: ", dbName)

	DB, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:"+portNo+")/"+dbName+"?charset=utf8")
	//DB, err := sql.Open("mysql", "root:xoriant123@tcp(localhost:3306)/golangdb?charset=utf8")
	if err != nil {
		fmt.Printf("Database connection error Exiting...")
		Log.Panicln(err)
		panic(err)
	}

	fmt.Print("Enter Employee Skill: ")
	scanner.Scan()
	skill := scanner.Text()
	Log.Println("Enter Employee Skill: ", skill)

	generateReport(skill, DB)
}

func generateReport(skill string, DB *sql.DB) {

	var employee Employee
	st, err := DB.Prepare("select * from employee where skills = ?;")
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}

	rows, err := st.Query(skill)
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}

	i := 0
	for rows.Next() {
		i++
		if err := rows.Scan(&employee.ID, &employee.Username, &employee.Password, &employee.Firstname, &employee.Lastname, &employee.Joiningdate, &employee.Skills); err != nil {
			Log.Panicln(err)
			panic(err)
		}
		allEmployee = append(allEmployee, employee)
		if err != nil {
			Log.Panicln(err)
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	result, err := json.MarshalIndent(allEmployee, "", "	")
	fmt.Println(string(result))
	Log.Println("Total Employee Records Found: ", i)
}
