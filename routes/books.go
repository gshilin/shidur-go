package routes

import (
	"net/http"
	"github.com/gshilin/shidur-go/models"
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
	App.DB.Where("id = ?", book_id).Select("title, slides").First(&book)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	var response map[string]string = make(map[string]string)
	response[book.Title] = book.Slides
	App.Render.JSON(w, http.StatusOK, response)
}

func BooksShowOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "X-Custom-Header")
	var response map[string]string = make(map[string]string)
	response["preFlight"] = "preFlight"
	App.Render.JSON(w, http.StatusOK, response)
}
