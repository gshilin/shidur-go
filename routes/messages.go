package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/gshilin/shidur-go/models"
)

func MessagesIndex(w http.ResponseWriter, req *http.Request) {
	type Response struct {
		Last_questions []models.Message `json:"last_questions"`
		Messages       []models.Message `json:"messages"`
	}

	// Find all messages in the DB
	messages := models.Messages{}
	App.DB.Order("id ASC").Find(&messages)
	replaceNewLines(messages)

	questions := models.Messages{}
	App.DB.Where("type = 'question' AND approved = true").Order("id ASC").Find(&questions)

	response := Response{
		Last_questions: findLastQuestions(questions, false),
		Messages:       messages,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, response)
}

func MessagesDestroy(w http.ResponseWriter, req *http.Request) {
	App.DB.Delete(models.Message{})
	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, "OK")
}

func QuestionsUnapprove(w http.ResponseWriter, req *http.Request) {
	type Response struct {
		Messages []models.Message `json:"questions"`
	}

	App.DB.Model(models.Message{}).Updates(map[string]interface{}{"approved": false})

	messages := models.Messages{}
	App.DB.Where("type = 'question' AND approved = true").Order("id ASC").Find(&messages)
	response := Response{
		Messages: findLastQuestions(messages, true),
	}
	if m, err := json.Marshal(response); err != nil {
		fmt.Println("Marshal Error: ", err)
	} else {
		H.broadcast <- m
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, "OK")
}

func QuestionsIndex(w http.ResponseWriter, req *http.Request) {
	type Response struct {
		Messages []models.Message `json:"questions"`
	}

	// Find all messages in the DB
	messages := models.Messages{}
	App.DB.Where("type = 'question' AND approved = true").Order("id ASC").Find(&messages)
	response := Response{
		Messages: messages,
	}

	replaceNewLines(response.Messages)

	// Pass them to the templates for rendering
	App.Render.JSON(w, http.StatusOK, response)
}

func MessagesApprove(w http.ResponseWriter, req *http.Request) {
	var err error

	params := mux.Vars(req)
	language := params["language"]
	message := models.Message{}

	App.DB.Where("type = 'question' AND language = ?", language).Last(&message)
	if message.ID != 0 {
		message.Approved = true
		err = App.DB.Save(&message).Error
		// Broadcast message to 3-questions & congress
		if m, err := json.Marshal(message); err != nil {
			fmt.Println("Marshal Error: ", err)
		} else {
			H.broadcast <- m
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, err)
}

func Questions3Index(w http.ResponseWriter, req *http.Request) {
	type Response struct {
		Messages []models.Message `json:"questions"`
	}

	// Find all messages in the DB
	messages := models.Messages{}
	App.DB.Where("type = 'question' AND approved = true").Order("id ASC").Find(&messages)
	response := Response{
		Messages: findLastQuestions(messages, false),
	}

	replaceNewLines(response.Messages)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, response)
}

func CongressIndex(w http.ResponseWriter, req *http.Request) {
	type Response struct {
		Messages []models.Message `json:"questions"`
	}

	// Find all approved 'cg' messages in the DB
	messages := models.Messages{}
	App.DB.Where("type = 'question' AND approved = true").Order("id ASC").Find(&messages)
	response := Response{
		Messages: findLastQuestions(messages, true),
	}

	replaceNewLines(response.Messages)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	App.Render.JSON(w, http.StatusOK, response)
}

func MessagesNew(w http.ResponseWriter, req *http.Request) {
	App.QRender.HTML(w, http.StatusOK, "messages/new", nil)
}

var Languages = []string{"it", "de", "nl", "fr", "pt", "tr", "pl", "ar", "hu", "fi", "lt", "ja", "bg", "ka", "no", "sv", "hr", "zh", "fa", "ro", "hi", "ua", "mk", "sl", "lv", "sk", "cs", "am",}

func findLastQuestions(messages []models.Message, include_cg bool) (msg []models.Message) {
	x := map[string]models.Message{}
	x["en"] = models.Message{Language: "en"}
	x["he"] = models.Message{Language: "he"}
	x["ru"] = models.Message{Language: "ru"}
	x["es"] = models.Message{Language: "es"}
	for _, language := range Languages {
		x[language] = models.Message{
			Language: language,
		}
	}

	for _, v := range messages {
		if v.Type == "question" {
			x[v.Language] = v
		}
	}

	msg = append(msg, x["en"])
	msg = append(msg, x["he"])
	msg = append(msg, x["ru"])
	msg = append(msg, x["es"])
	if include_cg {
		for _, language := range Languages {
			msg = append(msg, x[language])
		}
	}

	return
}

func replaceNewLines(messages []models.Message) {
	for i := 0; i < len(messages); i++ {
		message := &messages[i]
		message.Type = strings.Trim(message.Type, " \n")
	}
}
