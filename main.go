package main

import (
	"html/template"
	"net/http"
)

const (
	PORT                        = "8080"
	STARTING_TEMPLATE_DIRECTORY = "templates/"
	TEMPLATE_FILE_TYPE          = ".gohtml"
)

type Message struct {
	Content string
}

func ParseFiles(filenames ...string) (*template.Template, error) {
	for index := range filenames {
		filenames[index] = STARTING_TEMPLATE_DIRECTORY + filenames[index] + TEMPLATE_FILE_TYPE
	}
	return template.ParseFiles(filenames...)
}

func main() {
	messages := []Message{{"hi"}, {"hello"}, {"yo"}}
	chatTemplate := template.Must(ParseFiles("base", "chat"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := chatTemplate.Execute(w, messages)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			message := r.PostFormValue("message")
			if message == "" {
				http.Error(w, "message is empty", http.StatusBadRequest)
				return
			}

			messages = append(messages, Message{message})
		}

		err := chatTemplate.ExecuteTemplate(w, "messages", messages)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":"+PORT, nil)
}
