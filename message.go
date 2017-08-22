package main

import "time"

type Message struct {
  Id        int       `json:"id"`
  User      string    `json:"user"`
  Message   string    `json:"message"`
  CreatedOn time.Time `json:"createdOn"`
}

type Messages []Message
