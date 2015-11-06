package routes

import (
  "github.com/gshilin/shidur-go/models"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Messages", func() {

  Describe("#MessagesIndex", func() {

    It("renders a list of messages", func() {

      app.DB.Create(&models.Message{Message:"First Message", UserName:"Admin", Type:"question"})
      app.DB.Create(&models.Message{Message:"Second Message", UserName:"Admin", Type:"question"})

      Request("GET", "/messages", nil)
      Expect(response.Code).To(Equal(200))
      Expect(response.Body).To(ContainSubstring("last_question"))
      Expect(response.Body).To(ContainSubstring("First Message"))
      Expect(response.Body).To(ContainSubstring("Second Message"))
    })

  })

  Describe("#MessagesCreate", func() {

    It("creates a new message", func() {

      messages := []models.Message{}
      app.DB.Find(&messages)
      Expect(len(messages)).To(BeZero())

      Request("POST", "/messages", URLEncode(map[string]string{
        "message" : "First Message",
        "user_name" : "Admin",
        "type" : "question",
      }))

      Expect(response.Code).To(Equal(200))
      app.DB.Find(&messages)
      Expect(len(messages)).To(Equal(1))

      message := models.Message{}
      app.DB.First(&message)
      Expect(message.Message).To(Equal("First Message"))
      Expect(message.UserName).To(Equal("Admin"))
      Expect(message.Type).To(Equal("question"))
    })

  })

})
