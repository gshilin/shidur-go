package routes

import (
	"github.com/gshilin/shidur-go/config"
	_ "fmt"
)

var App *config.App

func Setup(app *config.App) {

	App = app

	go H.Run()

	// Define your routes here:
	var routes = config.Routes{
		config.Route{"Websocket", "GET", "/ws", ServeWs},
		config.Route{"Root", "GET", "/", HomeIndex},

		config.Route{"MessagesQIndex", "GET", "/questions", MessagesQIndex},
		config.Route{"MessagesQIndex", "GET", "/3questions", Messages3QIndex},
		config.Route{"MessagesIndex", "GET", "/messages", MessagesIndex},
		config.Route{"MessagesDestroy", "POST", "/messages", MessagesDestroy},
		config.Route{"MessagesDestroy", "POST", "/questions", QuestionsUnapprove},
		config.Route{"MessagesNew", "GET", "/questions/new", MessagesNew},
		config.Route{"MessagesApprove", "GET", "/questions/approve/{language}", MessagesApprove},

		config.Route{"BookmarksIndex", "GET", "/bookmarks", BookmarksIndex},
		config.Route{"BooksIndex", "GET", "/books", BooksIndex},
		config.Route{"BooksShow", "GET", "/books/{id}", BooksShow},
	}

	config.Setup(App.Router, routes)
}
