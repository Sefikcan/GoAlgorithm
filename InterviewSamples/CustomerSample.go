package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// It represents basic customer information
type Customer struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	OrderHistory []string `json:"order_history"`
}

var db *sql.DB

func main() {

	// database conn setup
	connStr := "user=username dbname=mydb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// router setup using gorilla mux
	router := mux.NewRouter()
	router.HandleFunc("/customers/{id}", GetCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")
	router.HandleFunc("/customers", GetAllCustomers).Methods("GET")

	// server start
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetCustomer retrieves customer informations
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	// Get customer id from request parameters
	vars := mux.Vars(r)
	customerId := vars["id"]

	rows, err := db.Query("select name, email from customers where id = $1", customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customer Customer

	// populate customers
	for rows.Next() {
		err := rows.Scan(&customer.Name, &customer.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// simulate fetching order history
	customer.OrderHistory = []string{"Order1", "Order2", "Order3"}

	// Convert customer struct to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// Creeate a new customer in the database
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert customer to database
	_, err = db.Exec("insert into customers(name, email) values ($1, $2)", customer.Name, customer.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update a existing customer
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update customer to database
	_, err = db.Exec("update customers set name=$1, email=$2 where id=$3", customer.Name, customer.Email, customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Retrieves all customers
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, name, email from customers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		customer.OrderHistory = []string{"Order1", "Order2", "Order3"}
		customers = append(customers, customer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
