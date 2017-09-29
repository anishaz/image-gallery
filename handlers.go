package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Homepage!")
    fmt.Println("Endpoint Hit: Homepage")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
  users := Users{
    User{FirstName: "Anisha", LastName: "Zach", Email:"anisha@anisha.com", Password:"password"},
    User{FirstName: "Vinoj", LastName: "Zach", Email:"vinoj@vinoj.com", Password:"password"},
  }
  fmt.Println("Endpoint Hit: GetUsers")

  json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {}

func GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["id"]

  fmt.Fprintf(w, "Key: " + key)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
