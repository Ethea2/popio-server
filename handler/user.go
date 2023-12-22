package handler

import (
	"fmt"
	"net/http"
)

type User struct{}

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
	fmt.Println("Get user")
}
