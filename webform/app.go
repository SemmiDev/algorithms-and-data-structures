package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId 	 uint
	CustomerName string
	SSN 		 string
}

func panicIfNeeded(err error)  {
	if err != nil {
		 panic(err.Error())
	}
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := ""
	databaseName := "crm"
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

func GetCustomerById(customerID uint) Customer {
	database := GetConnection()
	rows, err := database.Query("SELECT * FROM customer WHERE customer_id=?", customerID)
	panicIfNeeded(err)

	customer := Customer{}
	for rows.Next() {
		var customerID uint
		var customerName string
		var SSN string

		err = rows.Scan(&customerID, &customerName, &SSN)
		panicIfNeeded(err)

		customer.CustomerId = customerID
		customer.CustomerName = customerName
		customer.SSN = SSN
	}

	defer database.Close()
	return customer
}

func GetCustomers(sort string) []Customer {
	database := GetConnection()
	var rows *sql.Rows
	var err error

	if sort == "asc" {
		rows, err = database.Query("SELECT * FROM customer ORDER BY customer_id")
		panicIfNeeded(err)
	}else {
		rows, err = database.Query("SELECT * FROM customer ORDER BY customer_id DESC")
		panicIfNeeded(err)
	}

	customer := Customer{}
	var customers []Customer

	for rows.Next() {
		var customerId uint
		var customerName string
		var ssn string

		err = rows.Scan(&customerId, &customerName, &ssn)
		panicIfNeeded(err)

		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn

		customers = append(customers, customer)
	}

	defer database.Close()
	return customers
}

func InsertCustomer(customer Customer) {
	database := GetConnection()
	insert, err := database.Prepare("INSERT INTO customer(customer_name, ssn) VALUES(?,?)")
	panicIfNeeded(err)

	_, _ = insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

func UpdateCustomer(customer Customer) {
	database := GetConnection()
	update, err := database.Prepare("UPDATE customer SET customer_name=?, ssn=? WHERE customer_id=?")
	panicIfNeeded(err)

	_, _ = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

func DeleteCustomer(customer Customer) {
	database := GetConnection()
	delete, err := database.Prepare("DELETE FROM customer WHERE customer_id=?")
	panicIfNeeded(err)

	_, _ = delete.Exec(customer.CustomerId)
	defer database.Close()
}

var template_html = template.Must(template.ParseGlob("templates/*"))

func Create(writer http.ResponseWriter, request *http.Request) {
	_ = template_html.ExecuteTemplate(writer, "Create", nil)
}
func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	InsertCustomer(customer)

	var customers []Customer
	customers = GetCustomers("asc")
	_ = template_html.ExecuteTemplate(writer, "View", customers)
}
func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId uint
	var customerIdStr string

	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")

	UpdateCustomer(customer)
	var customers []Customer

	customers = GetCustomers("asc")
	_ = template_html.ExecuteTemplate(writer, "View", customers)
}func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId uint
	var customerIdStr string

	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	_ = template_html.ExecuteTemplate(writer, "Update", customer)
}
func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId uint
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)

	var customer Customer
	customer = GetCustomerById(customerId)

	DeleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers("asc")
	_ = template_html.ExecuteTemplate(writer, "View", customers)
}
func View(writer http.ResponseWriter, request *http.Request) {
	var customerId uint
	var customerIdStr string

	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)

	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	customers = append(customers, customer)
	_ = template_html.ExecuteTemplate(writer, "View", customers)
}
func Home(writer http.ResponseWriter, request *http.Request) {
	sort := request.FormValue("sort")
	var customers []Customer
	switch sort {
		case "ASC","asc":
			customers = GetCustomers("asc")
		case "DES","dsc":
			customers =  GetCustomers("dsc")
		default:
			customers = GetCustomers("asc")
	}
	log.Println(customers)
	if len(customers) != 0 {
		_ = template_html.ExecuteTemplate(writer, "View", customers)
	}
	_ = template_html.ExecuteTemplate(writer, "Empty", customers)
}
func main() {
	log.Println("Server started on: http://localhost:8000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	log.Fatal(http.ListenAndServe(":8000", nil))
}