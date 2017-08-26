package main

import (
   "fmt"
   "net/http"
)

func main() {
   http.ListenAndServe(":8080",nil)
   fmt.Println("Server is listening at port 8080")
}
