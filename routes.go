package main

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "MessageList",
    "GET",
    "/api/message",
    MessageList,
  },
  Route{
    "MessageDetail",
    "GET",
    "/api/message/{messageId}",
    MessageDetail,
  },
  Route{
    "MessageCreate",
    "POST",
    "/api/message",
    MessageCreate,
  },
}
