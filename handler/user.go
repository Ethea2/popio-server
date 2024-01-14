package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	db "github.com/Ethea2/popio-server/database"
	"github.com/Ethea2/popio-server/models"
)

type User struct {
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
	w.Header().Set("Content-Type", "application/json")

	type InputType struct {
		Username string
		Password string
	}
	var userinput InputType
	err := json.NewDecoder(r.Body).Decode(&userinput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userinput.Password), 12)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New()

	u.Password = string(hashedPassword)
	u.ID = id
	u.Username = userinput.Username
	db.Database.Create(&u)
	res := models.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       u,
	}
	jsonResponse, err := json.Marshal(res)

	w.Write(jsonResponse)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	db.Database.First(&u, "id = ?", userId)
	json.NewEncoder(w).Encode(u)
}
