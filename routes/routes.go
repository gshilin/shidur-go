package routes

import (
  "github.com/gshilin/shidur-slides/config"
  "github.com/gorilla/schema"
  _ "net/http"
)

var App *config.App
var decoder = schema.NewDecoder()

func Setup(app *config.App) {

  App = app

  go H.Run()

  // Define your routes here:
  var routes = config.Routes{
    config.Route{"Websocket", "GET", "/ws", ServeWs},
//    config.Route{"Root", "GET", "/", HomeIndex},

    config.Route{"MessagesIndex", "GET", "/questions", MessagesQIndex},
    config.Route{"MessagesIndex", "GET", "/messages", MessagesIndex},
    //    config.Route{"MessagesNew", "GET", "/questions/new", MessagesNew},

    config.Route{"BookmarksIndex", "GET", "/bookmarks", BookmarksIndex},
    config.Route{"BooksIndex", "GET", "/books", BooksIndex},
    config.Route{"BooksIndex", "GET", "/books/{id}", BooksShow},
  }

  config.Setup(App.Router, routes)
//  App.Router.Handle("/assets", http.FileServer(http.Dir("public")))
}
