package main

import (
	"github.com/icrowley/fake"
	"time"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var entries int = 100000

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	current_time := time.Now().Local()
	newdate := current_time.Format("2006-01-02")
	layout := "2006-01-02"
	t, _ := time.Parse(layout, newdate)
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
	if err != nil {
		fmt.Printf("Exiting...")
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO employee (Username, Password, Firstname, Lastname, Joiningdate, Skills) VALUES (?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	for i := 0; i < entries; i++ {
		username := fake.UserName()
		password := fake.SimplePassword()
		firstname := fake.FirstName()
		lastname := fake.LastName()
		skills := fake.Language()
		_, err = stmt.Exec(username, password, firstname, lastname, t, skills)

		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

/*
func BulkInsert(unsavedRows []*ExampleRowStruct) error {
    valueStrings := make([]string, 0, len(unsavedRows))
    valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
    for _, post := range unsavedRows {
        valueStrings = append(valueStrings, "(?, ?, ?)")
        valueArgs = append(valueArgs, post.Column1)
        valueArgs = append(valueArgs, post.Column2)
        valueArgs = append(valueArgs, post.Column3)
    }
    stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
    _, err := db.Exec(stmt, valueArgs...)
    return err
}


=====================================================================
type Value []interface{} // defined in the sql package

batch := []Value
for i := 0; i < N; i++ {
    batch = append(batch, Value{1, 1.3, "x"})
}
db.Exec("INSERT INTO films (code, title, did, date_prod, kind) VALUES ?", batch)

*/