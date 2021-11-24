package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)
var db *sql.DB
type Employee struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	setupDatabaseConnection()
	handleRequests()
}
// Function that handles the endpoints that are accessible
func handleRequests() {
	http.HandleFunc("/employees", returnAllEmployees)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// Function that queries for employees that have the specified name
func findEmployeesByNameQuery(name string) []Employee {
	// An employees slice to hold data from returned rows.
	var employees []Employee

	rows, _ := db.Query("SELECT * FROM emp WHERE name = ?", name)

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var employee Employee

		employees = append(employees, employee)
	}

	return employees
}
func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(getAllEmployeesQuery())
}
func getAllEmployeesQuery() []Employee {
	var employees []Employee

	rows, _ := db.Query("SELECT * FROM emp")

	defer rows.Close()

	for rows.Next() {
		var employee Employee
		rows.Scan(&employee.ID, &employee.Name, &employee.Age)

		employees = append(employees, employee)
	}

	return employees
}
// Function that sets up the mysql database connection
func setupDatabaseConnection() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "employees",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
