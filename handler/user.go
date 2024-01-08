package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	db "github.com/Ethea2/popio-server/database"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `                   json:"username"`
	Password string    `                   json:"password"`
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
	db.Database.First(&u)

	fmt.Println(u.Username)
}
