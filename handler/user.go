package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"

	db "github.com/Ethea2/popio-server/database"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `                   json:"username"`
	Password string    `                   json:"password"`
}

type UserGroup struct {
	gorm.Model
	ID    uuid.UUID
	User  User
	Group Group
}

func (u *User) SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUp")
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Database.First(&u)

	json.NewEncoder(w).Encode(u)
}
