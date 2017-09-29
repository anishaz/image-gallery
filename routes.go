package main

import (
    "net/http"

)

type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

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
