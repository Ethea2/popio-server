package handler

import (
	"fmt"
	"net/http"
)

type Group struct{}

func (g *Group) GetGroupsForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Groups for user")
}

func (g *Group) CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Group")
}

func (g *Group) EditGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit group")
}

func (g *Group) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Group")
}

func (g *Group) GetGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a group")
}
