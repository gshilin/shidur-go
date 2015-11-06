package models

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {

  It("saves a message", func() {

    // Make sure we have 0 messages
    messages := Messages{}
    app.DB.Find(&messages)
    Expect(len(messages)).To(BeZero())

    // Create a new message
    app.DB.Create(&Message{Message:"First Message", UserName:"Admin", Type:"question"})

    // Make sure we have 1 message
    app.DB.Find(&messages)
    Expect(len(messages)).To(Equal(1))

    // Make sure that message is the one we created
    message := Message{}
    app.DB.First(&message)
    Expect(message.Message).To(Equal("First Message"))
    Expect(message.UserName).To(Equal("Admin"))
    Expect(message.Type).To(Equal("question"))

  })

})
