package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Ethea2/popio-server/models"
)

type Bills struct {
	ID              uuid.UUID `gorm:"primary_key" json:"id"`
	Mode            string    `                   json:"mode_of_payment"`
	Name            string    `                   json:"name"`
	PeriodicPayment time.Time `                   json:"periodic_payment_date"`
	Status          string    `                   json:"status"`
	GroupID         uuid.UUID `                   json:"group_id,omitempty"`
	UserID          uuid.UUID `                   json:"user_id,omitempty"`
}

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
	w.Header().Set("Content-Type", "application/json")
	type BillInput struct {
		Name            string
		Mode            string
		PeriodicPayment time.Time
		UserID          uuid.UUID
	}
	// var billInput BillInput
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(b)
	// todo: find a way to parse dates
}

func (b *Bills) DeleteBillForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete bill for user")
}

func (b *Bills) DeleteBillForGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete bill for group")
}
