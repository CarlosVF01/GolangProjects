package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)
var db *sql.DB

type Employee struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	SetupDatabaseConnection()
	handleRequests()
}
// Function that handles the endpoints that are accessible
func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/employees/{name:[a-zA-Z]+}", GetEmployeeByNameHandler).Methods("GET")
	router.HandleFunc("/employees/{id:[0-9]+}", GetEmployeeByIdHandler).Methods("GET")
	router.HandleFunc("/employees", GetAllEmployeesHandler).Methods("GET")
	router.HandleFunc("/employees/add", PostNewEmployeeHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
func GetEmployeeByIdHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	fmt.Println("Endpoint Hit: GetEmployeeById")
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	employee := GetEmployeesByIdQuery(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}
func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllEmployees")

	employees := GetAllEmployeesQuery()
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(employees)
}
func GetEmployeeByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Endpoint Hit: GetEmployeeByName")

	employees := GetEmployeesByNameQuery(vars["name"])
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(employees)


}
func PostNewEmployeeHandler(w http.ResponseWriter, r *http.Request){
	var employee Employee
	fmt.Println("Endpoint Hit: PostNewEmployee")

	dec := json.NewDecoder(r.Body)
	dec.Decode(&employee)

	employee = PostNewEmployeeQuery(employee)

	json.NewEncoder(w).Encode(employee)

}
func PostNewEmployeeQuery(employee Employee) Employee{
	rows, err := db.Query("INSERT INTO emp (name, age) VALUES(?, ?)", employee.Name, employee.Age)

	if err != nil{
		fmt.Println(err)
	}

	defer rows.Close()


	for rows.Next() {
		var employeeOnDb Employee
		rows.Scan(&employeeOnDb.ID, &employeeOnDb.Name, &employeeOnDb.Age)
		employee = employeeOnDb
	}

	return employee
}
func GetEmployeesByNameQuery(name string) []Employee {
	// An employees slice to hold data from returned rows.
	var employees []Employee

	rows, _ := db.Query("SELECT * FROM emp WHERE name = ?", name)

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var employee Employee

		rows.Scan(&employee.ID, &employee.Name, &employee.Age)
		employees = append(employees, employee)
	}

	return employees
}
func GetAllEmployeesQuery() []Employee {
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
func GetEmployeesByIdQuery(id int64) Employee {
	var employee Employee

	rows, _ := db.Query("SELECT * FROM emp WHERE id = ?", id)

	defer rows.Close()

	for rows.Next() {
		var dbEmployee Employee
		rows.Scan(&dbEmployee.ID, &dbEmployee.Name, &dbEmployee.Age)

		employee = dbEmployee
	}

	return employee
}
// Function that sets up the mysql database connection
func SetupDatabaseConnection() {
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
