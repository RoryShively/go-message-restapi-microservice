package main

import "fmt"

var messageId int


var messages Messages

// Give us some seed data
func init() {
  RepoCreateMessage(Message{
    User: "kobebacon",
    Message: "Hey, what's up?",
  })
  RepoCreateMessage(Message{
    User: "nara_the_ham",
    Message: "Making a new react app with microservices, you?",
  })
  RepoCreateMessage(Message{
    User: "test_user",
    Message: "testing, testing, 1.2..3...",
  })
}

func RepoFindMessage(id int) Message {
  for _, m := range messages {
    if m.Id == id {
      return m
    }
  }
  // return empty Message if not found
  return Message{}
}

func RepoCreateMessage(m Message) Message {
  messageId += 1
  m.Id = messageId
  messages = append(messages, m)
  return m
}

func RepoDestroyMessage(id int) error {
  for i, m := range messages {
    if m.Id == id {
      messages = append(messages[:i], messages[i+1:]...) // TODO: find out about ellipses (...)
      return nil
    }
  }
  return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
