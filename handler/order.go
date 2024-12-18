package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create an order"))
	fmt.Println("Create an order")
}
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display list of created orders"))
	fmt.Println("Display list of created orders")
}
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("display a specific order by its ID")
}
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete an order")
}
