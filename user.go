package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "all users endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "new user endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete user endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update user endpoint hit")
}
