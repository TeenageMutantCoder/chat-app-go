package main

import (
	"net/http"
)

const (
	Port = "8080"
)

type Message struct {
	Content string
}

func main() {
	messages := []Message{{"hi"}, {"hello"}, {"yo"}}

	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

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

	http.ListenAndServe(":"+Port, nil)
}
