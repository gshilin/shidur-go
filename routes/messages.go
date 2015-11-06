package routes

import (
  "net/http"
  "github.com/gshilin/shidur-go/models"
  "strings"
)

func MessagesIndex(w http.ResponseWriter, req *http.Request) {

  type Response struct {
    Last_question models.Message    `json:"last_question"`
    Messages      []models.Message `json:"messages"`
  }

  // Find all messages in the DB
  messages := []models.Message{}
  App.DB.Order("id ASC").Find(&messages)

  replaceNewLines(messages)

  response := Response{
    Last_question: findLastQuestion(messages),
    Messages: messages,
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  App.Render.JSON(w, http.StatusOK, response)
}

func MessagesDestroy(w http.ResponseWriter, req *http.Request) {
  App.DB.Delete(models.Message{})
  w.Header().Set("Access-Control-Allow-Origin", "*")
  App.Render.JSON(w, http.StatusOK, "OK")
}

func MessagesQIndex(w http.ResponseWriter, req *http.Request) {

  type Response struct {
    Messages []models.Message `json:"questions"`
  }

  // Find all messages in the DB
  messages := []models.Message{}
  App.DB.Where("type = 'question'").Order("id ASC").Find(&messages)
  response := Response{
    Messages: messages,
  }

  replaceNewLines(response.Messages)

  // Pass them to the templates for rendering
  App.Render.JSON(w, http.StatusOK, response)
}

func MessagesNew(w http.ResponseWriter, req *http.Request) {
  App.QRender.HTML(w, http.StatusOK, "messages/new", nil)
}

func findLastQuestion(messages []models.Message) models.Message {

  for i := len(messages) - 1; i >= 0; i-- {
    v := messages[i]
    if v.Type == "question" {
      return v
    }
  }
  return models.Message{ID: 0}
}

func replaceNewLines(messages []models.Message) {
  for i := 0; i < len(messages); i++ {
    message := &messages[i]
    message.Type = strings.Trim(message.Type, " \n")
  }
}
