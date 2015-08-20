package routes

import (
  "net/http"
  "github.com/gshilin/shidur-slides/models"
  "fmt"
  "github.com/gorilla/mux"
)

func BooksIndex(w http.ResponseWriter, req *http.Request) {
  books := models.Books{}
  App.DB.Select("id, author, title").Find(&books)

  var response map[string]string
  response = make(map[string]string)
  for _, book := range books {
    fmt.Println(book)
    a := fmt.Sprint(response[book.Author], "<li><a href='/books/", book.ID, "'>", book.Title, "</a></li>")
    response[book.Author] = a
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  App.Render.JSON(w, http.StatusOK, response)
}

func BooksShow(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  book_id := vars["id"]

  book := models.Book{}
  App.DB.Where("id = ?", book_id).Select("slides").First(&book)

  w.Header().Set("Access-Control-Allow-Origin", "*")
  App.Render.JSON(w, http.StatusOK, book.Slides)
}
