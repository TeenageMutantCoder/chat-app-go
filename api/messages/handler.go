package messages

import (
	"chat/resources"
	"chat/templates"
	"net/http"
)

type templateData struct {
	Messages []resources.Message
	Name     string
}

func Get(w http.ResponseWriter, r *http.Request) {
	author, err := r.Cookie("name")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	data := templateData{resources.GetMessages(), author.Value}
	err = templates.ChatTemplate.ExecuteTemplate(w, "messages", data)
	if err != nil {
		panic(err)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	author, err := r.Cookie("name")
	if err != nil {
		http.Error(w, "name cookie not set", http.StatusBadRequest)
		return
	}

	message := r.PostFormValue("message")
	if message == "" {
		http.Error(w, "message is empty", http.StatusBadRequest)
		return
	}

	resources.AddMessage(message, author.Value)

	data := templateData{resources.GetMessages(), author.Value}
	err = templates.ChatTemplate.ExecuteTemplate(w, "messages", data)
	if err != nil {
		panic(err)
	}
}
