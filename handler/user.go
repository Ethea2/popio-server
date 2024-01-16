package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	db "github.com/Ethea2/popio-server/database"
	"github.com/Ethea2/popio-server/models"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `                   json:"username"`
	Password string    `                   json:"password,omitempty"`
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
		json.NewEncoder(w).Encode(models.Response{
			StatusCode: 400,
			Message:    "Signup failed",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userinput.Password), 12)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			StatusCode: 400,
			Message:    "Signup failed",
		})
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
		Data:       u.ID,
	}
	jsonResponse, err := json.Marshal(res)

	w.Write(jsonResponse)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	w.Header().Set("Content-Type", "application/json")
	type Payload struct {
		Token    string
		UserID   uuid.UUID
		Username string
	}
	type UserInput struct {
		Username string
		Password string
	}
	var userinput UserInput

	err := json.NewDecoder(r.Body).Decode(&userinput)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}
	db.Database.First(&u, "username = ?", userinput.Username)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userinput.Password)); err != nil {
		json.NewEncoder(w).Encode(models.Response{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	// token := jwt.New(jwt.SigningMethodHS256)
	//
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = u.Username
	// claims["id"] = u.ID
	// claims["admin"] = true
	// claims["expiration"] = time.Now().Add(time.Hour * 72).Unix()
	//
	// finalToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	// if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userinput.Password)); err != nil {
	// 	json.NewEncoder(w).Encode(models.Response{
	// 		StatusCode: 400,
	// 		Message:    "Login failed",
	// 	})
	// 	return
	// }

	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)
	_, finalToken, _ := tokenAuth.Encode(map[string]interface{}{
		"userID":     u.ID,
		"username":   u.Username,
		"expiration": time.Now().Add(time.Hour * 72).Unix(),
	})
	payload := Payload{
		Username: u.Username,
		UserID:   u.ID,
		Token:    finalToken,
	}
	res := models.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       payload,
	}

	json.NewEncoder(w).Encode(res)
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := chi.URLParam(r, "id")
	db.Database.First(&u, "id = ?", userId)
	json.NewEncoder(w).Encode(User{
		ID:       u.ID,
		Username: u.Username,
	})
}
