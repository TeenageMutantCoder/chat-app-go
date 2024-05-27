package chat

import (
	"chat/resources"
	"chat/templates"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	name, err := r.Cookie("name")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	type templateData struct {
		Messages []resources.Message
		Name     string
	}
	data := templateData{resources.GetMessages(), name.Value}
	err = templates.ChatTemplate.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
