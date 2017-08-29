package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Homepage!")
    fmt.Println("Endpoint Hit: Homepage")
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/", homePage)
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/user", CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", GetUser).Methods("GET")
    router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
  handleRequests()
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

type User struct {
  ID int `json:"id"`
  FirstName string `json:"FirstName"`
  LastName string `json:"LastName"`
  Email string `json:"Email"`
  Password string `json:"Password"`
}

type Users []User
