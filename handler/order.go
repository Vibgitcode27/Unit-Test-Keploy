package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an Order")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Order created successfully"}
	json.NewEncoder(w).Encode(response)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List an Order")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "List of orders"}
	json.NewEncoder(w).Encode(response)
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order by Id")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Order details by Id"}
	json.NewEncoder(w).Encode(response)
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Order by Id")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Order updated successfully"}
	json.NewEncoder(w).Encode(response)
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Order by Id")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Order deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
