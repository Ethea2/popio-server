package handler

import (
	"fmt"
	"net/http"
)

type Bills struct{}

func (b *Bills) GetBillsByUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all bills by UserID")
}

func (b *Bills) GetBillsByGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all bills by UserID")
}

func (b *Bills) AddBillForGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Bill for Group")
}

func (b *Bills) AddBillForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Bill for Group")
}

func (b *Bills) DeleteBillForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete bill for user")
}

func (b *Bills) DeleteBillForGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete bill for group")
}
