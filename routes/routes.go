package routes

import (
	_ "fmt"

	"github.com/gshilin/shidur-go/config"
)

var App *config.App

func Setup(app *config.App) {

	App = app

	go H.Run()

	// Define your routes here:
	var routes = config.Routes{
		config.Route{"Websocket", "GET", "/ws", ServeWs},
		config.Route{"Root", "GET", "/", HomeIndex},

		config.Route{"QuestionsIndex", "GET", "/questions", QuestionsIndex},
		config.Route{"Questions3Index", "GET", "/questions", Questions3Index},
		config.Route{"CongressIndex", "GET", "/congress", CongressIndex},
		config.Route{"MessagesIndex", "GET", "/messages", MessagesIndex},
		config.Route{"MessagesDestroy", "POST", "/messages", MessagesDestroy},
		config.Route{"QuestionsUnapprove", "POST", "/questions", QuestionsUnapprove},
		config.Route{"MessagesNew", "GET", "/questions/new", MessagesNew},
		config.Route{"MessagesApprove", "GET", "/questions/approve/{language}", MessagesApprove},

		config.Route{"BookmarksIndex", "GET", "/bookmarks", BookmarksIndex},
		config.Route{"BooksIndex", "GET", "/books", BooksIndex},
		config.Route{"BooksShow", "GET", "/books/{id}", BooksShow},
		config.Route{"BooksShow", "OPTIONS", "/books/{id}", BooksShowOptions},
	}

	config.Setup(App.Router, routes)
}
