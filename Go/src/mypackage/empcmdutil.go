package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	//Log --
	Log     *log.Logger
	logpath = flag.String("logpath", "D:/Projects/Go/src/mypackage/employee.log", "Log Path")
	//DB - Global database variable DB
	DB *sql.DB
)

// NewLog --
func NewLog(logpath string) {
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

var allEmployee []Employee

//GetAllEmployee gives list of all Employees
func GetAllEmployee(c *gin.Context) {

	var (
		employee Employee
		result   gin.H
	)
	st, err := DB.Prepare("select * from employee")
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}
	// ther rows produced by the Query, with exeption handler
	rows, err := st.Query()
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}

	//here - we print out the lines from the slice rows
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
		c.JSON(http.StatusOK, result)
	}
	Log.Println("Total Employee Records Found: ", i)
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": allEmployee,
		"count":  len(allEmployee),
	})
	if err := rows.Err(); err != nil {
		Log.Panicln(err)
		panic(err)
	}
	Log.Println("Total Employee Records Found: ", i)
	//fmt.Print("Total: ", i)
}

//GetEmployee give Eployee details by emp id
func GetEmployee(c *gin.Context) {

	var (
		employee Employee
		result   gin.H
	)
	username := c.Param("username")
	row := DB.QueryRow("select * from employee where username = ?;", username)
	err := row.Scan(&employee.ID, &employee.Username, &employee.Password, &employee.Firstname, &employee.Lastname, &employee.Joiningdate, &employee.Skills)
	if err != nil {
		Log.Panicln(err)
		panic(err)
	}
	// ther rows produced by the Query, with exeption handler
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": employee,
			"count":  1,
		}
	}
	Log.Println("Employee details: ", username)
	c.JSON(http.StatusOK, result)
}

//AddEmployee will add new Eployee
func AddEmployee(c *gin.Context) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("Enter Password: ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Print("Enter Firstname: ")
	scanner.Scan()
	firstname := scanner.Text()
	fmt.Print("Enter Lastname: ")
	scanner.Scan()
	lastname := scanner.Text()
	fmt.Print("Enter Skills: ")
	scanner.Scan()
	skills := scanner.Text()

	var buffer bytes.Buffer

	currentTime := time.Now().Local()
	newdate := currentTime.Format("2006-01-02")

	layout := "2006-01-02"
	t, _ := time.Parse(layout, newdate)

	stmt, err := DB.Prepare("INSERT INTO employee (Username, Password, Firstname, Lastname, Joiningdate, Skills) VALUES (?,?,?,?,?,?);")
	fmt.Print("stmt : ", stmt)
	fmt.Print(skills)
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(username, password, firstname, lastname, t, skills)
	Log.Println("SQL Query Executed for add employee details")
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(username)
	buffer.WriteString(" ")
	buffer.WriteString(password)
	buffer.WriteString(firstname)
	buffer.WriteString(" ")
	buffer.WriteString(lastname)
	buffer.WriteString(" ")
	buffer.WriteString(t.String())
	buffer.WriteString(" ")
	buffer.WriteString(skills)
	buffer.WriteString(" ")
	defer stmt.Close()
	name := buffer.String()
	Log.Println("Employee Added Successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name),
	})

}

//UpdateUser will update existing Eployee details by emp id
func UpdateUser(c *gin.Context) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("Enter Password: ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Print("Enter Firstname: ")
	scanner.Scan()
	firstname := scanner.Text()
	fmt.Print("Enter Lastname: ")
	scanner.Scan()
	lastname := scanner.Text()
	currentTime := time.Now().Local()
	newdate := currentTime.Format("2006-01-02")
	layout := "2006-01-02"
	t, _ := time.Parse(layout, newdate)
	fmt.Print("Enter Skills: ")
	scanner.Scan()
	skills := scanner.Text()
	id := c.Param("id")
	Log.Println("Employee id for update..:" + id)
	var buffer bytes.Buffer
	stmt, err := DB.Prepare("update employee set Username=?, Password=?, Firstname=?, Lastname=?, Skills=? where id= ?;")
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(username, password, firstname, lastname, t, skills, id)
	Log.Println("SQL Query Executed for update employee details")
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(username)
	buffer.WriteString(" ")
	buffer.WriteString(password)
	buffer.WriteString(" ")
	buffer.WriteString(firstname)
	buffer.WriteString(" ")
	buffer.WriteString(lastname)
	buffer.WriteString(" ")
	buffer.WriteString(skills)
	buffer.WriteString(" ")
	defer stmt.Close()
	name := buffer.String()
	Log.Println("Employee details updated Successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated to %s", name),
	})

}

//DeleteEmployee will remove existing Eployee details by emp id
func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")
	Log.Println("Employee id for delete..:" + id)
	stmt, err := DB.Prepare("delete from employee where id= ?;")
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	Log.Println("SQL Query Executed for delete")
	if err != nil {
		Log.Panicln(err)
		fmt.Print(err.Error())
	}
	Log.Println("Employee Deleted Successfully..")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})

}

func main() {
	flag.Parse()
	NewLog(*logpath)
	time.Sleep(3 * time.Second)
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
	fmt.Print("Enter Database Name: ")
	scanner.Scan()
	dbName := scanner.Text()
	fmt.Print("Enter API Server Port: ")
	scanner.Scan()
	apiPort := scanner.Text()
	Log.Println("API Server Port No: ", apiPort)

	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:"+portNo+")/"+dbName+"?charset=utf8")
	//db, err := sql.Open("mysql", "root:xoriant123@tcp(localhost:3306)/golangdb?charset=utf8")
	if err != nil {
		Log.Panicln(err)
		fmt.Printf("Database connection error Exiting...")
		Log.Println("Database connection error...")
		panic(err)
	}
	DB = db
	router := gin.Default()
	emp := router.Group("api/")
	{
		Log.Println("List of all employees function invoked..")
		emp.GET("employee", GetAllEmployee)
		Log.Println("List of all employees function completed..")
		Log.Println("Employee details function by id invoked..")
		emp.GET("employee/:username", GetEmployee)
		Log.Println("Employee details function by id completed..")
		Log.Println("Add Employee details function invoked..")
		emp.POST("employee", AddEmployee)
		Log.Println("Add Employee details function completed..")
		Log.Println("Update Employee details function invoked..")
		emp.PUT("/employee/:id", UpdateUser)
		Log.Println("Update Employee details function completed..")
		Log.Println("Delete Employee details function invoked..")
		emp.DELETE("employee/:id", DeleteEmployee)
		Log.Println("Delete Employee details function completed..")
	}
	router.Run(":" + apiPort)
}
