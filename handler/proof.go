package handler

import (
	"fmt"
	"net/http"
)

type Proof struct{}

func (p *Proof) CreateForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create by user")
}

func (p *Proof) CreateForGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create by group")
}

func (p *Proof) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete POP")
}

func (p *Proof) Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit POP")
}

func (p *Proof) GetProofByGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get proof by group")
}

func (p *Proof) GetUserProof(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get pop by user")
}
