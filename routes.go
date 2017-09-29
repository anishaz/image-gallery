package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  Route{
    "homePage",
    "GET",
    "/",
    homePage,
  },
  Route{
    "GetUsers",
    "GET",
    "/users",
    GetUsers,
  },
  Route{
    "CreateUser",
    "POST",
    "/user",
    CreateUser,
  },
  Route{
    "GetUser",
    "GET",
    "/user/{id}",
    GetUser,
  },
  Route{
    "DeleteUser",
    "DELETE",
    "/user/{id}",
    DeleteUser,
  },
}
