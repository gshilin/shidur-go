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

	type hash map[string]string
	type hash2 map[string]hash
  var (
			response hash2
			)
  response = make(hash2)
  for _, book := range books {
    response[book.Author] = hash{"links": "", "book_name": book.Title}
	}
  for _, book := range books {
    a := fmt.Sprintf("<li><a href='/books/%d'>%s</a></li>", book.ID, book.Title)
    response[book.Author]["links"] += a
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
