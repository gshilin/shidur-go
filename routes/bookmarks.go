package routes

import (
	"net/http"

	"shidur-go/models"
)

type Response struct {
	Bookmarks models.Bookmarks `json:"bookmarks"`
}

func BookmarksIndex(w http.ResponseWriter, req *http.Request) {
	bookmarks := []models.Bookmark{}
	App.DB.Order("id ASC").Find(&bookmarks)

	response := Response{
		Bookmarks: bookmarks,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, response)
}
